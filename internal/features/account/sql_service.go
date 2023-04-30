package account

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/crypto"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/db/sql/sqlschema"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type sqlService struct {
	sql *sqlx.DB
}

func NewSQLService(
	sql *sqlx.DB,
) Service {
	// TODO wrap the service with ctx cancellation
	return &sqlService{
		sql: sql,
	}
}

func (svc *sqlService) Find(cx context.Context, id string) (Account, error) {
	a := sqlschema.Account{}
	err := svc.sql.Get(
		&a,
		"SELECT * FROM accounts WHERE id = $1",
		id,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return &a, nil
}

func (svc *sqlService) FindByEmail(cx context.Context, email string) (Account, error) {
	a := sqlschema.Account{}
	err := svc.sql.Get(
		&a,
		"SELECT * FROM accounts WHERE email = $1",
		email,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return &a, nil
}

func (svc *sqlService) ExistsOfEmail(cx context.Context, email string) (bool, error) {
	a := sqlschema.Account{}
	err := svc.sql.Get(
		&a,
		"SELECT * FROM accounts WHERE email = $1",
		email,
	)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, errors.WithStack(err)
	}

	return true, nil
}

func (svc *sqlService) Create(cx context.Context, dto *CreateDTO) (Account, error) {
	passwordHash, err := crypto.Hash(dto.Password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	row := svc.sql.QueryRowx(
		`
		INSERT INTO accounts (
			id,
			email,
			password_hash,
			display_name
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) RETURNING *
		`,
		uuid.New().String(),
		dto.Email,
		passwordHash,
		dto.DisplayName,
	)

	if row.Err() != nil {
		return nil, errors.WithStack(row.Err())
	}

	a := sqlschema.Account{}
	if err := row.StructScan(&a); err != nil {
		return nil, errors.WithStack(err)
	}

	return &a, nil
}
