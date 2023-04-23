package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
)

type Cors struct {
	handler gin.HandlerFunc
}

func NewCors(e *env.Env) *Cors {
	handler := cors.New(cors.Config{
		AllowMethods: []string{"GET", "PATCH", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Length", "Content-Type", "Origin"},
		AllowOriginFunc: func(origin string) bool {
			if e.IsDevelopment() {
				return true
			} else {
				// TODO
				return false
			}
		},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	})

	return &Cors{
		handler: handler,
	}
}

func (c *Cors) Handler() gin.HandlerFunc {
	return c.handler
}
