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

	privateRoutes.POST("/game", controller.CreatePublisher)
	privateRoutes.GET("/game", controller.GetAllPublisher)
	privateRoutes.GET("/game/:id", controller.GetByIdPublisher)
	privateRoutes.PUT("/game/:id", controller.UpdatePublisher)
	privateRoutes.DELETE("/game/:id", controller.DeletePublisher)

	// routes for rating

	privateRoutes.POST("/game", controller.CreateRating)
	privateRoutes.GET("/game", controller.GetAllRating)
	privateRoutes.GET("/game/:id", controller.GetByIdRating)
	privateRoutes.PUT("/game/:id", controller.UpdateRating)
	privateRoutes.DELETE("/game/:id", controller.DeleteRating)
}
