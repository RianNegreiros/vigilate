package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error getting cookie")
		}

		if cookie.Value == "" {
			return c.JSON(http.StatusUnauthorized, "No token found")
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error parsing token")
		}

		if token.Valid {
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}
}
