package service

import (
	"github.com/RianNegreiros/vigilate/domain"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ServiceHandler struct {
	ServiceUsecase domain.ServiceUsecase
}

func NewServiceHandler(e *echo.Echo, us domain.ServiceUsecase) {
	handler := &ServiceHandler{
		ServiceUsecase: us,
	}

	e.GET("/services/:id", handler.GetServiceByID)

	e.POST("/services", handler.CreateService)
}

func (h *ServiceHandler) CreateService(c echo.Context) error {
	var s domain.ServiceRequest
	if err := c.Bind(&s); err != nil {
		c.JSON(400, echo.Map{"error": err.Error()})
		return nil
	}

	res, err := h.ServiceUsecase.CreateService(c.Request().Context(), &s)
	if err != nil {
		c.JSON(500, echo.Map{"error": err.Error()})
		return nil
	}

	c.JSON(201, res)
	return nil
}

func (h *ServiceHandler) GetServiceByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.ServiceUsecase.GetServiceByID(c.Request().Context(), id)
	if err != nil {
		c.JSON(500, echo.Map{"error": err.Error()})
		return nil
	}

	c.JSON(200, res)
	return nil
}
