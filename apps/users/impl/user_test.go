package impl_test

import (
	"context"
	"fmt"
	// "fmt"
	"testing"

	"github.com/hpp131/gblog/apps/users"
	"github.com/hpp131/gblog/apps/users/impl"
	// "github.com/hpp131/gblog/apps/users/impl"
)



var (
	i  = impl.NewUserServiceImple()
	ctx =  context.Background()
)

func TestCreateUser(t *testing.T){
	req := users.NewCreateUserRequest()
	req.Username = "admin9"
	req.Password = "123123"
	req.Role = users.ADMIN
	user, err := i.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal("user is nil")
	}
	// t.Log(user)
	// fmt.Println("hello world2")

	fmt.Printf("%v", user)
}