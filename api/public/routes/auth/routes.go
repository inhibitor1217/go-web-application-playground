package auth

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
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
