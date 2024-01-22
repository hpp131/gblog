package impl

import (
	"github.com/hpp131/gblog/apps/users"
	"github.com/hpp131/gblog/conf"
	"gorm.io/gorm"
)

// 定义tokens.Service interface的实现类
type TokenServiceImpl struct {
	db *gorm.DB
	// token模块依赖于user模块，但是不直接传入user模块的实现类，而是user的interface。先定义一个螺丝位置，然后在后期将制指定品牌的螺丝放进来，体现的是面向接口编程
	user users.Service
}

func NewTokenServiceImpl(db *gorm.DB) *TokenServiceImpl {
	return &TokenServiceImpl{db: conf.C().DB()}
}
