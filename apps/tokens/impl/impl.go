package impl

import (
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



