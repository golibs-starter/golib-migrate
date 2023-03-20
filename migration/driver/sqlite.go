package driver

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
)

type Sqlite struct {
}

func NewSqlite() Strategy {
	return &Sqlite{}
}

func (m Sqlite) Driver() string {
	return "sqlite"
}

func (m Sqlite) Init(db *sql.DB) (database.Driver, error) {
	return sqlite.WithInstance(db, &sqlite.Config{})
}
