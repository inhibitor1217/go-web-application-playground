package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/jwt"
	"github.com/pkg/errors"
)

var AuthRequired = errors.New("authentication required")

type Authenticator interface {
	Sign(cx *gin.Context, p Principal) error
	Authenticate(cx *gin.Context) (Principal, error)
	WillExpire(cx *gin.Context) bool
	Refresh(cx *gin.Context) (Principal, error)
	Clear(cx *gin.Context) error
}

func NewAuthenticator(
	accountSvc account.Service,
	e *env.Env,
	j *jwt.Jwt,
) Authenticator {
	return &tokenAuthenticator{
		accountSvc: accountSvc,
		env:        e,
		jwt:        j,
	}
}
