package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Trace struct{}

func NewTrace() *Trace {
	return &Trace{}
}

func (t *Trace) Handler() gin.HandlerFunc {
	return func(cx *gin.Context) {
		traceId := uuid.New()
		cx.Set("trace_id", traceId)
		cx.Header("X-Trace-Id", traceId.String())

		cx.Next()
	}
}
