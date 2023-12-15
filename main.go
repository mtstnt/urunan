package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/usecases"
)

func main() {
	app := fiber.New()
	database.Initialize()
	database.Recreate()
	usecases.Register(app)
	app.Listen(":8080")
}
