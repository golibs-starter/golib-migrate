package driver

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

type Postgres struct {
}

func NewPostgres() Strategy {
	return &Postgres{}
}

func (m Postgres) Driver() string {
	return "postgres"
}

func (m Postgres) Init(db *sql.DB) (database.Driver, error) {
	return postgres.WithInstance(db, &postgres.Config{})
}
