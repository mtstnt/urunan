package usecases

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
)

type GetBillDetailResponse struct {
	database.GetBillDetailRow
	Items []database.Item `json:"items"`
}

func GetBillDetailHandler(c *fiber.Ctx, deps *Dependencies) error {
	billCode := c.Params("code", "")
	if billCode == "" {
		return helpers.Error(c, http.StatusBadRequest, fmt.Errorf("requires bill code"))
	}

	bill, err := deps.Q.GetBillDetail(c.Context(), billCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return helpers.Error(c, http.StatusNotFound, err)
		}
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	items, err := deps.Q.GetBillItems(c.Context(), bill.ID)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	return helpers.Success(c, GetBillDetailResponse{
		GetBillDetailRow: bill,
		Items:            items,
	})
}
