package impl

import (
	"github.com/hpp131/gblog/apps/users"
	"github.com/hpp131/gblog/conf"
	"github.com/hpp131/gblog/ioc"
	"gorm.io/gorm"
)

func init(){
	ioc.Controller().Registry(users.AppName, &UserServiceImple{})
}



func NewUserServiceImple() *UserServiceImple{
	return &UserServiceImple{
		db: conf.C().DB(),
	}
}

// 单独使用一个文件定义一个实现类来实现Service interface
// 具体方法的实现位于./user.go
type UserServiceImple struct {
	db *gorm.DB
}


// 实现ioc.Objector interface
func (u *UserServiceImple) Init() error {
	u.db = conf.C().DB()
	return nil
}

func (u *UserServiceImple) Destroy() error{
	return nil
}




