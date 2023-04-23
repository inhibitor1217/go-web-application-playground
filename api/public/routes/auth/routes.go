package auth

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
)

type Handler struct {
	accountSvc *account.Service
}

func NewHandler(
	accountSvc *account.Service,
) *Handler {
	return &Handler{
		accountSvc: accountSvc,
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
