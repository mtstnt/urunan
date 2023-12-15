package usecases

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
	"github.com/mtstnt/urunan/models"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func CreateUserHandler(c *fiber.Ctx) error {
	var request CreateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.Error(c, http.StatusBadRequest, err)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	row, err := database.Q.CreateUser(c.Context(), database.CreateUserParams{
		FullName: request.FullName,
		Email:    request.Email,
		Password: string(hashed),
	})
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	var user = models.User{}
	if err := mapstructure.Decode(&row, &user); err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	token, err := helpers.GenerateTokenFromUserID(row.ID)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	return helpers.Success(c, CreateUserResponse{
		Token: token,
		User:  user,
	})
}
