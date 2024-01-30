package impl

import (
	"context"

	"github.com/hpp131/gblog/apps/blogs"
	"github.com/hpp131/gblog/ioc"
	"gorm.io/gorm"
)

func init(){
	ioc.Controller().Registry(blogs.AppName, &BlogServiceImpl{})
}


type BlogServiceImpl struct{
	db  *gorm.DB
}

func (b *BlogServiceImpl) CreateBlog(ctx context.Context, in *blogs.DescribeBlogRequest) (*blogs.Blog, error){
	return nil, nil
}

func (b *BlogServiceImpl) QueryBlog(ctx context.Context, in *blogs.DescribeBlogRequest) (*blogs.Blog, error){
	return nil, nil

}

func (b *BlogServiceImpl) DescribeBlog(ctx context.Context, in *blogs.DescribeBlogRequest) (*blogs.Blog, error){
	return nil, nil

}

func (b *BlogServiceImpl) PatchBlog(ctx context.Context, in *blogs.DescribeBlogRequest) (*blogs.Blog, error){
	return nil, nil

}

func (b *BlogServiceImpl) PutBlog(ctx context.Context, in *blogs.DescribeBlogRequest) (*blogs.Blog, error){
	return nil, nil

}

func (b *BlogServiceImpl) DeleteBlog(ctx context.Context, in *blogs.DescribeBlogRequest) (*blogs.Blog, error){
	return nil, nil

}


// 从ioc容器解决自身依赖

func (b *BlogServiceImpl) Init() error{
	return nil
}

func (b *BlogServiceImpl) Destroy() error{
	return nil
}


