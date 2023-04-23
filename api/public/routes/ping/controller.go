package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
)

// Ping godoc
//
//	@Summary		Ping
//	@Description	Ping
//	@Tags			Utility
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"pong"
//	@Router			/ping [get]
func (h *Handler) Ping(cx *gin.Context) {
	cx.JSON(http.OK, "pong")
}
