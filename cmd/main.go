package main

import (
	"github.com/inhibitor1217/go-web-application-playground/api/public"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/godotenv"

	"go.uber.org/fx"
)

func main() {
	godotenv.LoadEnv()

	fx.New(
		public.NewHttpServerModule(),

		fx.Invoke(func(s *http.Server, e *env.Env, l *log.Logger) {
			l.Info("Starting http server", log.String("app_name", e.App.Name), log.String("app_stage", string(e.App.Stage)), log.String("http_addr", s.Addr()))
			if err := s.Run(); err != nil {
				l.Fatal("Failed to run http server", log.Error(err))
			}
		}),
	).Run()
}
