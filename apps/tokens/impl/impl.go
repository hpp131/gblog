package impl

import (
	"context"
	"fmt"

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
	obj := ioc.Controller().Get(users.AppName)
	t.user = obj.(users.Service)
	t.db = conf.C().DB()
	return nil
}

func (t *TokenServiceImpl) Destroy() error {
	return nil
}


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
	err = t.db.Create(tk).Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}


func (t *TokenServiceImpl) RevokeToken(ctx context.Context, in *tokens.RevokeTokenRequest) (error) {
	// 从库里面删除某条token记录
	err := t.db.Where("access_token = ?", in.AccessToken).Delete(&tokens.Token{}).Error
	if err != nil {
		return err
	}
	return nil
}


func (t *TokenServiceImpl) ValidateToken(ctx context.Context, in *tokens.ValidateTokenRequest) (*tokens.Token, error) {
	// 校验token是否在过期时间内/是否是刷新令牌/是否已经注销
	tk := &tokens.Token{}
	err := t.db.Model(&tokens.Token{}).Where("access_token = ? and refresh_token = ?", in.AccessToken, in.RefreshToken).First(&tk).Error
	if err != nil {
		return nil, err
	}
	if err := tk.Validate();err != nil {
		return nil, err
	}
	return tk, nil
}
