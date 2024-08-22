package api

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

type Service struct {
	API *cloudflare.API
}

func NewApiService() (*Service, error) {
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		return nil, err
	}
	return &Service{
		API: api,
	}, nil
}

// Maximum upload size for cloudflare is 200mb
func (d *Service) UploadStandardVideo(location string) (cloudflare.StreamVideo, error) {
	absPath, err := filepath.Abs(location)
	if err != nil {
		return cloudflare.StreamVideo{}, err
	}

	ctx := context.Background()
	videoID := uuid.New().String()

	file, err := d.API.StreamUploadVideoFile(ctx, cloudflare.StreamUploadFileParameters{
		AccountID:         os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		VideoID:           videoID,
		FilePath:          absPath,
		ScheduledDeletion: nil,
	})
	if err != nil {
		return cloudflare.StreamVideo{}, err
	}

	return file, nil
}
