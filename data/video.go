package data

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/google/uuid"
	"os"
)

// Maximum upload size for cloudflare is 200mb
func (d *Service) uploadStandardVideo(location string) (err error, file cloudflare.StreamVideo) {
	videoID := uuid.New().String()
	ctx := context.Background()

	file, err = c.API.StreamUploadVideoFile(ctx, cloudflare.StreamUploadFileParameters{
		AccountID:         os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		VideoID:           videoID,
		FilePath:          location,
		ScheduledDeletion: nil,
	})
	return
}
