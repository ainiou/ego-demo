package http

import (
	"ego-demo/internal/pkg/http/example"
	"github.com/google/wire"
)

// ProviderSet ...
var ProviderSet = wire.NewSet(
	example.NewHttpExample,
)
