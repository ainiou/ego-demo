package pgsql

import (
	"ego-demo/internal/pkg/pgsql/example"
	"ego-demo/internal/pkg/pgsql/option"
	"github.com/google/wire"
)

// ProviderSet ...
var ProviderSet = wire.NewSet(
	option.New,
	example.NewExampleInterface,
)
