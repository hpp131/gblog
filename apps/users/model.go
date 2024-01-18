package users

import (
	"time"
	"github.com/infraboard/mcube/tools/pretty"
)

// 需要入库的对象（PO）
type User struct {
	Id       int64 `json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64	`json:"updated_at"`
	// 利用struct嵌套，将CreateUserRequest中的字段直接映射到User中，不必再一一枚举
	*CreateUserRequest
}

func NewUser(c *CreateUserRequest) *User {
	return &User{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		CreateUserRequest: c,
	}
}

// 实现stringer接口，方便使用fmt或其他类似包时对对象进行格式化输出
func (u *User) String() string {
	return pretty.ToJSON(u)
}

func (u *User) TableName() string {
	return "users"
}