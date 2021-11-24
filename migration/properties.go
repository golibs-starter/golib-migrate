package migration

import (
	_ "gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib/config"
)

func NewProperties(loader config.Loader) (*Properties, error) {
	props := Properties{}
	err := loader.Bind(&props)
	return &props, err
}

type Properties struct {
	Driver          string
	Database        string
	MigrationSource string `default:"file:///migrations"`
}

func (p Properties) Prefix() string {
	return "app.datasource"
}
