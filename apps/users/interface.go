package users

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/tools/pretty"
)

type Role int

const (
	ADMIN  Role = iota
	MEMBER      // 普通用户
	GUEST
)

var v = validator.New()

// 接口定义, 一定要考虑兼容性, 接口的参数不能变
// CreateUserRequest,QueryUserRequest,DescribeUserRequest可以修改，但是修改后不会影响接口方法以及不必改写实现类的入参和出参（返回值）
type Service interface {
	CreatedUser(context.Context, *CreateUserRequest) (*User, error)
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)

	// ToDo: 修改和删除用户信息的方法
	// ModifyUser(context.Context, *ModifyUserRequest) (*User, error)
	// DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
}

type UserSet struct {
	Items  []*User
	Totals int64
}

// 实现stringer接口，使用单元测试t.Log()或fmt.xxx()函数时起作用
func (s *UserSet) string() string {
	return pretty.ToJSON(s)
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Label: map[string]string{},
	}
}

type CreateUserRequest struct {
	// "创建时间" 字段使用int64类型而不使用time.Time类型，是因为time.Time类型在序列化时，会包含时区信息（不方便代码全球通用），而使用int64代表的timestamp（全球范围内是绝对的）则不包含时区信息，
	// CreatedAt int64
	Username string            `json:"username"`
	Password string            `json:"password"`
	Label    map[string]string `json:"label" gorm:"column:label;serializer:json"`
	Role     Role              `json:"role"`
}

func (c *CreateUserRequest) Validate() error {
	// 如果Label为nil map，则进行初始化
	// if c.Label == nil {
	// 	c.Label = make(map[string]string)
	// }
	err := v.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

// 查询用户列表
type QueryUserRequest struct {
	PageSize int
	PageNum  int
	// 可以根据指定用户名进行查询
	Username string
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{}
}

func (q *QueryUserRequest) Limit() int {
	return q.PageSize
}

func (q *QueryUserRequest) Offset() int {
	return (q.PageNum - 1) * q.PageSize
}

type DescribeUserRequest struct {
	UserId int64
}

func NewDescribeUserRequest() *DescribeUserRequest {
	return &DescribeUserRequest{}
}
