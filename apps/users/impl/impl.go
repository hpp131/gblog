package impl

import (
	"github.com/hpp131/gblog/conf"
	"gorm.io/gorm"
)

// 单独使用一个文件定义一个实现类来实现Service interface
// 具体方法的实现位于./user.go
type UserServiceImple struct {
	db *gorm.DB
}




func NewUserServiceImple() *UserServiceImple{
	return &UserServiceImple{
		db: conf.C().DB(),
	}
}