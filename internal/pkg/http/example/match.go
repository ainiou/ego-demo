package example

import (
	"github.com/gotomicro/ego/client/ehttp"
)

type HttpExample struct {
	client *ehttp.Component
}

func NewHttpExample() *HttpExample {
	return &HttpExample{
		client: ehttp.Load("http.matchsrv").Build(),
	}
}
