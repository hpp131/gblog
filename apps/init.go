package apps

// 执行各个包下面的init()，将需要的对象注册到IOC容器中
import(
	_ "github.com/hpp131/gblog/apps/users/impl"
	_ "github.com/hpp131/gblog/apps/tokens/api"
	_ "github.com/hpp131/gblog/apps/tokens/impl"
	_ "github.com/hpp131/gblog/apps/blogs/impl"
	_ "github.com/hpp131/gblog/apps/blogs/api"
)