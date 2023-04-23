package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
)

type ServerConfig struct {
	Host string
	Port string
}

type Server struct {
	engine *gin.Engine
	config ServerConfig
}

const (
	allHost = "0.0.0.0"
)

func NewServer(routes []Routes, middlewares []Middleware, port string, e *env.Env, l *log.Logger) (*Server, error) {
	gin.SetMode(selectMode(e))

	server := &Server{
		engine: gin.Default(),
		config: ServerConfig{
			Host: allHost,
			Port: port,
		},
	}

	for _, middleware := range middlewares {
		server.engine.Use(middleware.Handler())
	}

	for _, route := range routes {
		path := route.Path()
		route.Register(server.engine.Group(path))
	}

	return server, nil
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
}

func (s *Server) Run() error {
	return s.engine.Run(s.Addr())
}

func selectMode(e *env.Env) string {
	if e.IsDevelopment() {
		return gin.DebugMode
	} else {
		return gin.ReleaseMode
	}
}
