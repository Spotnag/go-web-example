package handlers

import "go-web-example/data"

type Handler struct {
	DataService *data.Service
}

func NewHandlers(dataService *data.Service) *Handler {
	return &Handler{DataService: dataService}
}
