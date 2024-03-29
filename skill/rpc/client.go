// package rpc

// import (
// 	"fmt"
// 	"net/rpc"
// )


// var _ RPCInterface =  (*rpcClient)(nil)

// type rpcClient struct {
// 	client *rpc.Client
// }

// func main()  {
// 	r, err := NewRpcClient("tcp", "127.0.0.1:1234")
// 	if err != nil {
// 		panic(err)
// 	}
// 	var res string
// 	r.Hello("Hello", &res)
// 	fmt.Printf("after rpc called, res is %s", res)
// }

// func NewRpcClient(net, addr string)  (*rpcClient, error) {
// 	c, err := rpc.Dial(net, addr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &rpcClient{
// 		client: c,
// 	}, nil
// }

// func (r *rpcClient) Hello(req string, res *string) error {

// 	err := r.client.Call("RPCServer.Hello", "hello", res)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

package rpc