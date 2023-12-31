package user

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/RianNegreiros/vigilate/internal/domain"
	"github.com/RianNegreiros/vigilate/internal/remote-server/delivery/http/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = os.Getenv("JWT_SECRET")

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(e *echo.Echo, us domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}

	e.POST("/signup", handler.CreateUser)
	e.POST("/login", handler.Login)
	e.GET("/logout", handler.Logout)
	e.PATCH("/email-notification-preferences", handler.UpdateEmailNotificationPreferences, middleware.JWTMiddleware)
	e.GET("/users/:id", handler.GetByID, middleware.JWTMiddleware)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var u domain.CreateUserRequest
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return nil
	}

	res, err := h.UserUsecase.CreateUser(c.Request().Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return nil
	}

	c.JSON(http.StatusCreated, res)
	return nil
}

func (h *UserHandler) Login(c echo.Context) error {
	var user domain.LoginUserRequest
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return nil
	}

	u, err := h.UserUsecase.Login(c.Request().Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return nil
	}

	cookieAge := 60 * 60 * 24 * 7 // 7 days
	writeCookie(c, "jwt", u.AccessToken, cookieAge)
	return c.JSON(http.StatusOK, u)
}

func (h *UserHandler) Logout(c echo.Context) error {
	writeCookie(c, "jwt", "", -1)
	c.JSON(http.StatusOK, echo.Map{"message": "logout success"})

	return nil
}

func (h *UserHandler) UpdateEmailNotificationPreferences(c echo.Context) error {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseError{Message: "error getting cookie"})
		log.Println("Error getting cookie: ", err)
		return nil
	}

	userID, err := getUserIDFromJWTToken(cookie.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseError{Message: "invalid token"})
		log.Println("Error getting user ID from JWT token: ", err)
		return nil
	}

	err = h.UserUsecase.UpdateEmailNotificationPreferences(c.Request().Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return nil
	}

	c.JSON(http.StatusOK, echo.Map{"message": "notification preferences updated"})
	return nil
}

func writeCookie(c echo.Context, name, value string, maxAge int) {
	domain := os.Getenv("DOMAIN")
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.MaxAge = maxAge
	cookie.Path = "/"
	cookie.Domain = domain
	cookie.HttpOnly = false
	cookie.Secure = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}

func getUserIDFromJWTToken(cookieValue string) (int, error) {
	token, err := jwt.Parse(cookieValue, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		log.Println("Error parsing JWT token", err)
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if idClaim, ok := claims["id"].(string); ok {
			userID, err := strconv.Atoi(idClaim)
			if err != nil {
				log.Println("Error converting ID claim to int", err)
				return 0, err
			}
			return userID, nil
		}
	}

	return 0, errors.New("ID claim not found in JWT token")
}

func (h *UserHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return nil
	}

	user, err := h.UserUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return nil
	}

	c.JSON(http.StatusOK, user)
	return nil
}
