package ports

import (
	"orden-trabajo-service/domain"
	"orden-trabajo-service/shared"
)

type OTRepository interface {
	GetAllOTEq(p shared.Pagination) ([]domain.OTEqAdminDTO, int, error)
	// GetAllOTPers() ([]domain.OrdenTrabajoPersona, error)
}
