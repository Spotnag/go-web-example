package handlers

import (
	"go-web-example/api"
	"go-web-example/data"
)

type Handler struct {
	data       *data.Service
	cloudflare *api.CloudflareService
	mailgun    *api.MailgunService
}

func NewHandlers(dataService *data.Service, cloudflareService *api.CloudflareService, mailgun *api.MailgunService) *Handler {
	return &Handler{
		data:       dataService,
		cloudflare: cloudflareService,
		mailgun:    mailgun,
	}
}
