package http

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/RianNegreiros/vigilate/domain"
	"github.com/RianNegreiros/vigilate/remote-server/delivery/http/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

var jwtSecret = os.Getenv("JWT_SECRET")

type ResponseError struct {
	Message string `json:"message"`
}

type RemoteServerHandler struct {
	RemoteServerUsecase domain.RemoteServerUsecase
}

func NewRemoteServerHandler(e *echo.Echo, us domain.RemoteServerUsecase) {
	handler := &RemoteServerHandler{
		RemoteServerUsecase: us,
	}

	e.POST("/remote-servers", handler.Create, middleware.JWTMiddleware)
	e.GET("/remote-servers", handler.GetByUserID, middleware.JWTMiddleware)
}

func (h *RemoteServerHandler) Create(c echo.Context) (err error) {
	var remoteServer domain.CreateRemoteServer
	err = c.Bind(&remoteServer)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	err = h.RemoteServerUsecase.Create(c.Request().Context(), &remoteServer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, remoteServer)
}

func (h *RemoteServerHandler) GetByUserID(c echo.Context) error {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
	}

	userID, err := getUserIDFromJWTToken(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
	}

	servers, err := h.RemoteServerUsecase.GetByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, servers)
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
