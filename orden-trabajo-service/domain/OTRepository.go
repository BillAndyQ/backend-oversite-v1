package domain

import "context"

// Definimos la interfaz OTRepository
type OTRepository interface {
	RegisterOTEq(ctx context.Context, orden *OrdenTrabajoEquipo) (bool, error)
}
