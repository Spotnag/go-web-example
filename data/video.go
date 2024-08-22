package data

import (
	"github.com/google/uuid"
	"go-web-example/models"
)

func (d *Service) CreateVideo(title, description, user, uploadedID string) (models.Video, error) {
	video := models.Video{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		UploadedID:  uploadedID,
		UploadedBy:  user,
	}

	_, err := d.DB.Exec("insert into video (id, title, description) values (?, ?, ?)", video.ID, video.Title, video.Description)
	if err != nil {
		return models.Video{}, err
	}
	return video, err
}
