package impl

import (
	"github.com/hpp131/gblog/apps/tokens"
	"github.com/hpp131/gblog/apps/users"
	"github.com/hpp131/gblog/apps/users/impl"
	"github.com/hpp131/gblog/conf"
	"github.com/hpp131/gblog/ioc"
	"gorm.io/gorm"
)

// 向ioc容器注册TokenServiceImpl的实例对象
func init() {
	ioc.Controller().Registry(tokens.AppName, &TokenServiceImpl{})
}

// 实体的构造函数一般写在对应实体的上面
func NewTokenServiceImpl() *TokenServiceImpl {
	return &TokenServiceImpl{db: conf.C().DB(),
		user: impl.NewUserServiceImple()}
}

// 定义tokens.Service interface的实现类
type TokenServiceImpl struct {
	db *gorm.DB
	// token模块依赖于user模块，但是不直接传入user模块的实现类，而是user的interface。先定义一个螺丝位置，然后在后期将制指定品牌的螺丝放进来，体现的是面向接口编程
	user users.Service
}

// 实现ioc.Objector interface
func (t *TokenServiceImpl) Init() error {
	if obj, err := ioc.Controller().Get(users.AppName); err != nil {
		return err
	}else{
		t.user = obj.(users.Service)
		t.db = conf.C().DB()
	}
	return nil
}

func (t *TokenServiceImpl) Destroy() error {
	return nil
}
