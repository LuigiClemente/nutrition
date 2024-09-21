package handlers

import (
	"nutrition/services"
)

type Handler struct {
	service services.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: *services.NewService(),
	}
}
