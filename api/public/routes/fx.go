package routes

import (
	"github.com/inhibitor1217/go-web-application-playground/api/public/routes/ping"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"go.uber.org/fx"
)

var Option = fx.Provide(
	fx.Annotate(
		ping.NewHandler,
		fx.As(new(http.Routes)),
		fx.ResultTags(`group:"routes"`),
	),
)
