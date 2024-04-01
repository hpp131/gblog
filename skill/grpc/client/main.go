package main

import (
	"context"
	"fmt"
	mygrpc "skill/grpc"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

func dialConn() mygrpc.DemoServiceClient {
	// 使用grpc包下封装的dial方法，直接返回一个
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	rpcClient := mygrpc.NewDemoServiceClient(conn)
	return rpcClient

}

func main()  {
	client := dialConn()
	res, err := client.SayHello(context.Background(), &mygrpc.DemoRequest{Value: "I'm client"})
	if err != nil {
		fmt.Println(err)	
	}
	fmt.Println(res)
}