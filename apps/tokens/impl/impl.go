package impl

import (
	"github.com/hpp131/gblog/apps/tokens"
	"github.com/hpp131/gblog/conf"
	"gorm.io/gorm"
)

// 定义tokens.Service interface的实现类
type TokenServiceImpl struct {
	db 	*gorm.DB
}


func NewTokenServiceImpl(db *gorm.DB) *TokenServiceImpl {
	return &TokenServiceImpl{db: conf.C().DB()}
}



func (t *TokenServiceImpl) GetToken(token string, userID uint64) (*tokens.Token, error) {
	return nil, nil
}


func (t *TokenServiceImpl) RevokeToken(token string, userID uint64) error {
	return nil
}


func (t *TokenServiceImpl) ValidateToken(token string, userID uint64) error {
	return nil
}