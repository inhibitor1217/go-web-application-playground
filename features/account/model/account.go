package model

import (
	"time"
)

type Account struct {
	Id           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string
	PasswordHash string
	DisplayName  string
	TouchedAt    time.Time
}

func (a *Account) Identifier() string {
	return a.Id
}

func (a *Account) TypeName() string {
	return "Account"
}
