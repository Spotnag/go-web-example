package api

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/google/uuid"
	"go-web-example/config"
	"path/filepath"
)

type CloudflareService struct {
	cloudflare *cloudflare.API
}

func NewCloudflareService() (*CloudflareService, error) {
	api, err := cloudflare.NewWithAPIToken(config.AppConfig.CloudflareAPIToken)
	if err != nil {
		return nil, err
	}
	return &CloudflareService{
		cloudflare: api,
	}, nil
}

// Maximum upload size for cloudflare is 200mb
func (d *CloudflareService) UploadStandardVideo(location string) (cloudflare.StreamVideo, error) {
	absPath, err := filepath.Abs(location)
	if err != nil {
		return cloudflare.StreamVideo{}, err
	}

	ctx := context.Background()
	videoID := uuid.New().String()

	file, err := d.cloudflare.StreamUploadVideoFile(ctx, cloudflare.StreamUploadFileParameters{
		AccountID:         config.AppConfig.CloudflareAccountID,
		VideoID:           videoID,
		FilePath:          absPath,
		ScheduledDeletion: nil,
	})
	if err != nil {
		return cloudflare.StreamVideo{}, err
	}

	return file, nil
}
