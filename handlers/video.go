package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type VideoHandler struct {
}

func (u *Handler) GetVideo() {
}

func (u *Handler) CreateVideo(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	path := c.FormValue("path")

	uploadedVideo, err := u.APIService.UploadStandardVideo(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Failed to upload video")
	}

	u.DataService.CreateVideo(title, description)

}

func (u *Handler) UpdateVideo() {
}

func (u *Handler) DeleteVideo() {
}
