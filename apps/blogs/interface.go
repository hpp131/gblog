package blogs

import "context"

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	// 非全量字段更新
	PatchBlog(context.Context, *PatchBlogRequest) (*Blog, error)
	// 全量字段更新
	PutBlog(context.Context, *PutBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
}

type BlogSet struct {
	Total int
	Items []*Blog
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]interface{}{},
	}
}

type CreateBlogRequest struct {
	Title   string         `json:"title" gorm:"colume:title" validate:"required"`
	Content string         `json:"content" gorm:"colume:content" validate:"required"`
	Tags    map[string]any `json:"tags" gorm:"colume:tags;serializer:json" validate:"required"`
	Summary string         `json:"summary" gorm:"colume:summary" validate:"required"`
	Author  string
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{}
}

// 使用分页返回
type QueryBlogRequest struct {
	PageSize int
	PageNum  int
}

func NewDescribeBlogRequest(id int64) *DescribeBlogRequest {
	return &DescribeBlogRequest{Id: id}
}

type DescribeBlogRequest struct {
	Id int64
}

type PatchBlogRequest struct {
}

type PutBlogRequest struct {
}

func NewDeleteBlogRequest(id int64) *DeleteBlogRequest {
	return &DeleteBlogRequest{Id: id}
}

type DeleteBlogRequest struct {
	Id int64
}
