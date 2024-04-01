package protobufrpc

type CustomHellor interface {
	Hello(req *Request, res *Response) error
}