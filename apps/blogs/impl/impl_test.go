package impl_test

import (
	"context"
	"testing"

	"github.com/hpp131/gblog/apps/blogs"
	// "github.com/hpp131/gblog/apps/blogs/impl"
	// "github.com/hpp131/gblog/conf"
	"github.com/hpp131/gblog/ioc"
	// "gorm.io/gorm"
)


var (
	ctx       = context.Background()
	blImplObj blogs.Service
)

// 使用ioc的方式解决依赖；使用ioc重写单元测试
func init() {
	blImplObj = ioc.Controller().Get(blogs.AppName).(blogs.Service)
	if err := ioc.Init(); err != nil {
		panic(err)
	}
}


func TestCreateBlog(t *testing.T){
	// obj, _ := ioc.Controller().Get(blogs.AppName)
	// if i, ok := obj.(blogs.Service);ok{
	// 	imple = i
	// }
	req := blogs.NewCreateBlogRequest()
	req.Author = "lin"
	req.Content = "this is content"
	req.Title = "unit_test"
	req.Tags["lang"] = "golang"
	req.Summary = "a summary"
	bl, err := blImplObj.CreateBlog(ctx, req);
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bl)
	
	
}

func NewCreateBlogRequest() {
}

func TestQueryBlog(*testing.T){
	
}

func TestPatchBlog(*testing.T){
	
}


func TestPutBlog(*testing.T){
	
}


func TestDeleteBlog(*testing.T){
	
}