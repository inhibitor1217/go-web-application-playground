package http

import "github.com/gin-gonic/gin"

type Router interface {
	gin.IRouter
}

type Routes interface {
	Path() string
	Register(r Router)
}
