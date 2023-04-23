package healthcheck

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
)

type Handler struct {
	env *env.Env
}

func NewHandler(e *env.Env) *Handler {
	return &Handler{
		env: e,
	}
}

func (h *Handler) Path() string {
	return "/healthcheck"
}

func (h *Handler) Register(r http.Router) {
	r.GET("", h.Healthcheck)
}
