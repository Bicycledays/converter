package handler

import "github.com/bicycledays/converter/src/service"

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}
