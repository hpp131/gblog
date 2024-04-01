package main

import (
	"context"
	"fmt"
	"net"
	mygrpc "skill/grpc"

	"google.golang.org/grpc"
)



type DemoServiceImpl struct {
	// inherit
	*mygrpc.UnimplementedDemoServiceServer
}

func (d *DemoServiceImpl) SayHello(ctx context.Context, in *mygrpc.DemoRequest) (out *mygrpc.DemoResponse, err error) {
	return &mygrpc.DemoResponse{Value: fmt.Sprintf("hello world %s",  in.Value)}, nil
}

func main()  {
	server := grpc.NewServer()
	// protobuf的grpc插件自动生成的register方法
	mygrpc.RegisterDemoServiceServer(server, &DemoServiceImpl{})
	// grpc依赖go内置的net包进行网络通信
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(listen)
	}
	server.Serve(listen)
	
}