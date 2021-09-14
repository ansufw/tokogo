package jwtauth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	Name string
	Kind string
}

type ConsumerClaim struct {
	*jwt.StandardClaims
	TokenType string
	UserInfo
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractTokenUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		return userId
	}
	return 0
}
