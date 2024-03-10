package application

import (
	"ego-demo/internal/server/controller"
	"github.com/gotomicro/ego/server/egin"
)

type HttpServer struct {
	ginServer *egin.Component
	*controller.Options
}

func NewHttpServer(opts *controller.Options) *HttpServer {
	s := egin.Load("server.http").Build()

	g := s.Group("/api/example-app")
	// 注册路由
	opts.HttpHandler.RegisterRouter(g)
	return &HttpServer{
		ginServer: s,
	}
}
