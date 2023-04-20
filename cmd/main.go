package main

import (
	log "github.com/inhibitor1217/go-web-application-playground/internal/lib/log"
	zap "github.com/inhibitor1217/go-web-application-playground/internal/service/zap"

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
		zap.Option,
	).Run()
}
