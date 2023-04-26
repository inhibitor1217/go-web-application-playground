package model

import (
	"time"
)

type Account interface {
	Id() string
	CreatedAt() time.Time
	UpdatedAt() time.Time
	Email() string
	PasswordHash() string
	DisplayName() *string
	TouchedAt() *time.Time
}

func TypeName() string {
	return "Account"
}
