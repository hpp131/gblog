package tokens

import "github.com/hpp131/gblog/exception"

// 自定义异常通常使用"Err"开头，方便使用时进行索引/搜索
var (
	ErrAccessTokenExpired  = exception.NewAPIException(1000, "access token expired")
	ErrRefreshTokenExpired = exception.NewAPIException(1001, "refresh token invalid")
	ErrTokenNotFound = exception.NewAPIException(1002, "Token is not found in request")
)
