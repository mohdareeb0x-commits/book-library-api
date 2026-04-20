package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohdareeb0x-commits/book-library-api/models"
)

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data: data,
	})
}

func Fail(c *gin.Context, status int, code, message string) {
  c.JSON(status, models.Response{
    Success: false,
    Error:   &models.ErrorInfo{Code: code, Message: message},
  })
}
