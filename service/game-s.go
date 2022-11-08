package service

import (
	"MINIPROJECT/models"
	"MINIPROJECT/repository"
)

type GameService struct {
	GameRepository repository.GameRepository
}

func NewGameService() GameService {
	return GameService{
		GameRepository: &repository.GameRepositoryImpl{},
	}
}

func (gs *GameService) GetAllGame() []models.Game {
	return gs.GameRepository.GetAllGame()
}

func (gs *GameService) GetByIdGame(id string) models.Game {
	return gs.GameRepository.GetByIdGame(id)
}

func (gs *GameService) CreateGame(input models.InputGame) models.Game {
	return gs.GameRepository.CreateGame(input)
}

func (gs *GameService) UpdateGame(id string, input models.InputGame) models.Game {
	return gs.GameRepository.UpdateGame(id, input)
}

func (bs *GameService) DeleteGame(id string) bool {

	return bs.GameRepository.DeleteGame(id)

}
