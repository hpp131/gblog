package exception



type APIExeeption struct {
	Code    int
	Message string
}

// 实现error接口的Error方法
func (a *APIExeeption) Error() string{
	return a.Message
}