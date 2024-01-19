package users_test

import (
	// "os/user"
	"fmt"
	"testing"

	"github.com/hpp131/gblog/apps/users"
)



func TestEncryptPassword(t *testing.T) {
	req := users.NewCreateUserRequest()
	req.Password = "123124"
	req.EncryptyPassword()
	t.Log(req.Password)
	err := req.CheckPassword("123124")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("test")
	
}