package persistence

import (
	"database/sql"
)

type PostgresOTRepo struct {
	db *sql.DB
}

// NewPostgresUserRepo crea una nueva instancia de PostgresUserRepo utilizando una conexión existente.
func NewPostgresRepo(db *sql.DB) *PostgresOTRepo {
	return &PostgresOTRepo{db: db}
}

// Close cierra la conexión a la base de datos.
func (r *PostgresOTRepo) Close() error {
	return r.db.Close()
}
