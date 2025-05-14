package persistence

import (
	"database/sql"
	"fmt"
	"user-service/domain"
)

type PostgresUserRepo struct {
	db *sql.DB
}

// NewPostgresUserRepo crea una nueva instancia de PostgresUserRepo utilizando una conexión existente.
func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Save(user *domain.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Username, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepo) FindByUsername(username string) (*domain.User, error) {
	stmt, err := r.db.Prepare("SELECT id, username, password, role FROM users WHERE username = $1")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var user domain.User
	err = stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}
	return &user, nil
}

// Close cierra la conexión a la base de datos.
func (r *PostgresUserRepo) Close() error {
	return r.db.Close()
}
