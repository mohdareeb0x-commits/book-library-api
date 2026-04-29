package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/mohdareeb0x-commits/book-library-api/internal/dto"
)

func OK(c *gin.Context, data interface{}, meta *dto.Meta) {
	if meta == nil {
		c.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data:    data,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func Fail(c *gin.Context, status int, code, message string) {
	c.JSON(status, dto.Response{
		Success: false,
		Error:   &dto.ErrorInfo{Code: code, Message: message},
	})
}
