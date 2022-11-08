package repository

import (
	"MINIPROJECT/database"
	"MINIPROJECT/models"
)

type GameRepositoryImpl struct{}

func (gr *GameRepositoryImpl) GetAllGame() []models.Game {
	var games []models.Game

	database.DB.Find(&games)

	return games

}

func (gr *GameRepositoryImpl) GetByIdGame(id string) models.Game {
	var game models.Game

	database.DB.First(&game, "id =?", id)
	return game
}

func (gr *GameRepositoryImpl) CreateGame(input models.InputGame) models.Game {
	var newGame models.Game = models.Game{
		Game_name:   input.Game_name,
		Game_type:   input.Game_type,
		Game_desc:   input.Game_desc,
		Game_access: input.Game_access,
	}

	var createGame models.Game = models.Game{}

	result := database.DB.Create(&newGame)

	result.Last(&createGame)

	return createGame

}

func (gr *GameRepositoryImpl) UpdateGame(id string, input models.InputGame) models.Game {

	var game models.Game = gr.GetByIdGame(id)

	game.Game_name = input.Game_name
	game.Game_type = input.Game_type
	game.Game_desc = input.Game_desc
	game.Game_access = input.Game_access

	database.DB.Save(&game)

	return game
}

func (gr *GameRepositoryImpl) DeleteGame(id string) bool {
	var game models.Game = gr.GetByIdGame(id)

	result := database.DB.Delete(&game)

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}
