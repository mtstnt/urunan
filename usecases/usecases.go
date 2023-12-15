package usecases

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
)

func jwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) < 2 {
			return fmt.Errorf("unauthenticated")
		}

		token := authHeaderParts[1]
		userID, err := helpers.GetUserIDFromJWT(token)
		if err != nil {
			return err
		}

		user, err := database.Q.GetUserByID(c.Context(), userID)
		if err != nil {
			return err
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func Register(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	api := app.Group("/api")
	api.Post("/signup", CreateUserHandler)
	api.Post("/auth", AuthenticateUserHandler)

	protected := api.Group("", jwtMiddleware())
	protected.Get("/", GetUserDataHandler)
}
