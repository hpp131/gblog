package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/apps/tokens"
	"github.com/hpp131/gblog/response"
	"github.com/hpp131/gblog/ioc"
)

func init(){
	ioc.API().Registry(tokens.AppName, &TokenAPIHandler{})
}

type TokenAPIHandler struct {
	svc tokens.Service
}

func NewTokenAPIHandler(svc tokens.Service) *TokenAPIHandler {
	return &TokenAPIHandler{svc}
}

// 注册token模块内所有api的路由路径
func (t *TokenAPIHandler) Registry(r gin.IRouter) {
	moduleRouter := r.Group(tokens.AppName)
	moduleRouter.POST("/", t.Login)
	moduleRouter.DELETE("/", t.LogOut)
}

// 该func供登录api调用
func (t *TokenAPIHandler) Login(c *gin.Context) {
	req := tokens.NewIssueTokenRequest("", "")
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(401, gin.H{"error": err})
	}
	tk, err := t.svc.IssueToken(c, req)
	// 设置token到cookie中
	c.SetCookie("token", tk.AccessToken, int(tk.AccessTokenExpiredAt), "/", "127.0.0.1", false, true)

	if err != nil {
		// 自定义RestfulAPIResponse
		response.Failed(c, err)
	} else {
		response.Success(c, tk)
	}


}

// 该func供退出api调用
func (t *TokenAPIHandler) LogOut(c *gin.Context) {
	// 获取请求头中的Authorization字段值，如果没有就去cookie中获取
	var req *tokens.RevokeTokenRequest

	at := GetAccessTokenFromHttp(c.Request)
	req = tokens.NewRevokeTokenRequest(at)
	err := t.svc.RevokeToken(c, req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 删除客户端的token
	c.SetCookie("token", "", -1, "/", "127.0.0.1", false, true)

	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "退出成功")
}


// 分别尝试从Authorization和cookie字段中获取accessToken
func GetAccessTokenFromHttp(r *http.Request) string {
	at := r.Header.Get("Authorization")
	if at != "" {
		return at
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// 实现ioc.Objector interface: 从ioc容器中获取自己所需要的依赖
func (t *TokenAPIHandler) Init() error{
	obj := ioc.Controller().Get(tokens.AppName)
	t.svc = obj.(tokens.Service)
	return nil
}

func (t *TokenAPIHandler) Destroy() error{
	return nil
}