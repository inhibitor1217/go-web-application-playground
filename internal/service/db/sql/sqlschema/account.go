package sqlschema

import (
	"database/sql"
	"time"
)

type Account struct {
	SId           string       `db:"id"`
	SCreatedAt    time.Time    `db:"created_at"`
	SUpdatedAt    time.Time    `db:"updated_at"`
	SEmail        string       `db:"email"`
	SPasswordHash string       `db:"password_hash"`
	SDisplayName  *string      `db:"display_name"`
	STouchedAt    sql.NullTime `db:"touched_at"`
}

func (a *Account) TypeName() string {
	return "account"
}

func (a *Account) Id() string {
	return a.SId
}

func (a *Account) CreatedAt() time.Time {
	return a.SCreatedAt
}

func (a *Account) UpdatedAt() time.Time {
	return a.SUpdatedAt
}

func (a *Account) Email() string {
	return a.SEmail
}

func (a *Account) PasswordHash() string {
	return a.SPasswordHash
}

func (a *Account) DisplayName() *string {
	return a.SDisplayName
}

func (a *Account) TouchedAt() *time.Time {
	if a.STouchedAt.Valid {
		return &a.STouchedAt.Time
	}
	return nil
}
