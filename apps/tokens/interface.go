package tokens

import (
	"context"
)

type Service interface {
	// 用户登录后，需要颁发token给用户
	IssueToken(ctx context.Context, in *IssueTokenRequest) (*Token, error)
	// 用户注销后，需要把token给吊销掉
	RevokeToken(ctx context.Context, in *RevokeTokenRequest) error
	// 验证用户每次请求中所携带的token
	ValidateToken(ctx context.Context, in *ValidateTokenRequest) (*Token, error)
}

type IssueTokenRequest struct {
	// 首先搞清楚什么时候为用户颁发token，应该是用户首次登录的时候把？
	// 为用户颁发token前得需要校验用户密码吧？校验成功了就颁发token
	Username string
	Password string
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{Username: username, Password: password}
}


type RevokeTokenRequest struct {
	// 从库里面删除某条token记录
	AccessToken  string
	// RefreshToken string
}

func NewRevokeTokenRequest(accessToken string) *RevokeTokenRequest {
	return &RevokeTokenRequest{AccessToken: accessToken}
}

type ValidateTokenRequest struct {
	// 校验token是否在过期时间内/是否是刷新令牌/是否已经注销
	AccessToken        string
	// RefreshToken string
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{AccessToken: accessToken}	
}
