package api

import "github.com/hpp131/gblog/exception"

var (
	ErrQueryParam = exception.NewAPIException(5002, "QueryParam Error, lack pagenum or pagenum")
)