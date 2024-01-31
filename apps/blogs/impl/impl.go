package impl

import (
	"context"

	// "github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/apps/blogs"
	"github.com/hpp131/gblog/conf"
	"github.com/hpp131/gblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registry(blogs.AppName, &BlogServiceImpl{})
}

type BlogServiceImpl struct {
	db *gorm.DB
}

func (b *BlogServiceImpl) CreateBlog(ctx context.Context, in *blogs.CreateBlogRequest) (*blogs.Blog, error) {
	blog := blogs.NewBlog(in)
	if err := b.db.Create(&blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func (b *BlogServiceImpl) QueryBlog(ctx context.Context, in *blogs.QueryBlogRequest) (*blogs.BlogSet, error) {
	blogSet := blogs.NewBlogSet()
	if err := b.db.Offset(in.Offset()).Limit(in.Limit()).Find(blogSet.Items).Error; err != nil {
		return nil, nil
	}
	blogSet.Total = len(blogSet.Items)
	return blogSet, nil

}

func (b *BlogServiceImpl) DescribeBlog(ctx context.Context, in *blogs.DescribeBlogRequest) (bl *blogs.Blog, err error) {
	// var blog  *blogs.Blog
	if err := b.db.Where("id = ?", in.Id).First(&bl).Error; err != nil {
		return nil, err
	}
	return bl, nil
}

func (b *BlogServiceImpl) PutBlog(ctx context.Context, in *blogs.PutBlogRequest) (bl *blogs.Blog, err error) {
	bl.Id = in.Id
	bl.UpdatedAt = in.UpdatedAt
	bl.CreateBlogRequest = in.CreateBlogRequest
	if err := b.db.Model(bl).Updates(blogs.Blog{UpdatedAt: in.UpdatedAt, CreateBlogRequest: in.CreateBlogRequest}).Error; err != nil{
		return nil, err
	}
	return bl, nil
}

// func (b *BlogServiceImpl) PutBlog(ctx context.Context, in *blogs.PutBlogRequest) (bl *blogs.Blog, err error) {
// 	// bl.Id = in.Id
// 	// bl.UpdatedAt = in.UpdatedAt
// 	// bl.CreateBlogRequest = in.CreateBlogRequest
// 	// b.db.Save(&bl)
// 	// return bl, nil
// 	return nil, nil
// }

func (b *BlogServiceImpl) DeleteBlog(ctx context.Context, in *blogs.DeleteBlogRequest) (bl *blogs.Blog, err error) {
	if err := b.db.Where("id = ?", in.Id).Delete(&bl).Error; err != nil {
		return nil, err
	}
	return bl, nil
}

// 从ioc容器中获取自身需要的依赖
func (b *BlogServiceImpl) Init() error {
	b.db = conf.C().DB()
	return nil
}

func (b *BlogServiceImpl) Destroy() error {
	return nil
}
