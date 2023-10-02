# Golib Migrate

> **Note**
> We are moving out from [Gitlab](https://gitlab.com/golibs-starter). All packages are now migrated
> to `github.com/golibs-starter/*`. Please consider updating.

Migration solutions for Golang project.

### Setup instruction

Base setup, see [GoLib Instruction](https://github.com/golibs-starter/golib#readme)

Both `go get` and `go mod` are supported.

```shell
go get github.com/golibs-starter/golib-migrate
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
	"context"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib-data"
	"github.com/golibs-starter/golib-migrate"
	"github.com/golibs-starter/golib/log"
	"go.uber.org/fx"
)

func main() {
	if err := fx.New(
		// Required options for migration
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),
		golibdata.DatasourceOpt(),

		// When you want to run migration
		golibmigrate.MigrationOpt(),
	).Start(context.Background()); err != nil {
		log.Fatal("Error when migrate database: ", err)
	}
}
```

### Configuration

```yaml
app:
    datasource:
        # Define the database driver.
        # Supports: mysql, postgres, sqlite
        driver: mysql
        host: localhost
        port: 3306
        database: sample
        username: root
        password: secret
        # Define location of migration files
        migrationSource: file://migrations
```
