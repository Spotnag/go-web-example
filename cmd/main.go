package main

import (
	"context"
	"errors"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-web-example/handlers"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	e := echo.New()

	//e.Renderer = newTemplate()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(checkLoggedInMiddleware)
	//e.Use(middleware.CSRF()) // TODO FIX

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Debug = true // TODO REMOVE IN PRODUCTION

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.GET("/", handlers.HandleHome)
	e.GET("/login", handlers.HandleLoginIndex)
	e.POST("/login", loginHandler)
	e.POST("/logout", logoutHandler)

	e.Logger.Fatal(e.Start("localhost:3000")) // TODO REMOVE IN PRODUCTION
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		if err := c.String(he.Code, he.Message.(string)); err != nil {
			c.Logger().Error(err)
		}
		return
	}
	c.Logger().Error(err) // log the original error
	if err = c.String(http.StatusInternalServerError, "internal server error"); err != nil {
		c.Logger().Error(err)
	}
}

func checkLoggedInMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "session")
		isLoggedIn := session.Values["loggedIn"]
		if isLoggedIn == nil {
			isLoggedIn = false
		}
		c.SetRequest(c.Request().WithContext(context.WithValue(
			c.Request().Context(),
			"isLoggedIn",
			isLoggedIn)))
		return next(c)
	}
}

func loginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Validate the username and password (this is just an example, in real application use hashed passwords)
	if username == "admin@admin" && password == "password" {
		session, _ := store.Get(c.Request(), "session")

		// Set user as authenticated
		session.Values["loggedIn"] = true
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   48 * 60 * 60, // 48 hours
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode, // TODO what are the differences here?
		}
		if err := session.Save(c.Request(), c.Response()); err != nil {
			return err
		}
		return hxRedirect(c, "/")
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
}

func logoutHandler(c echo.Context) error {
	session, _ := store.Get(c.Request(), "session")

	// Revoke users authentication
	session.Values["loggedIn"] = false
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return hxRedirect(c, "/")
}

func hxRedirect(c echo.Context, url string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", url)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func hxLocation(c echo.Context, url string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Location", url)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
