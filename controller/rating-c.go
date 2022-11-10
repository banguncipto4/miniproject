package controller

import (
	"MINIPROJECT/models"
	"MINIPROJECT/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ratingService service.RatingService = service.NewRatingService()

func GetAllRating(e echo.Context) error {
	var ratings []models.Rating = ratingService.GetAllRating()

	return e.JSON(http.StatusOK, ratings)
}

func GetByIdRating(e echo.Context) error {
	var ratingId string = e.Param("id")
	rating := ratingService.GetByIdRating(ratingId)
	if rating.ID == 0 {
		return e.JSON(http.StatusNotFound, map[string]string{
			"message": "note not found",
		})
	}
	return e.JSON(http.StatusOK, rating)
}

func CreateRating(e echo.Context) error {
	var createRating *models.InputRating = new(models.InputRating)

	if err := e.Bind(createRating); err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}

	err := createRating.ValidasiRating()

	if err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "Wajib Di isi",
		})
	}

	rating := ratingService.CreateRating(*createRating)

	return e.JSON(http.StatusAccepted, rating)
}

func UpdateRating(e echo.Context) error {
	var updateRating *models.InputRating = new(models.InputRating)

	if err := e.Bind(updateRating); err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}
	err := updateRating.ValidasiRating()

	if err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}

	var idRating string = e.Param("id")

	if idRating == "" {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "kosong",
		})
	}

	result := ratingService.UpdateRating(idRating, *updateRating)

	return e.JSON(http.StatusAccepted, result)

}

func DeleteRating(e echo.Context) error {
	var ratingId string = e.Param("id")

	successDelete := ratingService.DeleteRating(ratingId)

	if !successDelete {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"messege": "failed delete",
		})
	}

	return e.JSON(http.StatusAccepted, map[string]string{
		"messege": "data deleted",
	})
}
