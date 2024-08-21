package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/labstack/echo/v4"
	"go-web-example/data"
	"go-web-example/models"
	"go-web-example/views/courses"
	"net/http"
)

type CourseHandler struct {
}

func (u *Handler) Course(c echo.Context) error {
	return Render(c, courses.Course())
}

func (c CourseHandler) ListCourses(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(data.ListCourses())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (c CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	course := data.GetCourse(ID)
	if course == nil {
		http.Error(w, "course not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(course)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (c CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	data.CreateCourse(course)
	err = json.NewEncoder(w).Encode(course)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (c CourseHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	updatedCourse := data.UpdateCourse(ID, course)
	if updatedCourse == nil {
		http.Error(w, "course not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedCourse)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (c CourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	deletedCourse := data.DeleteCourse(ID)
	if deletedCourse == nil {
		http.Error(w, "course not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
