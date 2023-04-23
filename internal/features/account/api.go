package account

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account/model"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account/service"
)

type Account = model.Account

type Service = service.Service

var NewService = service.NewService

type CreateDTO = service.CreateDTO
