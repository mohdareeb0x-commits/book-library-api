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

		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		c.Set("role", role)
		c.Set("user_id", claims["user_id"])

		c.Next()
	}
}

func AdminOnly(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		response.Fail(c, http.StatusForbidden, "FORBIDDEN", "can't access, admin only")
		c.Abort()
		return
	}
	c.Next()
}
