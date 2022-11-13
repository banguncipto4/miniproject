package repository

import (
	"MINIPROJECT/database"
	"MINIPROJECT/models"
	"MINIPROJECT/tokenjwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct{}

func (a *AuthRepositoryImpl) Register(input models.Register) models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser models.User = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(password),
	}

	var createdUser models.User = models.User{}

	result := database.DB.Create(&newUser)

	result.Last(&createdUser)

	return createdUser
}

func (a *AuthRepositoryImpl) Login(input models.Login) string {
	var user models.User = models.User{}

	database.DB.First(&user, "email", input.Email)
	if user.ID == 0 {
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}
	tokens := tokenjwt.CreateToken(user.ID)

	return tokens
}
