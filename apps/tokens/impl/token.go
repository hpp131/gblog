package impl

import (
	"context"

	"github.com/hpp131/gblog/apps/tokens"
)



func (t *TokenServiceImpl) GetToken(ctx context.Context, in *tokens.GetTokenRequest) (*tokens.Token, error) {
	return nil, nil
}


func (t *TokenServiceImpl) RevokeToken(ctx context.Context, in *tokens.RevokeTokenRequest) error {
	return nil
}


func (t *TokenServiceImpl) ValidateToken(ctx context.Context, in *tokens.ValidateTokenRequest) error {
	return nil
}