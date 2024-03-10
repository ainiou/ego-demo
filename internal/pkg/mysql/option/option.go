package option

import "github.com/ego-component/egorm"

// Option ...
// 所有的mysql依赖
type Option struct {
	ExampleDB *egorm.Component
}

// New ...
func New() *Option {
	return &Option{
		ExampleDB: egorm.Load("mysql.example").Build(),
	}
}
