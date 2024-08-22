package handlers

import (
	"go-web-example/api"
	"go-web-example/data"
)

type Handler struct {
	DataService *data.Service
	APIService  *api.Service
}

func NewHandlers(dataService *data.Service, apiService *api.Service) *Handler {
	return &Handler{
		DataService: dataService,
		APIService:  apiService,
	}
}
