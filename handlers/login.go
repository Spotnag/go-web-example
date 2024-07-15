package handlers

import (
	"github.com/labstack/echo/v4"
	"go-web-example/views/auth"
)

func HandleLoginIndex(c echo.Context) error {
	//return auth.Login().Render(context.Background(), c.Response().Writer)
	return Render(c, auth.Login())
}
