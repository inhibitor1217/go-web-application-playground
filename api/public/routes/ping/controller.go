package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	cx.JSON(http.StatusOK, "pong")
}
