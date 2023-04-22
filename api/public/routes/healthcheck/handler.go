package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
)

type Handler struct {
	env *env.Env
}

func NewHandler(e *env.Env) *Handler {
	return &Handler{
		env: e,
	}
}

func (h *Handler) Path() string {
	return "/healthcheck"
}

func (h *Handler) Register(r http.Router) {
	r.GET("", h.healthcheck)
}

func (h *Handler) healthcheck(cx *gin.Context) {
	cx.JSON(http.OK, HealthcheckView{
		AppName:  h.env.App.Name,
		AppStage: h.env.App.Stage.String(),
		AppBuild: h.env.App.Build,
	})
}
