package service

import (
	"context"

	"github.com/inhibitor1217/go-web-application-playground/internal/features/account/model"
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

type service struct{}

func NewService() Service {
	// TODO wrap the service with ctx cancellation
	return &service{}
}

func (svc *service) Create(cx context.Context, dto *CreateDTO) (*model.Account, error) {
	return nil, errors.WithStack(errors.New("TODO"))
}

func (svc *service) ExistsOfEmail(cx context.Context, email string) (bool, error) {
	return false, errors.WithStack(errors.New("TODO"))
}

func (svc *service) FindByEmail(cx context.Context, email string) (*model.Account, error) {
	return nil, errors.WithStack(errors.New("TODO"))
}
