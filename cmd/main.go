package main

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/godotenv"
	"github.com/inhibitor1217/go-web-application-playground/internal/service/zap"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		fx.WithLogger(func(logger *log.Logger) fxevent.Logger {
			return logger
		}),

		// internal/lib
		log.Option,

		// internal/service
		godotenv.Option,
		zap.Option,
	).Run()
}
