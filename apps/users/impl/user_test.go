package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hpp131/gblog/apps/users"
	"github.com/hpp131/gblog/apps/users/impl"
)



var (
	i  = impl.NewUserServiceImple()
	ctx =  context.Background()
)

func TestCreateUser(t *testing.T){
	req := users.NewCreateUserRequest()
	req.Username = "admin1"
	req.Password = "123123"
	req.Role = users.ADMIN
	user, err := i.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal("user is nil")
	}
	fmt.Printf("%v", user)
}


func TestQueryUser(t *testing.T){
	req := users.NewQueryUserRequest()
	req.Username = "admin8"
	// req.PageNum = 1
	// req.PageSize = 5
	result, err := i.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}


func TestDescribeUser(t *testing.T){
	req := users.NewDescribeUserRequest()
	req.UserId = 2
	res, err := i.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
