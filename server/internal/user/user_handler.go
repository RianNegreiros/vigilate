package user

import (
	"net/http"

	"github.com/labstack/echo"
)

const cookieAge = 7 * 24 * 60 * 60 // 7 days

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserService UserService
}

func NewUserHandler(e *echo.Echo, userService UserService) {
	handler := &UserHandler{
		UserService: userService,
	}

	e.POST("/signup", handler.CreateUser)
	e.POST("/login", handler.Login)
	e.GET("/logout", handler.Logout)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var u CreateUserRequest
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		return nil
	}

	res, err := h.UserService.CreateUser(c.Request().Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		return nil
	}

	c.JSON(http.StatusCreated, res)
	return nil
}

func (h *UserHandler) Login(c echo.Context) error {
	var user LoginUserRequest
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		return nil
	}

	u, err := h.UserService.Login(c.Request().Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		return nil
	}

	writeCookie(c, "jwt", u.accessToken, cookieAge) // 7 days
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
	cookie.Domain = domain
	cookie.Secure = false
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}
