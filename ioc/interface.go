package ioc

// ioc对象应该实现Init功能和Destroy功能
// Init: 该方法会引入当前对象的下级依赖对象，如TokenAPIHandler.Init()会补充TokenAPIHandler对象内部的svc
// Destroy:
type Objector interface {
	Init() error
	Destroy() error
}