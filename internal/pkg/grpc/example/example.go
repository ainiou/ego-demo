package example

import (
	"context"
	"github.com/ego-component/eetcd"
	"github.com/ego-component/eetcd/registry"
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/client/egrpc/resolver"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

type Example struct {
	client helloworld.GreeterClient
}

func NewExampleInterface() ExampleInterface {
	// 必须注册在grpc前面
	resolver.Register("etcd", registry.Load("registry").Build(registry.WithClientEtcd(eetcd.Load("etcd").Build())))
	grpcConn := egrpc.Load("grpc.example").Build()
	return &Example{
		client: helloworld.NewGreeterClient(grpcConn.ClientConn),
	}
}

// SayHello ...
func (h *Example) SayHello(ctx context.Context) (record string, err error) {
	var resp *helloworld.HelloReply
	if resp, err = h.client.SayHello(ctx, &helloworld.HelloRequest{
		Name: "ego",
	}); err != nil {
		return
	}
	record = resp.GetMessage()
	return
}
