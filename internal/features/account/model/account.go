package model

import (
	"time"

	"github.com/inhibitor1217/go-web-application-playground/internal/lib/entity"
)

type Account struct {
	Id           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string
	PasswordHash string
	DisplayName  *string
	TouchedAt    *time.Time
}

func (a *Account) Identifier() string {
	return a.Id
}

func (a *Account) TypeName() string {
	return "Account"
}

func (a *Account) String() string {
	return entity.String(a)
}
