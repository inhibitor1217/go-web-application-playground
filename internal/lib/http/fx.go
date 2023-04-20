package http

import "go.uber.org/fx"

var Option = fx.Option(
	fx.Provide(NewHttpServer),
)
