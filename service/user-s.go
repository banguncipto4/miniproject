package service

import (
	"MINIPROJECT/models"
	"MINIPROJECT/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUser() UserService {
	return UserService{
		userRepository: &repository.UserRepositoryImpl{},
	}
}

func (us *UserService) GetAllUser() []models.User {
	return us.userRepository.GetAllUser()
}

func (us *UserService) GetByIdUser(id string) models.User {
	return us.userRepository.GetByIdUser(id)
}

func (us *UserService) UpdateUser(id string, input models.Register) models.User {
	return us.userRepository.UpdateUser(id, input)
}

func (us *UserService) DeleteUser(id string) bool {
	return us.userRepository.DeleteUser(id)
}
