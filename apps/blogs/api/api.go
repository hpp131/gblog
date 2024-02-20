package api

import (
	// "fmt"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/apps/blogs"
	"github.com/hpp131/gblog/apps/users"
	"github.com/hpp131/gblog/ioc"
	"github.com/hpp131/gblog/middleware"
	"github.com/hpp131/gblog/response"
)

func init() {
	ioc.API().Registry(blogs.AppName, &BlogAPIHandler{})
}

type BlogAPIHandler struct {
	svc blogs.Service
}

func (b *BlogAPIHandler) Registry(rr gin.IRouter) {
	blogRouter := rr.Group(blogs.AppName)
	// queryBlog
	blogRouter.GET("/", b.queryBlog)
	// describBlog
	blogRouter.GET("/:id", b.describeBlog)

	// 对于增删写入api,在操作前需要先通过中间件的认证
	blogRouter.Use(middleware.Auth)
	// createBlog
	blogRouter.POST("/", b.createBlog)
	// // patchBlog
	// blogRouter.PATCH("/", b.patchBlog)
	// putBlog
	blogRouter.PUT("/:id", b.putBlog)
	// deleteBlog
	blogRouter.DELETE("/:id", middleware.Required(users.ADMIN), b.deleteBlog)
}

// 实现ioc.Objector interface
func (b *BlogAPIHandler) Init() error {
	obj := ioc.Controller().Get(blogs.AppName)
	if value, ok := obj.(blogs.Service); ok {
		b.svc = value
	}
	return nil
}

func (b *BlogAPIHandler) Destroy() error {
	return nil
}

func (b *BlogAPIHandler) createBlog(c *gin.Context) {
	req := blogs.NewCreateBlogRequest()
	if err := c.BindJSON(req); err != nil {
		response.Failed(c, err)
	}
	fmt.Println(req)
	bl, err := b.svc.CreateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
	}
	response.Success(c, bl)

}

func (b *BlogAPIHandler) queryBlog(c *gin.Context) {
	var err error
	req := blogs.NewQueryBlogRequest()
	req.PageNum, err = strconv.Atoi(c.Query("pagenum"))
	if err != nil {
		response.Failed(c, ErrQueryParam)
	}
	blogSet, _ := b.svc.QueryBlog(c, req)
	response.Success(c, blogSet)
}

func (b *BlogAPIHandler) describeBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	req := blogs.NewDescribeBlogRequest(int64(id))
	if bl, err := b.svc.DescribeBlog(c, req); err != nil {
		response.Failed(c, err)
	} else {
		response.Success(c, bl)
	}
}

// func (b *BlogAPIHandler) patchBlog(*gin.Context) {
// 	response.Success(c, )

// }

func (b *BlogAPIHandler) putBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	req := blogs.NewPutBlogRequest((int64(id)))
	if bl, err := b.svc.PutBlog(c, req); err != nil {
		response.Failed(c, err)
	} else {
		response.Success(c, bl)
	}
}

func (b *BlogAPIHandler) deleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	req := blogs.NewDeleteBlogRequest(int64(id))
	if bl, err := b.svc.DeleteBlog(c, req); err != nil {
		response.Failed(c, err)
	} else {
		response.Success(c, bl)
	}
}
