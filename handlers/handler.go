package handlers

import (
	"go-web-example/api"
	"go-web-example/data"
)

type Handler struct {
	db           *data.Service
	videoService *api.CloudflareService
	mailService  api.MailService
}

func NewHandlers(dbService *data.Service, videoService *api.CloudflareService, mailService api.MailService) *Handler {
	return &Handler{
		db:           dbService,
		videoService: videoService,
		mailService:  mailService,
	}
}
