package routes

import (
	"MINIPROJECT/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoute(server *echo.Echo) {

	// routes for login and register
	server.POST("/user/register", controller.Register)
	server.POST("/user/login", controller.Login)

	privateRoutes := server.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))

	// routes for user

	privateRoutes.GET("/user", controller.GetAllUser)
	privateRoutes.GET("/user/:id", controller.GetByIdUser)
	privateRoutes.PUT("/user/:id", controller.UpdateUser)
	privateRoutes.DELETE("/user/:id", controller.DeleteUser)

	// routes for game

	privateRoutes.POST("/game", controller.CreateGame)
	privateRoutes.GET("/game", controller.GetAllGame)
	privateRoutes.GET("/game/:id", controller.GetByIdGame)
	privateRoutes.PUT("/game/:id", controller.UpdateGame)
	privateRoutes.DELETE("/game/:id", controller.DeleteGame)

	// routes for publisher

	privateRoutes.POST("/publisher", controller.CreatePublisher)
	privateRoutes.GET("/publisher", controller.GetAllPublisher)
	privateRoutes.GET("/publisher/:id", controller.GetByIdPublisher)
	privateRoutes.PUT("/publisher/:id", controller.UpdatePublisher)
	privateRoutes.DELETE("/publisher/:id", controller.DeletePublisher)

	// routes for rating

	privateRoutes.POST("/rating", controller.CreateRating)
	privateRoutes.GET("/rating", controller.GetAllRating)
	privateRoutes.GET("/rating/:id", controller.GetByIdRating)
	privateRoutes.PUT("/rating/:id", controller.UpdateRating)
	privateRoutes.DELETE("/rating/:id", controller.DeleteRating)
}
