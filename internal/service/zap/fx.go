package zap

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Option = fx.Option(
	fx.Provide(zap.NewDevelopment),
)
