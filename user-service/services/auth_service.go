package services

import (
	"fmt"
	"user-service/domain"
	"user-service/ports"
	"user-service/utils"
)

type AuthService struct {
	userRepo ports.UserRepository
}

func NewAuthService(userRepo ports.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("usuario no encontrado")
	}
	if user.Password != password {
		return "", fmt.Errorf("contraseña incorrecta")
	}

	// Genera JWT real
	token, err := utils.GenerateJWT(user.Username, string(user.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(username, password string, role domain.Role) (string, error) {
	existingUser, _ := s.userRepo.FindByUsername(username)
	if existingUser != nil {
		return "", fmt.Errorf("el usuario ya existe")
	}

	user := &domain.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	if err := s.userRepo.Save(user); err != nil {
		return "", err
	}

	// Generar JWT para el nuevo usuario
	token, err := utils.GenerateJWT(user.Username, string(user.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GetUserByUsername(username string) (*domain.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el usuario: %v", err)
	}
	return user, nil
}

func (s *AuthService) GetRole(username string) (domain.Role, error) {
	role, err := s.userRepo.GetRole(username)
	if err != nil {
		return "", fmt.Errorf("error al obtener el rol del usuario: %v", err)
	}
	return role, nil
}
