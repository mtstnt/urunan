package usecases

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
	"github.com/mtstnt/urunan/sessions"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUserResponse struct {
	Token string        `json:"token"`
	User  database.User `json:"user"`
}

func AuthenticateUserHandler(c *fiber.Ctx, deps *Dependencies) error {
	var request AuthenticateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.Error(c, http.StatusBadRequest, err)
	}

	row, err := deps.Q.GetUserByEmail(c.Context(), request.Email)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(request.Password)); err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	var user database.User
	if err := mapstructure.Decode(&row, &user); err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	token, err := sessions.Create(row.ID)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	return helpers.Success(c, AuthenticateUserResponse{
		Token: token,
		User:  user,
	})
}
