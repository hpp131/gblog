package response

import (
	"github.com/gin-gonic/gin"
	"github.com/hpp131/gblog/exception"
)

func Success(c *gin.Context, data any) {
	c.JSON(200, data)
}

func Failed(c *gin.Context, err error) {
	var resp  *exception.APIException
	if e, ok := err.(*exception.APIException);ok{
		resp = e
	}else{
		resp = exception.NewAPIException(500, err.Error())
	}
	
	c.JSON(500, resp)
	c.Abort()
}