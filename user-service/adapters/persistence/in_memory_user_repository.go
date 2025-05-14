package persistence

import (
	"database/sql"
	"fmt"
	"user-service/domain" // Importar correctamente el paquete domain
)

type InMemoryUserRepo struct {
	users map[string]domain.User // Usar domain.User en vez de User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	// Contraseñas hasheadas con bcrypt
	hash := "$2a$12$VplG3Yp8gY.T/qb3cz4yaeLPMKrT2yL/UQNDaIWGz/hfLWr0i7MNe" // "123456"
	return &InMemoryUserRepo{
		users: map[string]domain.User{ // Usar domain.User
			"admin": {
				ID:       1,
				Username: "admin",
				Password: hash,
				Role:     domain.RoleAdmin, // Usar domain.RoleAdmin
			},
		},
	}
}

func (r *InMemoryUserRepo) Save(user *domain.User) error {
	if _, exists := r.users[user.Username]; exists {
		return fmt.Errorf("usuario ya existe")
	}
	user.ID = uint(len(r.users) + 1)
	r.users[user.Username] = *user
	return nil
}

func (r *InMemoryUserRepo) FindByUsername(username string) (*domain.User, error) { // Usar domain.User
	user, ok := r.users[username]
	if !ok {
		return nil, nil
	}
	return &user, nil // Devolver una referencia al tipo correcto
}

func (r *PostgresUserRepo) GetRole(username string) (domain.Role, error) {
	var role domain.Role
	err := r.db.QueryRow("SELECT role FROM users WHERE username = $1", username).Scan(&role)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	}
	if err != nil {
		return "", fmt.Errorf("error fetching role: %v", err)
	}
	return role, nil
}
