package users

import "context"

// 接口定义, 一定要考虑兼容性, 接口的参数不能变
// CreateUserRequest,QueryUserRequest,DescribeUserRequest可以修改，但是修改后不会影响接口方法以及不必改写实现类的入参和出参（返回值）
type Service interface {
	CreatedUser(context.Context, *CreatedUserRequest) (*User, error)
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

type CreatedUserRequest struct {
	// "创建时间" 字段使用int64类型而不使用time.Time类型，是因为time.Time类型在序列化时，会包含时区信息（不方便代码全球通用），而使用int64代表的timestamp（全球范围内是绝对的）则不包含时区信息，
	// CreatedAt int64
	Username string
	Password string
	Label    map[string]string
	Role     string
}

// 查询用户列表
type QueryUserRequest struct {
	PageSize int
	PageNum  int
	// 可以根据指定用户名进行查询
	Username string
}

type DescribeUserRequest struct {
	UserId int64
}
