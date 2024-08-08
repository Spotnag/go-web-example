package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

//func Make(h HTTPHandler) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		if err := h(w, r); err != nil {
//			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
//		}
//	}
//}

func Render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
