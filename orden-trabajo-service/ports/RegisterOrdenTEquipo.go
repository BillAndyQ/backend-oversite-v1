package ports

import (
	"context"
	"orden-trabajo-service/domain"
)

type RegisterOrden interface {
	RegisterOTEq(ctx context.Context, orden *domain.OrdenTrabajoEquipo) (bool, error)
	RegisterOTPers(ctx context.Context, orden *domain.OrdenTrabajoPersona) (bool, error)
}
