package trace

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Trace struct{}

func NewTrace() *Trace {
	return &Trace{}
}

func (t *Trace) Handler() gin.HandlerFunc {
	return handler
}

func handler(cx *gin.Context) {
	traceId := uuid.New()
	cx.Set("trace_id", traceId)
	cx.Header("X-Trace-Id", traceId.String())

	cx.Next()
}
