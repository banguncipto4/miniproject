package service

import (
	"MINIPROJECT/models"
	"MINIPROJECT/repository"
)

type AuthService struct {
	authRepository repository.AutRepository
}

func NewAuthService() AuthService {
	return AuthService{
		authRepository: &repository.AuthRepositoryImpl{},
	}

}

func (a *AuthService) Register(input models.Register) models.User {
	return a.authRepository.Register(input)
}

func (a *AuthService) Login(input models.Login) string {
	return a.authRepository.Login(input)
}
