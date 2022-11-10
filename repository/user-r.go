package repository

import (
	"MINIPROJECT/database"
	"MINIPROJECT/models"
)

type UserRepositoryImpl struct{}

func (ur *UserRepositoryImpl) GetAllUser() []models.User {
	var users []models.User

	database.DB.Find(&users)

	return users

}

func (ur *UserRepositoryImpl) GetByIdUser(id string) models.User {
	var user models.User

	database.DB.First(&user, "id_user =?", id)

	return user

}

func (ur *UserRepositoryImpl) UpdateUser(id string, input models.Register) models.User {
	var user models.User = ur.GetByIdUser(id)

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	database.DB.Save(&user)

	return user

}

func (ur *UserRepositoryImpl) DeleteUser(id string) bool {
	var user models.User = ur.GetByIdUser(id)

	result := database.DB.Delete(&user)

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
