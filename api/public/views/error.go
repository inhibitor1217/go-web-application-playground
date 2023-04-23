package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorView[T interface{}] struct {
	Type    string `json:"type" binding:"required"`
	Message string `json:"msg"`
	Payload T      `json:"payload"`
}

func ClientError[T interface{}](cx *gin.Context, e ErrorView[T]) {
	cx.JSON(http.StatusBadRequest, e)
}

func ValidationError(cx *gin.Context, err error) {
	e := ErrorView[interface{}]{
		Type:    "validation_error",
		Message: err.Error(),
	}

	cx.JSON(http.StatusBadRequest, e)
}

func Panic(cx *gin.Context, err error) {
	e := ErrorView[interface{}]{
		Type:    "panic",
		Message: "Internal server error",
	}

	cx.JSON(http.StatusInternalServerError, e)
}
