package migration

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/golibs-starter/golib-migrate/migration/driver"
)

func NewMigration(db *sql.DB, resolver *driver.Resolver, props *Properties, logger migrate.Logger) (*migrate.Migrate, error) {
	if props.MigrationSource == "" {
		return nil, errors.New("missing migrationSource")
	}
	strategy, err := resolver.Resolve(props.Driver)
	if err != nil {
		return nil, err
	}
	drive, err := strategy.Init(db)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot init migration driver")
	}
	m, err := migrate.NewWithDatabaseInstance(props.MigrationSource, props.Database, drive)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot create migrate instance")
	}
	m.Log = logger
	return m, nil
}
