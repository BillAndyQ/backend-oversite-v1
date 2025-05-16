package http

import "orden-trabajo-service/services"

type OTHandler struct {
	otService *services.OTService
}

func NewOTHandler(auth *services.OTService) *OTHandler {
	return &OTHandler{otService: auth}
}
