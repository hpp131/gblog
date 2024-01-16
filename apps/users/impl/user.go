package impl

import (
	"context"

	"github.com/hpp131/gblog/apps/users"
)




func (u *UserServiceImple) CreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.User, error){
	u.db = u.db.WithContext(ctx).Table("users")
	// 创建用户的时候先查询一下数据库中有没有同名的用户
	// 校验入参格式和类型的合法性

	err := in.Validate()
	if err != nil {
		return nil, err
	}
	user := users.NewUser(in)
	record := u.db.Create(&user)
	if record.Error != nil {
		return nil, record.Error
	}
	return user, nil
}


func (u *UserServiceImple) QueryUser(ctx context.Context, in *users.QueryUserRequest) (*users.UserSet, error){
	return nil, nil
}

func (u *UserServiceImple) DescribeUser(ctx context.Context, in *users.DescribeUserRequest) (*users.User, error){
		return nil, nil
}