package main

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/http"
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/godotenv"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/zap"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	godotenv.LoadEnv()

	fx.New(
		// internal/lib
		env.Option,
		http.Option,
		log.Option,

		// internal/service
		zap.Option,

		fx.WithLogger(func(logger *log.Logger) fxevent.Logger {
			return logger
		}),

		fx.Invoke(func(s *http.HttpServer, e *env.Env, l *log.Logger) {
			l.Info("Starting http server", log.String("app_name", e.App.Name), log.String("app_stage", string(e.App.Stage)), log.String("http_addr", s.Addr()))
			if err := s.Run(); err != nil {
				l.Fatal("Failed to run http server", log.Error(err))
			}
		}),
	).Run()
}
