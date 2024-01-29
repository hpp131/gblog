package ioc

//ioc container

func NewContainer() *Container {
	return &Container{storage: make(map[string]Objector)}
}

// ioc container
type Container struct {
	storage map[string]Objector
}

// 需要为ioc对象提供注册方法。该对象所在的go文件中需要有init()方法来把该对象注册到ioc容器中。例如注册tokens.TokenAPIHander对象时，需要在tokens模块下api.go文件中添加init()方法
func (c *Container) Registry(name string, value Objector) {
	c.storage[name] = value
}

func (c *Container) Get(name string) (Objector, error) {
	if v, ok := c.storage[name]; ok {
		return v, nil
	}
	return nil, nil
}
