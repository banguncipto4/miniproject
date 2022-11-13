package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint      `json:"id_user" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type Register struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (input *Register) ValidasiRegister() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (input *Login) ValidasiLogin() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err
}
