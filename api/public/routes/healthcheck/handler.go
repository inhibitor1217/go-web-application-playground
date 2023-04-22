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

// healthcheck godoc
//
//	@Summary		Healthcheck
//	@Description	Checks if the application is healthy
//	@Tags			Utility
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	healthcheck.healthcheck.view
//	@Router			/healthcheck [get]
func (h *Handler) healthcheck(cx *gin.Context) {
	type view struct {
		AppName  string `json:"app_name"`
		AppStage string `json:"app_stage"`
		AppBuild string `json:"app_build"`
	}

	cx.JSON(http.OK, view{
		AppName:  h.env.App.Name,
		AppStage: h.env.App.Stage.String(),
		AppBuild: h.env.App.Build,
	})
}
