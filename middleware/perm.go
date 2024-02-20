package middleware

// 使用中间件来实现简易的RBAC功能，复杂的RBAC功能如何实现？

import (
	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/apps/tokens"
	"github.com/hpp131/gblog/apps/users"
	"github.com/hpp131/gblog/response"
)

func Required(role ...users.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		if value, ok := c.Get(tokens.MiddleWareKey);ok{
			tkRole := value.(*tokens.Token)
			for _, r := range role{
				if tkRole.Role == r{
					return
				}
			}
			// 遍历完成后发现token中的role信息，与api允许的Role不匹配，直接返回permission denied
			response.Failed(c, tokens.ErrPermissDenied)
		}
	}
}
