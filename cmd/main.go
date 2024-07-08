package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"

	"io"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Renderer = newTemplate()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Debug = true // TODO REMOVE IN PRODUCTION

	e.Static("/images", "images")
	e.Static("/css", "css")

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "admin" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/dashboard", func(c echo.Context) error {
		return c.String(200, "Admin Dashboard")
	})
	e.GET("/", func(c echo.Context) error {

		return c.Render(200, "base.html", nil)
	})
	e.GET("/hello", handler)

	e.Logger.Fatal(e.Start("localhost:3000")) // TODO REMOVE IN PRODUCTION
}

func handler(c echo.Context) error {
	return c.Render(200, "base.html", nil)
}
