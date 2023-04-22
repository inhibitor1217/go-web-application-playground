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

// ping godoc
//
//	@Summary		Ping
//	@Description	Ping
//	@Tags			Utility
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"pong"
//	@Router			/ping [get]
func (h *Handler) ping(cx *gin.Context) {
	cx.JSON(http.OK, "pong")
}
