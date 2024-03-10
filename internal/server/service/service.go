package service

import (
	"ego-demo/internal/pkg/http"
	"ego-demo/internal/server/service/example"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(

	example.NewService,
	wire.Struct(new(example.Option), "*"),

	http.ProviderSet,
)
