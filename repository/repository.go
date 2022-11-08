package repository

import "MINIPROJECT/models"

type AutRepository interface {
	Register(input models.Register) models.User
	Login(input models.Login) string
}

type UserRepository interface {
	GetAllUser() []models.User
	GetByIdUser(id string) models.User
	UpdateUser(id string, input models.Register) models.User
	DeleteUser(id string) bool
}

type GameRepository interface {
	GetAllGame() []models.Game
	GetByIdGame(id string) models.Game
	CreateGame(input models.InputGame) models.Game
	UpdateGame(id string, input models.InputGame) models.Game
	DeleteGame(id string) bool
}

type PublisherRepository interface {
	GetAllPublisher() []models.Publisher
	GetByIdPublisher(id string) models.Publisher
	CreatePublisher(input models.InputPublisher) models.Publisher
	UpdatePublisher(id string, input models.InputPublisher) models.Publisher
	DeletePublisher(id string) bool
}

type RatingRepository interface {
	GetAllRating() []models.Rating
	GetByIdRating(id string) models.Rating
	CreateRating(input models.InputRating) models.Rating
	UpdateRating(id string, input models.InputRating) models.Rating
	DeleteRating(id string) bool
}
