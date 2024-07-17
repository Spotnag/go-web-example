package main

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-web-example/handlers"
	"net/http"
)

//type Template struct {
//	tmpl *template.Template
//}
//
//func newTemplate() *Template {
//	return &Template{
//		tmpl: template.Must(template.ParseGlob("views/*.html")),
//	}
//}
//
//func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
//	return t.tmpl.ExecuteTemplate(w, name, data)
//}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	e := echo.New()
	//e.Renderer = newTemplate()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middleware.CSRF()) // TODO FIX
	e.Debug = true // TODO REMOVE IN PRODUCTION
	e.Static("/images", "images")
	e.Static("/css", "css")

	e.Use(checkAuthMiddleware)
	e.GET("/", handlers.HandleHome)
	e.POST("/login", handlers.HandleLoginIndex)
	e.GET("/logout", logoutHandler)

	e.Logger.Fatal(e.Start("localhost:3000")) // TODO REMOVE IN PRODUCTION
}

func checkAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "session")
		isAuthenticated := session.Values["authenticated"]
		if isAuthenticated == nil {
			isAuthenticated = false
		}
		c.SetRequest(
			c.Request().WithContext(
				context.WithValue(
					c.Request().Context(),
					"isAuthenticated",
					isAuthenticated),
			),
		)
		return next(c)
	}
}

func loginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Validate the username and password (this is just an example, in real application use hashed passwords)
	if username == "admin" && password == "password" {
		session, _ := store.Get(c.Request(), "session")

		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   48 * 60 * 60, // 48 hours
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		session.Save(c.Request(), c.Response())
		return c.String(http.StatusOK, "Logged in successfully!")
	}
	return c.String(http.StatusUnauthorized, "Invalid username or password")
}

func logoutHandler(c echo.Context) error {
	session, _ := store.Get(c.Request(), "session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(c.Request(), c.Response())
	return c.String(http.StatusOK, "Logged out successfully!")
}
