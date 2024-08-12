package handlers

import "go-web-example/data"

type Handler struct {
	DataService *data.DataService
}

func NewHandlers(dataService *data.DataService) *Handler {
	return &Handler{DataService: dataService}
}
