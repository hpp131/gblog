package impl

import (
	"context"

	"github.com/hpp131/gblog/apps/users"
)


func (*UserServiceImple) CreateUser(context.Context, *users.CreatedUserRequest) (*users.User, error){
	return nil, nil
}

func (*UserServiceImple) QueryUser(context.Context, *users.QueryUserRequest) (*users.UserSet, error){
	return nil, nil
}

func (*UserServiceImple) DescribeUser(context.Context, *users.DescribeUserRequest) (*users.User, error){
		return nil, nil
}