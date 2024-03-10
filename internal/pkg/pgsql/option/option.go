package option

import "github.com/gotomicro/ego-component/egorm"

// Option ...
type Option struct {
	ExamplePG *egorm.Component
}

// New ...
func New() *Option {
	return &Option{
		ExamplePG: egorm.Load("pgsql.example").Build(),
	}
}
