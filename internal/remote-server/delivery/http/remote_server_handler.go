package http

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

type RemoteServerHandler struct {
	RemoteServerUsecase domain.RemoteServerUsecase
}

func NewRemoteServerHandler(e *echo.Echo, r domain.RemoteServerUsecase) {
	handler := &RemoteServerHandler{
		RemoteServerUsecase: r,
	}

	e.POST("/remote-servers", handler.Create, middleware.JWTMiddleware)
	e.GET("/remote-servers", handler.GetByUserID, middleware.JWTMiddleware)
	e.GET("/remote-servers/:id", handler.GetByID, middleware.JWTMiddleware)
	e.POST("/remote-servers/:id/start-monitoring", handler.StartMonitoring, middleware.JWTMiddleware)
	e.PUT("/remote-servers/:id", handler.UpdateNameAddress, middleware.JWTMiddleware)
	e.DELETE("/remote-servers/:id", handler.Delete, middleware.JWTMiddleware)
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

func (h *RemoteServerHandler) GetByID(c echo.Context) error {
	serverID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	server, err := h.RemoteServerUsecase.GetByID(c.Request().Context(), serverID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, server)
}

func (h *RemoteServerHandler) StartMonitoring(c echo.Context) error {
	serverID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	err = h.RemoteServerUsecase.StartMonitoring(serverID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseError{Message: "Monitoring started successfully"})
}

func (h *RemoteServerHandler) UpdateNameAddress(c echo.Context) error {
	serverID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	var updateRemoteServer domain.UpdateRemoteServer
	err = c.Bind(&updateRemoteServer)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	updateRemoteServer.ID = int64(serverID)

	server, err := h.RemoteServerUsecase.UpdateNameAddress(c.Request().Context(), &updateRemoteServer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, server)
}

func (h *RemoteServerHandler) Delete(c echo.Context) error {
	serverID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	err = h.RemoteServerUsecase.Delete(c.Request().Context(), serverID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseError{Message: "Server deleted successfully"})
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
