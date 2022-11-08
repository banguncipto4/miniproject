package controller

import (
	"MINIPROJECT/models"
	"MINIPROJECT/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var authService service.AuthService = service.NewAuthService()

func Register(c echo.Context) error {
	var userRegister *models.Register = new(models.Register)

	if err := c.Bind(userRegister); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}
	user := authService.Register(*userRegister)

	err := userRegister.ValidasiRegister()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	var userLogin *models.Login = new(models.Login)
	if err := c.Bind(userLogin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	token := authService.Login(*userLogin)
	err := userLogin.ValidasiLogin()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

}
