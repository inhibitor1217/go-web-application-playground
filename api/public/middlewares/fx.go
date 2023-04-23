package middlewares

import (
	"github.com/inhibitor1217/go-web-application-playground/api/public/middlewares/cors"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"go.uber.org/fx"
)

var Option = fx.Provide(
	middleware(cors.NewCors),
)

func middleware(fn interface{}) interface{} {
	return fx.Annotate(
		fn,
		fx.As(new(http.Middleware)),
		fx.ResultTags(`group:"middlewares:public"`),
	)
}
