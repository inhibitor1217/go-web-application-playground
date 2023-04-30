package account

import (
	"time"

	"github.com/inhibitor1217/go-web-application-playground/internal/lib/crypto"
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

func ValidatePassword(a Account, password string) (bool, error) {
	return crypto.Validate(password, a.PasswordHash())
}
