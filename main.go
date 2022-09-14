package main

import (
	"net/http"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	swagger, err := GetSwagger()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(oapimiddleware.OapiRequestValidator(swagger))

	v1, impl := e.Group("v1"), &ServerImpl{}
	RegisterHandlers(v1, impl)

	e.Logger.Fatal(e.Start(":8000"))
}

type ServerImpl struct{}

// GetHello implements ServerInterface
func (*ServerImpl) GetHello(ctx echo.Context, params GetHelloParams) error {
	return ctx.JSON(http.StatusOK,
		map[string]interface{}{"hello": params.Name})
}
