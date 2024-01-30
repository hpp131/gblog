package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/apps/blogs"
	"github.com/hpp131/gblog/ioc"
)

func init() {
	ioc.API().Registry(blogs.AppName, &BlogAPIHandler{})
}

type BlogAPIHandler struct {
	svc blogs.Service
}

func (b *BlogAPIHandler) Registry(rr gin.IRouter) {
	blogRouter := rr.Group(blogs.AppName)
	// createBlog
	blogRouter.POST("/", b.createBlog)
	// queryBlog
	blogRouter.POST("/", b.queryBlog)
	// describBlog
	blogRouter.POST("/", b.describeBlog)
	// patchBlog
	blogRouter.POST("/", b.patchBlog)
	// putBlog
	blogRouter.POST("/", b.putBlog)
	// deleteBlog
	blogRouter.POST("/", b.deleteBlog)
}

// 实现ioc.Objector interface
func (b *BlogAPIHandler) Init() error {
	obj, err := ioc.Controller().Get(blogs.AppName)
	if err != nil {
		return err
	}
	if value, ok := obj.(blogs.Service); ok {
		b.svc = value
	}
	return nil
}

func (b *BlogAPIHandler) Destroy() error {
	return nil
}

func (b *BlogAPIHandler) createBlog(*gin.Context) {

}

func (b *BlogAPIHandler) queryBlog(*gin.Context) {

}

func (b *BlogAPIHandler) describeBlog(*gin.Context) {

}
func (b *BlogAPIHandler) patchBlog(*gin.Context) {

}

func (b *BlogAPIHandler) putBlog(*gin.Context) {

}

func (b *BlogAPIHandler) deleteBlog(*gin.Context) {

}
