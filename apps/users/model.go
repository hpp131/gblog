package users

// 需要入库的对象（PO）
type User struct {
	Id       int64
	CreateAt int64
	UpdateAt int64
	// 利用struct嵌套，将CreateUserRequest中的字段直接映射到User中，不必再一一枚举
	*CreatedUserRequest
}
