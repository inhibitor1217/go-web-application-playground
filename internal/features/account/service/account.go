package service

import (
	"errors"

	"github.com/inhibitor1217/go-web-application-playground/internal/features/account/model"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type CreateDTO struct {
	Email       string `binding:"required"`
	Password    string `binding:"required"`
	DisplayName string
}

func (svc *Service) Create(dto *CreateDTO) (*model.Account, error) {
	return nil, errors.New("TODO")
}

func (svc *Service) ExistsOfEmail(email string) (bool, error) {
	return false, errors.New("TODO")
}
