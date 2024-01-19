package tokens

type Token struct {
	AccessToken           string
	AccessTokenExpiresAt  int64
	RefreshToken          string
	RefreshTokenExpiresAt int64
	CreatedAt             int64
	UpdatedAt             int64
	UserId                int64
	Username              string
}


func NewToken() *Token {
	return &Token{}
}