// NewDBConnection establece una nueva conexión a la base de datos PostgreSQL.
package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDBConnection(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}
