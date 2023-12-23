package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/sessions"
	"github.com/mtstnt/urunan/usecases"
)

func main() {
	app := fiber.New()
	database.Initialize()

	if len(os.Args) >= 2 && os.Args[1] == "refreshdb" {
		database.Recreate()
	}

	sessions.Initialize()
	usecases.Register(app)

	slog.Error("Error: ", app.Listen(":8080"))
}
