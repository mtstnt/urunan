package usecases

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/mtstnt/urunan/helpers"
	"github.com/mtstnt/urunan/models"
)

type UserDTO struct {
	ID       int64
	FullName string
	Email    string
}

type GetUserDataResponse struct {
	User UserDTO
}

func GetUserDataHandler(c *fiber.Ctx) error {
	user := c.Context().UserValue("user").(models.User)

	// TODO: Join with participated bills and created bills.

	var userDTO UserDTO
	if err := mapstructure.Decode(&user, &userDTO); err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	return helpers.Success(c, GetUserDataResponse{
		User: userDTO,
	})
}
