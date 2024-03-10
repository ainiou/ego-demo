package grpc

import (
	"ego-demo/internal/pkg/grpc/example"
	"github.com/google/wire"
)

// ProviderSet ...
var ProviderSet = wire.NewSet(
	example.NewExampleInterface,
)
