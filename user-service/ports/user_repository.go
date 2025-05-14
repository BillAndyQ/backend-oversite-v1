package ports

import "user-service/domain"

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
	Save(user *domain.User) error
	GetRole(username string) (domain.Role, error)
}
