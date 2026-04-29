package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohdareeb0x-commits/book-library-api/internal/response"
	"github.com/mohdareeb0x-commits/book-library-api/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", "unable to authorize user")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.JwtSecret), nil
		})
		
		if err != nil || !token.Valid {
			response.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
