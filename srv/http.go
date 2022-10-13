package srv

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	*echo.Echo
}

func NewHttpServer() *HttpServer {
	e := echo.New()

	e.GET("/livez", func(c echo.Context) error {
		return c.String(http.StatusOK, "alive and well")
	})

	e.GET("/readyz", func(c echo.Context) error {
		// @TODO hook dependencies in here to know when services that come up are actually ready
		return c.String(http.StatusOK, "ready for duty")
	})

	return &HttpServer{e}
}
