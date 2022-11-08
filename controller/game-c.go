package controller

import (
	"MINIPROJECT/models"
	"MINIPROJECT/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var gameService service.GameService = service.NewGameService()

func GetAllGame(e echo.Context) error {
	var games []models.Game = gameService.GetAllGame()

	return e.JSON(http.StatusOK, games)
}

func GetByIdGame(e echo.Context) error {
	var gameId string = e.Param("id")
	game := gameService.GetByIdGame(gameId)
	if game.Id_game == 0 {
		return e.JSON(http.StatusNotFound, map[string]string{
			"message": "note not found",
		})
	}
	return e.JSON(http.StatusOK, game)
}

func CreateGame(e echo.Context) error {
	var createGame *models.InputGame = new(models.InputGame)

	if err := e.Bind(createGame); err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}

	err := createGame.ValidasiGame()

	if err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "Wajib Di isi",
		})
	}

	game := gameService.CreateGame(*createGame)

	return e.JSON(http.StatusAccepted, game)
}

func UpdateGame(e echo.Context) error {
	var updateGame *models.InputGame = new(models.InputGame)

	if err := e.Bind(updateGame); err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}
	err := updateGame.ValidasiGame()

	if err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}

	var idGame string = e.Param("id")

	if idGame == "" {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "kosong",
		})
	}

	result := gameService.UpdateGame(idGame, *updateGame)

	return e.JSON(http.StatusAccepted, result)

}

func DeleteGame(e echo.Context) error {
	var gameId string = e.Param("id")

	successDelete := gameService.DeleteGame(gameId)

	if !successDelete {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"messege": "failed delete",
		})
	}

	return e.JSON(http.StatusAccepted, map[string]string{
		"messege": "data deleted",
	})
}
