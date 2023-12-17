package usecases

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
)

type GetUserDataResponseUserDTO struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type GetUserDataResponse struct {
	User            GetUserDataResponseUserDTO   `json:"user"`
	Bills           []database.GetBillsByUserRow `json:"bills"`
	TotalUnpaid     int64                        `json:"total_unpaid"`
	TotalIncomplete int64                        `json:"total_incomplete"`
}

func GetUserDataHandler(c *fiber.Ctx, deps *Dependencies) error {
	user := c.Context().UserValue("user").(database.User)

	rows, err := deps.Q.GetBillsByUser(c.Context(), user.ID)
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
