package http

import (
	"net/http"

	"github.com/RianNegreiros/vigilate/domain"
	"github.com/labstack/echo"
)

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

	e.POST("/remote-servers", handler.Create)
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
