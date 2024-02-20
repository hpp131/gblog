package middleware

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/apps/tokens"
	"github.com/hpp131/gblog/apps/tokens/api"
	"github.com/hpp131/gblog/ioc"

	// "github.com/hpp131/gblog/exception"
	"github.com/hpp131/gblog/response"
)

// 对需要携带认证信息才能访问的请求生效,如果api本身就是公开的，那么不需要用到该中间件
func Auth(c *gin.Context) {
	tkSVC := ioc.Controller().Get(tokens.AppName).(tokens.Service)
	ak := api.GetAccessTokenFromHttp(c.Request)
	if ak == "" {
		response.Failed(c, tokens.ErrTokenNotFound)
		return
	}
	// 如果能获取到ak,则对ak进行validate,判断token是否过期
	token, err := tkSVC.ValidateToken(c.Request.Context(), tokens.NewValidateTokenRequest(ak))
	if err != nil {
		response.Failed(c, tokens.ErrAccessTokenExpired)
	}
	c.Set(tokens.MiddleWareKey, token)
}
