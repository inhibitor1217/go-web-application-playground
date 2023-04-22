package swagger

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Path() string {
	return "/"
}

func (h *Handler) Register(r http.Router) {
	r.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
