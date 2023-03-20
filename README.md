# Golib Migrate

Migration solutions for Golang project.

### Setup instruction

Base setup, see [GoLib Instruction](https://gitlab.com/golibs-starter/golib/-/blob/main/README.md)

Both `go get` and `go mod` are supported.
```shell
go get gitlab.com/golibs-starter/golib-migrate
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
    "context"
    "gitlab.com/golibs-starter/golib"
    "gitlab.com/golibs-starter/golib-data"
    "gitlab.com/golibs-starter/golib-migrate"
    "gitlab.com/golibs-starter/golib/log"
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
