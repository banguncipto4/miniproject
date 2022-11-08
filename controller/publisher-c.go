package controller

import (
	"MINIPROJECT/models"
	"MINIPROJECT/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var publisherService service.PublisherService = service.NewPublisherService()

func GetAllPublisher(e echo.Context) error {
	var publishers []models.Publisher = publisherService.GetAllPublisher()

	return e.JSON(http.StatusOK, publishers)
}

func GetByIdPublisher(e echo.Context) error {
	var publisherId string = e.Param("id")
	publisher := publisherService.GetByIdPublisher(publisherId)
	if publisher.Id_publisher == 0 {
		return e.JSON(http.StatusNotFound, map[string]string{
			"message": "note not found",
		})
	}
	return e.JSON(http.StatusOK, publisher)
}

func CreatePublisher(e echo.Context) error {
	var createPublisher *models.InputPublisher = new(models.InputPublisher)

	if err := e.Bind(createPublisher); err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}

	err := createPublisher.ValidasiPublisher()

	if err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "Wajib Di isi",
		})
	}

	publisher := publisherService.CreatePublisher(*createPublisher)

	return e.JSON(http.StatusAccepted, publisher)
}

func UpdatePublisher(e echo.Context) error {
	var updatePublisher *models.InputPublisher = new(models.InputPublisher)

	if err := e.Bind(updatePublisher); err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}
	err := updatePublisher.ValidasiPublisher()

	if err != nil {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}

	var idPublisher string = e.Param("id")

	if idPublisher == "" {
		e.JSON(http.StatusBadRequest, map[string]string{
			"messege": "kosong",
		})
	}

	result := publisherService.UpdatePublisher(idPublisher, *updatePublisher)

	return e.JSON(http.StatusAccepted, result)

}

func DeletePublisher(e echo.Context) error {
	var publisherId string = e.Param("id")

	successDelete := publisherService.DeletePublisher(publisherId)

	if !successDelete {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"messege": "failed delete",
		})
	}

	return e.JSON(http.StatusAccepted, map[string]string{
		"messege": "data deleted",
	})
}
