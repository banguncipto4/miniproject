package main

import (
	"MINIPROJECT/database"
	"MINIPROJECT/middlewares"
	"MINIPROJECT/routes"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	server := echo.New()

	middlewares.LogMiddleware(server)

	routes.SetupRoute(server)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	server.Logger.Fatal(server.Start(port))
}
