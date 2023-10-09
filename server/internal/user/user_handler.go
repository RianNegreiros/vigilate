package user

import (
	"net/http"

	"github.com/labstack/echo"
)

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
