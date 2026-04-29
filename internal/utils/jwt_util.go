package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var AuthParams = viper.GetStringMapString("jwt")
var JwtSecret = []byte(AuthParams["jwt_secret"])

func GenerateToken(userID uint, userName, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"user_name": userName,
		"role":      role,
		"exp":       time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
