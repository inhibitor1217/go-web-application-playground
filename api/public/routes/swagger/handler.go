package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Path() string {
	return "/swagger"
}

func (h *Handler) Register(r http.Router) {
	r.GET("", func(cx *gin.Context) {
		cx.Redirect(http.MOVED_PERMANENTLY, "/index.html")
	})
	r.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
