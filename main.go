package main

import (
	"MINIPROJECT/database"
	"MINIPROJECT/middlewares"
	"MINIPROJECT/routes"

	"github.com/labstack/echo/v4"
)

// koneksi lokal
func main() {
	database.Koneksi()

	server := echo.New()

	middlewares.LogMiddleware(server)

	routes.SetupRoute(server)

	server.Logger.Fatal(server.Start(":8080"))

}
