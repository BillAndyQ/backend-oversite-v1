package services

import (
	"fmt"
	"orden-trabajo-service/domain"
	"orden-trabajo-service/ports"
	"orden-trabajo-service/shared"
)

type OTService struct {
	otRepo ports.OTRepository
}

func NewOTService(otRepo ports.OTRepository) OTService {
	return OTService{otRepo: otRepo}
}

func (s *OTService) ListOT_Equipo_Admin(p shared.Pagination) ([]domain.OTEqAdminDTO, int, error) {
	lista, total, err := s.otRepo.GetAllOTEq(p)
	if err != nil {
		return nil, 0, fmt.Errorf("error al obtener la lista de órdenes: %w", err)
	}
	return lista, total, nil
}
