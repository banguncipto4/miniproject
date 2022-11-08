package service

import (
	"MINIPROJECT/models"
	"MINIPROJECT/repository"
)

type RatingService struct {
	RatingRepository repository.RatingRepository
}

func NewRatingService() RatingService {
	return RatingService{
		RatingRepository: &repository.RatingRepositoryImpl{},
	}
}

func (rs *RatingService) GetAllRating() []models.Rating {
	return rs.RatingRepository.GetAllRating()
}

func (rs *RatingService) GetByIdRating(id string) models.Rating {
	return rs.RatingRepository.GetByIdRating(id)
}

func (rs *RatingService) CreateRating(input models.InputRating) models.Rating {
	return rs.RatingRepository.CreateRating(input)
}

func (rs *RatingService) UpdateRating(id string, input models.InputRating) models.Rating {
	return rs.RatingRepository.UpdateRating(id, input)
}

func (rs *RatingService) DeleteRating(id string) bool {

	return rs.RatingRepository.DeleteRating(id)

}
