package blogs

type Blog struct {
	CreateBlogRequest

	Status      string
	PublishedAt int64

	Id          int64
	UpdatedAt   int64
	CreatedAt   int64

	AuditAt     int64
	IsAuditPass int
}
