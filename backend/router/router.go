package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/satokiy/giraffe/adaptor/restapi/controller"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func NewRouter(gc controller.GiraffeController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.BodyDump(bodyDumpHandler))

	giraffe := e.Group("/giraffe")
	giraffe.GET("/random", gc.GetRandom)
	return e
}
