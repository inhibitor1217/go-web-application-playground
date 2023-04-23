package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(cx *gin.Context) {
	cx.JSON(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) SignIn(cx *gin.Context) {
	cx.JSON(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) Touch(cx *gin.Context) {
	cx.JSON(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) SignOut(cx *gin.Context) {
	cx.JSON(http.StatusNotImplemented, "Not implemented")
}
