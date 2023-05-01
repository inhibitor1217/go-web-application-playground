package auth

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/auth"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
)

type Handler struct {
	accountSvc account.Service
	auth       *auth.Authenticator
	l          *log.Logger
}

func NewHandler(
	accountSvc account.Service,
	auth *auth.Authenticator,
	l *log.Logger,
) *Handler {
	return &Handler{
		accountSvc: accountSvc,
		auth:       auth,
		l:          l,
	}
}

func (h *Handler) Path() string {
	return "/auth"
}

func (h *Handler) Register(r http.Router) {
	r.POST("/sign-up", h.SignUp)
	r.POST("/sign-in", h.SignIn)
	r.POST("/touch", h.Touch)
	r.DELETE("/sign-out", h.SignOut)
}
