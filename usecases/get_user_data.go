package usecases

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
	"github.com/mtstnt/urunan/models"
)

type GetUserDataResponseUserDTO struct {
	ID       int64
	FullName string
	Email    string
}

type GetUserDataResponse struct {
	User            GetUserDataResponseUserDTO
	Bills           []database.GetBillsByUserRow
	TotalUnpaid     int64
	TotalIncomplete int64
}

func GetUserDataHandler(c *fiber.Ctx) error {
	user := c.Context().UserValue("user").(models.User)

	// TODO: Join with participated bills and created bills.
	rows, err := database.Q.GetBillsByUser(c.Context(), user.ID)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	var userDTO GetUserDataResponseUserDTO
	if err := mapstructure.Decode(&user, &userDTO); err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	return helpers.Success(c, GetUserDataResponse{
		User:  userDTO,
		Bills: rows,
	})
}
