package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Game struct {
	Id_game      uint      `json:"id_game" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Game_name    string    `json:"game_name"`
	Game_type    string    `json:"game_type"`
	Game_desc    string    `json:"game_desc"`
	Game_access  string    `json:"game_access"`
	Publisher    Publisher `json:"publisher" gorm:"foreignKey:Id_publisher"`
	Id_publisher uint      `json:"id_publisher"`
	Rating       Rating    `json:"rating" gorm:"foreignKey:Id_rating"`
	Id_rating    uint      `json:"id_rating"`
}

type InputGame struct {
	Game_name   string `json:"game_name" validate:"required"`
	Game_type   string `json:"game_type" validate:"required"`
	Game_desc   string `json:"game_desc" validate:"required"`
	Game_access string `json:"game_access" validate:"required"`
}

func (input *InputGame) ValidasiGame() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err
}
