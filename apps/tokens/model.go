package tokens

import (
	"fmt"
	"time"

	"github.com/infraboard/mcube/tools/pretty"
	"github.com/rs/xid"
)

type Token struct {
	AccessToken           string
	AccessTokenExpiresAt  int64
	RefreshToken          string
	RefreshTokenExpiresAt int64
	CreatedAt             int64
	UpdatedAt             int64
	UserId                int64
	Username              string
}


func NewToken() *Token {
	return &Token{
		AccessTokenExpiresAt: 7200, // 2 hoursa
		RefreshTokenExpiresAt: 86400, // 1 day
		CreatedAt: time.Now().Unix(),
		// 暂用uuid作为token
		AccessToken: xid.New().String(),
		RefreshToken: xid.New().String(),
	}
}


func (t *Token)	Validate() error {
	if (t.CreatedAt + t.RefreshTokenExpiresAt < time.Now().Unix()) {
		return fmt.Errorf("RefreshToken is expired")
	}
	if (t.CreatedAt + t.AccessTokenExpiresAt < time.Now().Unix()){
		return fmt.Errorf("AccessToken is expired")
	}
	return nil
}

func (t *Token) String() string {
	return pretty.ToJSON(t)
}

func (t *Token) TableName() string {
	return "tokens"
}

