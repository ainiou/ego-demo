package example

import (
	"context"
	"ego-demo/internal/pkg/redis/option"
)

type Example struct {
	*option.Option
}

func NewExampleRedis(opt *option.Option) *Example {
	return &Example{
		Option: opt,
	}
}

// GetHelloWorld ...
func (h *Example) GetHelloWorld(ctx context.Context, key string) (res string, err error) {
	return h.ExampleRedis.Get(ctx, key)
}
