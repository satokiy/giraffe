package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satokiy/giraffe/adaptor/restapi/apiresponse"
)

type HealthCheckController interface {
	GetHealthCheck(c echo.Context) error
}
type HealthCheckControllerImpl struct{}

func NewHealthCheckControllerImpl() HealthCheckController {
	return &HealthCheckControllerImpl{}
}

func (*HealthCheckControllerImpl) GetHealthCheck(c echo.Context) error {
	res := apiresponse.HealthCheckResponse{
		Message: "OK",
	}
	return c.JSON(http.StatusOK, res)
}
