package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MyTestController struct{}

func (c *MyTestController) Register(e *echo.Echo) error {
	e.GET("/users", c.GetUser)

	return nil
}

func (mtc *MyTestController) GetUser(ctx echo.Context) error {
	ctx.JSON(http.StatusOK, "hello")

	return nil
}
