package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Publisher struct {
	Id_publisher   uint      `json:"id_publisher" gorm:"primaryKey"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Publisher_name string    `json:"publisher_name"`
	Publisher_desc string    `json:"publisher_desc"`
}

type InputPublisher struct {
	Publisher_name string `json:"publisher_name" validate:"required"`
	Publisher_desc string `json:"publisher_desc" validate:"required"`
}

func (input *InputPublisher) ValidasiPublisher() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err
}
