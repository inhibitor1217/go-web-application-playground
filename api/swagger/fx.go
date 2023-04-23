package swagger

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/docs"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"go.uber.org/fx"
)

func NewSwaggerModule() fx.Option {
	return fx.Module(
		"swagger",

		// internal/lib
		docs.Option,

		fx.Provide(
			fx.Annotate(
				NewHandler,
				fx.As(new(http.Routes)),
				fx.ResultTags(`group:"routes:swagger"`),
			),
		),

		fx.Provide(
			fx.Annotate(
				func(e *env.Env) string {
					return e.Swagger.Port
				},
				fx.ResultTags(`name:"swagger:http-port"`),
			),
		),

		fx.Provide(
			fx.Annotate(
				http.NewServer,
				fx.ParamTags(`group:"routes:swagger"`, `group:"middlewares:swagger"`, `name:"swagger:http-port"`),
				fx.ResultTags(`group:"servers"`, `name:"swagger"`),
			),
		),
	)
}
