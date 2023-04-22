package main

import (
	"github.com/inhibitor1217/go-web-application-playground/api/public"
	"github.com/inhibitor1217/go-web-application-playground/api/swagger"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.LoadEnv()

	fx.New(
		public.NewHttpServerModule(),
		swagger.NewSwaggerModule(),

		fx.Invoke(
			fx.Annotate(
				runServers,
				fx.ParamTags(`group:"servers"`),
			),
		),
	).Run()
}

func runServers(s []*http.Server, l *log.Logger) {
	for _, s := range s {
		go runServer(s, l)
	}
}

func runServer(s *http.Server, l *log.Logger) {
	if err := s.Run(); err != nil {
		l.Fatal("Failed to run http server", log.Error(err))
	}
}
