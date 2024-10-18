package handlers

import (
	"github.com/labstack/echo/v4"
	"go-web-example/views/courses"
	"net/http"
)

func (u *Handler) Course(c echo.Context) error {
	return Render(c, courses.Course())
}

func (u *Handler) GetCourse(w http.ResponseWriter, r *http.Request) {
}

func (u *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
}

func (u *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
}

func (u *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
}
