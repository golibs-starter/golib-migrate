package golibmigrate

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib-migrate/migration"
	"github.com/golibs-starter/golib-migrate/migration/driver"
	"github.com/golibs-starter/golib/log"
	"go.uber.org/fx"
)

func MigrationOpt() fx.Option {
	return fx.Options(
		fx.Invoke(Migrate),
		fx.Provide(migration.NewMigration),
		fx.Provide(migration.NewLogging),
		fx.Provide(newDriverResolver),
		golib.ProvideProps(migration.NewProperties),
		ProvideMigrationDriverStrategy(driver.NewMySql),
		ProvideMigrationDriverStrategy(driver.NewPostgres),
		ProvideMigrationDriverStrategy(driver.NewSqlite),
	)
}

func ProvideMigrationDriverStrategy(constructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  "migration_driver_strategy",
		Target: constructor,
	})
}

type NewDriverResolverIn struct {
	fx.In
	Strategies []driver.Strategy `group:"migration_driver_strategy"`
}

func newDriverResolver(in NewDriverResolverIn) *driver.Resolver {
	return driver.NewResolver(in.Strategies)
}

func Migrate(migration *migrate.Migrate, props *migration.Properties) error {
	log.Infof("Start migrate database [%s], driver [%s], source [%s]",
		props.Database, props.Driver, props.MigrationSource)
	err := migration.Up()
	if err == migrate.ErrNoChange {
		log.Infof("No change when migrate database [%s], driver [%s], source [%s]",
			props.Database, props.Driver, props.MigrationSource)
		return nil
	}
	if err != nil {
		return err
	}
	log.Infof("Finish migrate database [%s], driver [%s], source [%s]",
		props.Database, props.Driver, props.MigrationSource)
	return nil
}
