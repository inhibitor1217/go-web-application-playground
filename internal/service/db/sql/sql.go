package sql

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/jmoiron/sqlx"
)

func NewSqlDB(e *env.Env) (*sqlx.DB, error) {
	return sqlx.Open("postgres", e.PSQL.DatasourceName())
}
