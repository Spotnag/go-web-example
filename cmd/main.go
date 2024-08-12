package main

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-web-example/handlers"
	"go-web-example/service"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("create table if not exists user (id text not null primary key, username text, password text);")
	if err != nil {
		log.Fatalf("%q: %s\n", err)
	}

	userService := service.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	// TODO: Remove this in production
	_, err = userService.CreateUser("admin@time", "passtime")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	//e.Renderer = newTemplate()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handlers.CheckLoggedInMiddleware)
	//e.Use(middleware.CSRF()) // TODO FIX

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Debug = true // TODO REMOVE IN PRODUCTION

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.GET("/", handlers.HandleHome)
	e.GET("/login", userHandler.HandleLoginIndex)
	e.POST("/login", userHandler.HandleLogin)
	e.POST("/logout", userHandler.HandleLogout)

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
