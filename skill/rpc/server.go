package rpc

import (
	"net"
	"net/rpc"
	// "test/service"
)

// var _ service.RPCInterface = (*rpcServer)(nil)
// var _ service.RPCInterface

type RPCServer struct {}

func (r *RPCServer) Hello (req string, res *string) error {
	*res = req + " test..."
	return nil
}

func main()  {
	rpc.Register(new(RPCServer))
	// rpc.RegisterName("helloFunc", new(rpcServer))
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := listen.Accept()
		go rpc.ServeConn(conn)
	}
	
}