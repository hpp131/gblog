package exception

type APIException struct {
	Code    int
	Message string
}

func NewAPIException(code int, message string) *APIException {
	return &APIException{
		Code:    code,
		Message: message,
	}
}

// 实现error接口的Error方法
func (a *APIException) Error() string {
	return a.Message
}

// 链式编程，函数返回值是仍然是*APIException类型
func (a *APIException) WithCode(code int) *APIException {
	a.Code = code
	return a
}

func (a *APIException) WithMessage(message string) *APIException {
	a.Message = message
	return a
}

// 判断error是否是自定义错误类型APIException下的一种(根据code判断)。具体的APIException类型定义在各模块下，如tokens.excepion.go
func IsAPIException(err error, a *APIException) bool {
	if exception, ok := err.(*APIException); ok {
		return exception.Code == a.Code
	}
	return false
}
