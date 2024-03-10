package example

import (
	"context"
	"ego-demo/internal/pkg/pgsql/option"
	"ego-demo/internal/pkg/pgsql/pgdb"
)

type Example struct {
	*option.Option
}

func NewExampleInterface(opt *option.Option) *Example {
	return &Example{
		Option: opt,
	}
}

// GetRecord ...
func (h *Example) GetRecord(ctx context.Context) (record pgdb.HelloWorld, err error) {
	if err = h.ExamplePG.WithContext(ctx).Model(new(pgdb.HelloWorld)).Find(&record).Error; err != nil {
		return
	}
	return
}
