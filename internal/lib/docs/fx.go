package docs

import (
	"github.com/inhibitor1217/go-web-application-playground/docs"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"go.uber.org/fx"
)

var Option = fx.Option(
	fx.Invoke(fillDocsInfo),
)

func fillDocsInfo(e *env.Env) {
	docs.SwaggerInfo.Title = e.App.Name
	docs.SwaggerInfo.Host = e.PublicHttp.BaseUrl.Host
	docs.SwaggerInfo.Version = e.App.Build
}
