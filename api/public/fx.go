package public

import (
	"github.com/inhibitor1217/go-web-application-playground/api/public/middlewares"
	"github.com/inhibitor1217/go-web-application-playground/api/public/routes"
	"github.com/inhibitor1217/go-web-application-playground/internal/features/account"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/db/sql"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/zap"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewHttpServerModule() fx.Option {
	return fx.Module(
		"public-api",

		// internal/lib
		env.Option,
		log.Option,

		// internal/service
		sql.Option,
		zap.Option,

		// features
		account.Option,

		// middlewares
		middlewares.Option,

		// routes
		routes.Option,

		fx.WithLogger(func(logger *log.Logger) fxevent.Logger {
			return logger
		}),

		fx.Provide(
			fx.Annotate(
				func(e *env.Env) string {
					return e.PublicHttp.Port
				},
				fx.ResultTags(`name:"public-api:http-port"`),
			),
		),

		fx.Provide(
			fx.Annotate(
				http.NewServer,
				fx.ParamTags(`group:"routes:public"`, `group:"middlewares:public"`, `name:"public-api:http-port"`),
				fx.ResultTags(`group:"servers"`, `name:"public-api"`),
			),
		),
	)
}
