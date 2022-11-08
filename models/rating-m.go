package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Rating struct {
	Id_rating uint      `json:"id_rating" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Star      uint      `json:"star"`
	Reaction  string    `json:"reaction"`
}

type InputRating struct {
	Star     uint   `json:"star"     validate:"required"`
	Reaction string `json:"reaction" validate:"required"`
}

func (input *InputRating) ValidasiRating() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err
}
