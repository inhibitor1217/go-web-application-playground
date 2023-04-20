package main

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
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
		log.Option,

		// internal/service
		zap.Option,

		fx.WithLogger(func(logger *log.Logger) fxevent.Logger {
			return logger
		}),
	).Run()
}
