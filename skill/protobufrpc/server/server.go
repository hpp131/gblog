// 这里的server.go和client.go只是使用protobuf的代码生成功能自动生成了Request和Response结构体
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"skill/protobufrpc"
)

type CusHelloImpl struct{}

// 在程序编译时检查CusHelloImpl是否实现了CustomHellor接口
var _ protobufrpc.CustomHellor = (*CusHelloImpl)(nil)

func (c *CusHelloImpl) Hello(req *protobufrpc.Request, res *protobufrpc.Response) error {
	res.Value = fmt.Sprintf("hello, %s", req.Value)
	return nil
}

func main() {
	rpc.RegisterName("helloservice", new(CusHelloImpl))
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		// rpc标准库默认使用gob编码，要使用自定义编码需要调用ServeCodec方法传入自定义编解码工具
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
