package blogs

import "time"


func NewBlog(in  *CreateBlogRequest) *Blog{
	return &Blog{
		CreatedAt: time.Now().Unix(),
		CreateBlogRequest: in,
	}
}
type Blog struct {
	*CreateBlogRequest

	Status      string
	PublishedAt int64

	Id          int64
	UpdatedAt   int64
	CreatedAt   int64

	AuditAt     int64
	IsAuditPass int
}
