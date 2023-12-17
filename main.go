package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/sessions"
	"github.com/mtstnt/urunan/usecases"
)

func main() {

	app := fiber.New()

	database.Initialize()

	fmt.Printf("%+v\n", os.Args)
	if len(os.Args) >= 2 {
		database.Recreate()
	}

	sessions.Initialize()

	usecases.Register(app)
	app.Listen(":8080")
}
