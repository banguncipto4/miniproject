package repository

import (
	"MINIPROJECT/database"
	"MINIPROJECT/models"
)

type RatingRepositoryImpl struct{}

func (rr *RatingRepositoryImpl) GetAllRating() []models.Rating {
	var ratings []models.Rating

	database.DB.Find(&ratings)

	return ratings

}

func (rr *RatingRepositoryImpl) GetByIdRating(id string) models.Rating {
	var rating models.Rating

	database.DB.First(&rating, "id =?", id)
	return rating
}

func (rr *RatingRepositoryImpl) CreateRating(input models.InputRating) models.Rating {
	var newRating models.Rating = models.Rating{
		Star:     input.Star,
		Reaction: input.Reaction,
	}

	var createRating models.Rating = models.Rating{}

	result := database.DB.Create(&newRating)

	result.Last(&createRating)

	return createRating

}

func (rr *RatingRepositoryImpl) UpdateRating(id string, input models.InputRating) models.Rating {

	var rating models.Rating = rr.GetByIdRating(id)

	rating.Star = input.Star
	rating.Reaction = input.Reaction

	database.DB.Save(&rating)

	return rating
}

func (rr *RatingRepositoryImpl) DeleteRating(id string) bool {
	var rating models.Rating = rr.GetByIdRating(id)

	result := database.DB.Delete(&rating)

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}
