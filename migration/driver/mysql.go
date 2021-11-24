package driver

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

type MySql struct {
}

func NewMySql() Strategy {
	return &MySql{}
}

func (m MySql) Driver() string {
	return "mysql"
}

func (m MySql) Init(db *sql.DB) (database.Driver, error) {
	return mysql.WithInstance(db, &mysql.Config{})
}
