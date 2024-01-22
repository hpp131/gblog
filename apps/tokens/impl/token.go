package impl

import (
	"context"
	"fmt"

	"github.com/hpp131/gblog/apps/tokens"
	"github.com/hpp131/gblog/apps/users"
)



func (t *TokenServiceImpl) IssueToken(ctx context.Context, in *tokens.IssueTokenRequest) (*tokens.Token, error) {
	user := users.NewQueryUserRequest()
	user.Username = in.Username
	userSet, err := t.user.QueryUser(ctx, user)
	if err != nil {
		return nil, err
	}
	if len(userSet.Items) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	// 判断用户存在后校验其密码
	if err := userSet.Items[0].CheckPassword(in.Password);err != nil {
		return nil, err
	}
	// 密码校验通过后，生成token
	tk := tokens.NewToken()
	tk.UserId = userSet.Items[0].Id
	tk.Username = in.Username
	return tk, nil
}


func (t *TokenServiceImpl) RevokeToken(ctx context.Context, in *tokens.RevokeTokenRequest) error {
	// 从库里面删除某条token记录
	t.db.Model(&tokens.Token{}).Where("access_token = ?", in.AccessToken).Delete(&tokens.Token{})
	return nil
}


func (t *TokenServiceImpl) ValidateToken(ctx context.Context, in *tokens.ValidateTokenRequest) error {
	// 校验token是否在过期时间内/是否是刷新令牌/是否已经注销
	tk := &tokens.Token{}
	err := t.db.Model(&tokens.Token{}).Find("access_token = ?, ", in.AccessToken).First(&tk).Error
	if err != nil {
		return err
	}
	if err := tk.Validate();err != nil {
		return err
	}
	return nil
}