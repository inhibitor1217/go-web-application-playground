package log

import (
	"go.uber.org/fx"
)

var Option = fx.Options(
	fx.Provide(NewLogger),
)
