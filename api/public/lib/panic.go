package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/inhibitor1217/go-web-application-playground/api/public/views"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
)

func Panic(
	cx *gin.Context,
	err error,
	l *log.Logger,
) {
	cx.AbortWithError(http.StatusInternalServerError, err)
	views.Panic(cx, err)

	details := []log.Field{}
	traceId, exists := cx.Get("trace_id")
	if exists {
		details = append(details, log.Field{
			Key:   "trace_id",
			Value: traceId.(uuid.UUID).String(),
		})
	}
	l.Error(err, details...)
}
