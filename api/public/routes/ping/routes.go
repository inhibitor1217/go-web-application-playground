package ping

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Path() string {
	return "/ping"
}

func (h *Handler) Register(r http.Router) {
	r.GET("", h.Ping)
}
