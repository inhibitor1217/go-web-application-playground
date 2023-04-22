package ping

import (
	"github.com/gin-gonic/gin"
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
	r.GET("", h.ping)
}

func (h *Handler) ping(cx *gin.Context) {
	cx.JSON(http.OK, "pong")
}