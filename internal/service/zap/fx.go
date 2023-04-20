package zap

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Option = fx.Option(
	fx.Provide(func(e *env.Env) (*zap.Logger, error) {
		return zap.NewDevelopment()
	}),
)

func NewLogger(e *env.Env) (*zap.Logger, error) {
	switch e.App.Stage {
	case env.AppStageDev:
		return zap.NewDevelopment()
	default:
		return zap.NewProduction()
	}
}
