package domain

type Historial struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	OrdenTrabajoID string `json:"orden_trabajo_id"`
	FechaHora      string `json:"fecha_hora"`
	Accion         string `json:"accion"`
	Nombres        string `json:"nombres"`
}
