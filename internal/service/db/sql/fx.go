package sql

import "go.uber.org/fx"

var Option = fx.Provide(NewSqlDB)
