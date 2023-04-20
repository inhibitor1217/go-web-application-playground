package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
)

type HttpServerConfig struct {
	Host string
	Port string
}

type HttpServer struct {
	engine *gin.Engine
	config HttpServerConfig
}

func NewHttpServer(e *env.Env) (*HttpServer, error) {
	gin.SetMode(selectMode(e))
	return &HttpServer{
		engine: gin.Default(),
		config: HttpServerConfig{
			Host: "0.0.0.0",
			Port: e.Http.Port,
		},
	}, nil
}

func (s *HttpServer) Addr() string {
	return fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
}

func (s *HttpServer) Run() error {
	return s.engine.Run(s.Addr())
}

func selectMode(e *env.Env) string {
	if e.IsDevelopment() {
		return gin.DebugMode
	} else {
		return gin.ReleaseMode
	}
}
