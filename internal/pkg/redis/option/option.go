package option

import (
	"github.com/gotomicro/ego-component/eredis"
)

type Option struct {
	ExampleRedis *eredis.Component
}

func New() *Option {
	return &Option{
		ExampleRedis: eredis.Load("redis.example").Build(),
	}
}
