package redis

import (
	"ego-demo/internal/pkg/redis/example"
	"ego-demo/internal/pkg/redis/option"
	"github.com/google/wire"
)

// ProviderSet .
var ProviderSet = wire.NewSet(
	option.New,
	example.NewExampleRedis,
)
