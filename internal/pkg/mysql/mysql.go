package mysql

import (
	"ego-demo/internal/pkg/mysql/example"
	"ego-demo/internal/pkg/mysql/option"
	"github.com/google/wire"
)

// ProviderSet ...
var ProviderSet = wire.NewSet(
	option.New,
	example.NewExampleInterface,
)
