package tokens

import (
	"time"

	// "github.com/hpp131/gblog/apps/tokens/impl"
	"github.com/hpp131/gblog/apps/users"
	"github.com/infraboard/mcube/tools/pretty"
	"github.com/rs/xid"
)

type Token struct {
	AccessToken           string
	AccessTokenExpiredAt  int64
	RefreshToken          string
	RefreshTokenExpiredAt int64
	CreatedAt             int64
	UpdatedAt             int64
	UserId                int64
	Username              string
	// 用于RBAC中间件,middleware.Required()
	Role                  users.Role
}

func NewToken() *Token {
	return &Token{
		AccessTokenExpiredAt:  7200,  // 2 hoursa
		RefreshTokenExpiredAt: 86400, // 1 day
		CreatedAt:             time.Now().Unix(),
		// 暂用uuid作为token
		AccessToken:  xid.New().String(),
		RefreshToken: xid.New().String(),
	}
}

func (t *Token) Validate() error {
	if t.CreatedAt+t.RefreshTokenExpiredAt < time.Now().Unix() {
		return ErrAccessTokenExpired
	}
	if t.CreatedAt+t.AccessTokenExpiredAt < time.Now().Unix() {
		return ErrRefreshTokenExpired
	}
	return nil
}

func (t *Token) String() string {
	return pretty.ToJSON(t)
}

func (t *Token) TableName() string {
	return "tokens"
}
