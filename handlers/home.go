package handlers

import (
	"github.com/labstack/echo/v4"
	"go-web-example/views/home"
)

func HandleHome(c echo.Context) error {
	return Render(c, home.Index())
}
