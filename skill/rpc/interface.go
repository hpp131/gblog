package rpc

type RPCInterface interface{
	Hello(req string, res *string) error
}