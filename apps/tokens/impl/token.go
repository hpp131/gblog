package impl

import (
	"context"
	"fmt"

	"github.com/hpp131/gblog/apps/tokens"
	"github.com/hpp131/gblog/apps/users"
)



func (t *TokenServiceImpl) IssueToken(ctx context.Context, in *tokens.GetTokenRequest) (*tokens.Token, error) {
	user := users.NewQueryUserRequest()
	user.Username = in.Username
	userSet, err := t.user.QueryUser(ctx, user)
	if err != nil {
		return nil, err
	}
	if len(userSet.Items) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	
	return nil, nil
}


func (t *TokenServiceImpl) RevokeToken(ctx context.Context, in *tokens.RevokeTokenRequest) error {
	return nil
}


func (t *TokenServiceImpl) ValidateToken(ctx context.Context, in *tokens.ValidateTokenRequest) error {
	return nil
}