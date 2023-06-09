package routes

import (
	"github.com/inhibitor1217/go-web-application-playground/api/public/routes/accounts"
	"github.com/inhibitor1217/go-web-application-playground/api/public/routes/auth"
	"github.com/inhibitor1217/go-web-application-playground/api/public/routes/healthcheck"
	"github.com/inhibitor1217/go-web-application-playground/api/public/routes/ping"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"go.uber.org/fx"
)

var Option = fx.Provide(
	route(accounts.NewHandler),
	route(auth.NewHandler),
	route(healthcheck.NewHandler),
	route(ping.NewHandler),
)

func route(fn interface{}) interface{} {
	return fx.Annotate(
		fn,
		fx.As(new(http.Routes)),
		fx.ResultTags(`group:"routes:public"`),
	)
}
