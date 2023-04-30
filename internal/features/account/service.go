package account

import (
	"context"
)

type Service interface {
	Find(cx context.Context, id string) (Account, error)
	FindByEmail(cx context.Context, email string) (Account, error)
	ExistsOfEmail(cx context.Context, email string) (bool, error)
	Create(cx context.Context, dto *CreateDTO) (Account, error)
}

type CreateDTO struct {
	Email       string
	Password    string
	DisplayName *string
}
