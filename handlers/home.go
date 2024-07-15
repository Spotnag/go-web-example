package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-web-example/views/home"
)

func HandleHome(c echo.Context) error {
	return home.Index().Render(context.Background(), c.Response().Writer)
}
