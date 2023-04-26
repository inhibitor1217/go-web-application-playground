package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account/model"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/db/sql/sqlmapper"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/db/sql/sqlschema"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Service interface {
	Create(cx context.Context, dto *CreateDTO) (*model.Account, error)
	ExistsOfEmail(cx context.Context, email string) (bool, error)
	FindByEmail(cx context.Context, email string) (*model.Account, error)
}

type CreateDTO struct {
	Email       string
	Password    string
	DisplayName string
}

type service struct {
	sql *sqlx.DB
}

func NewService(
	sql *sqlx.DB,
) Service {
	// TODO wrap the service with ctx cancellation
	return &service{
		sql: sql,
	}
}

func (svc *service) Create(cx context.Context, dto *CreateDTO) (*model.Account, error) {
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
		dto.Password, // TODO hash
		dto.DisplayName,
	)

	if row.Err() != nil {
		return nil, errors.WithStack(row.Err())
	}

	a := sqlschema.Account{}
	if err := row.StructScan(&a); err != nil {
		return nil, errors.WithStack(err)
	}

	createdAt, err := sqlmapper.MapTimestampToTime(&a.CreatedAt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	updatedAt, err := sqlmapper.MapTimestampToTime(&a.UpdatedAt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	touchedAt, err := sqlmapper.MapTimestampToTime(a.TouchedAt)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.Account{
		Id:           a.Id,
		CreatedAt:    *createdAt,
		UpdatedAt:    *updatedAt,
		Email:        a.Email,
		PasswordHash: a.PasswordHash,
		DisplayName:  a.DisplayName,
		TouchedAt:    touchedAt,
	}, nil
}

func (svc *service) ExistsOfEmail(cx context.Context, email string) (bool, error) {
	a := []sqlschema.Account{}
	if err := svc.sql.Select(
		&a,
		`
		SELECT * FROM accounts
			WHERE email = $1
			LIMIT 1
		`,
		email,
	); err != nil {
		return false, errors.WithStack(err)
	}

	return len(a) > 0, nil
}

func (svc *service) FindByEmail(cx context.Context, email string) (*model.Account, error) {
	return nil, errors.WithStack(errors.New("TODO"))
}
