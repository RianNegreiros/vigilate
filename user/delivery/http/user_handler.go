package user

import (
	"net/http"
	"os"

	"github.com/RianNegreiros/vigilate/domain"
	"github.com/labstack/echo"
)

var domainURL = os.Getenv("DOMAIN_URL")

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

func writeCookie(c echo.Context, name, value string, maxAge int) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.MaxAge = maxAge
	cookie.Path = "/"
	cookie.Domain = domainURL
	cookie.Secure = false
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}
