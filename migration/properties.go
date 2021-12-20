package migration

import (
	_ "gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib/config"
)

func NewProperties(loader config.Loader) (*Properties, error) {
	props := Properties{}
	err := loader.Bind(&props)
	return &props, err
}

type Properties struct {
	Driver          string `validate:"required"`
	Database        string `validate:"required"`
	MigrationSource string `validate:"required" default:"file:///migrations"`
}

func (p Properties) Prefix() string {
	return "app.datasource"
}
