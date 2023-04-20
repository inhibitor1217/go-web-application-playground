package public

import (
	"github.com/inhibitor1217/go-web-application-playground/api/public/routes"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/zap"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewHttpServerModule() fx.Option {
	return fx.Module(
		"http-api",

		// internal/lib
		env.Option,
		log.Option,

		// internal/service
		zap.Option,

		// routes
		routes.Option,

		fx.WithLogger(func(logger *log.Logger) fxevent.Logger {
			return logger
		}),

		fx.Provide(
			fx.Annotate(
				http.NewServer,
				fx.ParamTags(`group:"routes"`),
			),
		),
	)
}
