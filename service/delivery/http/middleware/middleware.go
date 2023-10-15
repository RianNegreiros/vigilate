package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

type GoMiddleware struct {
}

func (m *GoMiddleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := c.Cookie("jwt")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error getting cookie")
		}

		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		return next(c)
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
