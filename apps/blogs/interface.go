package blogs

import (
	"context"
	"time"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	// 非全量字段更新
	// PatchBlog(context.Context, *PatchBlogRequest) (*Blog, error)
	// 全量字段更新
	PutBlog(context.Context, *PutBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	Total int
	Items []*Blog
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

type CreateBlogRequest struct {
	Title   string         `json:"title" gorm:"colume:title"`
	Content string         `json:"content" gorm:"colume:content"`
	Tags    map[string]string `json:"tags" gorm:"colume:tags;serializer:json"`
	Summary string         `json:"summary" gorm:"colume:summary"`
	Author  string         `json:"author" gorm:"colume:author"`
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize: 20,
		PageNum:  1,
	}
}

// 使用分页返回
type QueryBlogRequest struct {
	PageSize int
	PageNum  int
}

func (q *QueryBlogRequest) Offset() int {
	return q.PageSize * (q.PageNum - 1)
}

func (q *QueryBlogRequest) Limit() int {
	return q.PageSize
}

func NewDescribeBlogRequest(id int64) *DescribeBlogRequest {
	return &DescribeBlogRequest{Id: id}
}

type DescribeBlogRequest struct {
	Id int64
}

// func NewPatchBlogRequest(id int64) *PatchBlogRequest {
// 	return &PatchBlogRequest{
// 		Id: id,
// 		UpdatedAt:         time.Now().Unix(),
// 		CreateBlogRequest: NewCreateBlogRequest(),
// 	}
// }

// type PatchBlogRequest struct {
// 	Id        int64
// 	UpdatedAt int64
// 	*CreateBlogRequest
// }

func NewPutBlogRequest(id int64) *PutBlogRequest {
	return &PutBlogRequest{
		Id: id,
		UpdatedAt:         time.Now().Unix(),
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type PutBlogRequest struct {
	Id        int64
	UpdatedAt int64
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id int64) *DeleteBlogRequest {
	return &DeleteBlogRequest{Id: id}
}

type DeleteBlogRequest struct {
	Id int64
}
