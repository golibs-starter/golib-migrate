package driver

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4/database"
)

type Strategy interface {

	// Driver returns the driver name
	Driver() string

	// Init database driver
	Init(db *sql.DB) (database.Driver, error)
}
