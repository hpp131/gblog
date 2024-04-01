package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"skill/protobufrpc"

)


type CusHelloClient struct {
	c *rpc.Client
}

var _ protobufrpc.CustomHellor = (*CusHelloClient)(nil)

func (c *CusHelloClient) Hello(req *protobufrpc.Request, res *protobufrpc.Response) error {
	c.c.Call("helloservice", req, res)
	return nil
}

func dialConn() *CusHelloClient {
	conn, err := net.Dial("tcp", "127.0.0.1")
	if err != nil {
		panic(err)
	}

	// 使用json作为序列化协议
	// 因为server端是使用json作为序列化协议，因此client端也应该使用json进行序列化
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return  &CusHelloClient{c: client}
}


func main()  {
	var response = &protobufrpc.Response{}
	client := dialConn()
	err := client.Hello(&protobufrpc.Request{Value: "client"}, response)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)	
}