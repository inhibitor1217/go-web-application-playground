package account

import "go.uber.org/fx"

var Option = fx.Options(
	fx.Provide(NewService),
)
