package usecases

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
	"github.com/mtstnt/urunan/repos"
	"github.com/mtstnt/urunan/sessions"

	_ "github.com/mattn/go-sqlite3"
)

type Dependencies struct {
	Q  database.Querier
	DB *sqlx.DB
}

type HandlerWithDependencies func(c *fiber.Ctx, d *Dependencies) error

func WithDependencies(handler HandlerWithDependencies, d *Dependencies) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler(c, d)
	}
}

func JWTMiddleware(deps *Dependencies) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) < 2 {
			return fmt.Errorf("unauthenticated")
		}

		token := authHeaderParts[1]
		userID, err := sessions.GetUserIDFromJWT(token)
		if err != nil {
			return helpers.Error(c, http.StatusUnauthorized, err)
		}

		if !sessions.Exists(token) {
			return helpers.Error(c, http.StatusUnauthorized, err)
		}

		user, err := repos.GetUserByID(c.Context(), deps.DB, userID)
		if err != nil {
			return helpers.Error(c, http.StatusUnauthorized, err)
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func Register(app *fiber.App) {
	deps := &Dependencies{
		Q:  database.Q,
		DB: sqlx.NewDb(database.DB, "sqlite3"),
	}

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	api := app.Group("/api")
	api.Post("/signup", WithDependencies(CreateUserHandler, deps))
	api.Post("/auth", WithDependencies(AuthenticateUserHandler, deps))

	protected := api.Group("", JWTMiddleware(deps))
	protected.Get("/user", WithDependencies(GetUserDataHandler, deps))
	protected.Post("/bill", WithDependencies(CreateBillHandler, deps))
	protected.Get("/bills/:code", WithDependencies(GetBillDetailHandler, deps))
}
