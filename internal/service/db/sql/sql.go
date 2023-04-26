package sql

import (
	"github.com/inhibitor1217/go-web-application-playground/internal/lib/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewSqlDB(e *env.Env) (*sqlx.DB, error) {
	return sqlx.Open("postgres", e.PSQL.DatasourceName())
}
