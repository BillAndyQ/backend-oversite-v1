package persistence

import (
	"database/sql"
	"fmt"
	"orden-trabajo-service/domain"
	"orden-trabajo-service/shared"
)

type PostgresOTRepo struct {
	db *sql.DB
}

// NewPostgresUserRepo crea una nueva instancia de PostgresUserRepo utilizando una conexión existente.
func NewPostgresRepo(db *sql.DB) *PostgresOTRepo {
	return &PostgresOTRepo{db: db}
}

func (r *PostgresOTRepo) GetAllOTEq(p shared.Pagination) ([]domain.OTEqAdminDTO, int, error) {
	p.Normalize(10) // Establece límite por defecto si no se proporciona

	// Obtener total de registros
	var total int
	countQuery := `SELECT COUNT(*) FROM public.ot_equipos`
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("error al obtener total de registros: %w", err)
	}

	// Obtener registros paginados
	query := `SELECT id, n_ord_trabajo, empresa_matriz, empresa_socia, n_ord_servicio, ruc, fecha_servicio, empresa_certificadora, tipo_unidad, placa, area 
			  FROM public.ot_equipos
			  ORDER BY id
			  LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, p.Limit, p.Offset())
	if err != nil {
		return nil, 0, fmt.Errorf("error al obtener registros paginados: %w", err)
	}
	defer rows.Close()

	var lista []domain.OTEqAdminDTO
	for rows.Next() {
		var ot domain.OTEqAdminDTO
		if err := rows.Scan(
			&ot.ID,
			&ot.NOrdenTrabajo,
			&ot.EmpresaMatriz,
			&ot.EmpresaSocia,
			&ot.NOrdenService,
			&ot.Ruc,
			&ot.FechaServicio,
			&ot.Certificadora,
			&ot.TipoUnidad,
			&ot.Placa,
			&ot.Area,
		); err != nil {
			return nil, 0, err
		}
		lista = append(lista, ot)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return lista, total, nil
}

func (r *PostgresOTRepo) GetAllOTPers() ([]domain.OrdenTrabajoPersona, error) {
	var lista []domain.OrdenTrabajoPersona

	rows, err := r.db.Query("SELECT id, n_orden_trabajo, empresa_matriz, empresa_socia FROM OT_equipo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ot domain.OrdenTrabajoPersona
		if err := rows.Scan(&ot.ID, &ot.NOrdenTrabajo, &ot.EmpresaMatriz); err != nil {
			return nil, err
		}
		lista = append(lista, ot)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lista, nil
}

// Close cierra la conexión a la base de datos.
func (r *PostgresOTRepo) Close() error {
	return r.db.Close()
}
