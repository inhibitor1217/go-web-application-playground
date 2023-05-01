package jwt

import "go.uber.org/fx"

var Option = fx.Option(
	fx.Provide(NewJwt),
)
