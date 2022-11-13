package controller

import (
	"MINIPROJECT/models"
	"MINIPROJECT/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var userService service.UserService = service.NewUser()

func GetAllUser(c echo.Context) error {
	var users []models.User = userService.GetAllUser()

	return c.JSON(http.StatusOK, users)
}

func GetByIdUser(c echo.Context) error {
	var userId string = c.Param("id")

	user := userService.GetByIdUser(userId)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "note not found",
		})
	}
	return c.JSON(http.StatusOK, user)

}

func UpdateUser(c echo.Context) error {

	var UpdateUser *models.Register = new(models.Register)

	if err := c.Bind(UpdateUser); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed",
		})
	}
	err := UpdateUser.ValidasiRegister()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	var userId string = c.Param("id")

	result := userService.UpdateUser(userId, *UpdateUser)

	return c.JSON(http.StatusAccepted, result)

}

func DeleteUser(c echo.Context) error {
	var userId string = c.Param("id")

	successDelete := userService.DeleteUser(userId)

	if !successDelete {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messege": "failed delete",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"messege": "data deleted",
	})
}
