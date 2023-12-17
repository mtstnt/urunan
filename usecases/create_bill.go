package usecases

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
)

type CreateBillRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Items       []struct {
		Name       string  `json:"name"`
		Price      float64 `json:"price"`
		InitialQty int64   `json:"initial_qty"`
	} `json:"items"`
}

type CreateBillResponse struct {
	Bill database.Bill `json:"bill"`
}

func CreateBillHandler(c *fiber.Ctx, deps *Dependencies) error {
	user := c.Context().UserValue("user").(database.User)

	var request CreateBillRequest
	if err := c.BodyParser(&request); err != nil {
		return helpers.Error(c, http.StatusBadRequest, err)
	}

	code := strings.ToUpper(uuid.New().String())[:5]
	bill, err := deps.Q.CreateBill(c.Context(), database.CreateBillParams{
		Code:        code,
		Title:       request.Title,
		Description: request.Description,
		HostUserID:  user.ID,
	})
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	if _, err := deps.Q.AddParticipantToBill(c.Context(), database.AddParticipantToBillParams{
		BillID:   bill.ID,
		UserID:   user.ID,
		JoinedAt: time.Now().Unix(),
	}); err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	{
		tx, err := deps.DB.BeginTx(c.Context(), &sql.TxOptions{
			Isolation: sql.LevelDefault,
		})
		if err != nil {
			return helpers.Error(c, http.StatusInternalServerError, err)
		}

		q := deps.Q.(*database.Queries).WithTx(tx)
		defer tx.Rollback()
		for _, item := range request.Items {
			if _, err := q.AddItemToBill(c.Context(), database.AddItemToBillParams{
				Name:       item.Name,
				Price:      item.Price,
				InitialQty: item.InitialQty,
				BillID:     bill.ID,
			}); err != nil {
				return helpers.Error(c, http.StatusInternalServerError, err)
			}
		}

		if err := tx.Commit(); err != nil {
			return helpers.Error(c, http.StatusInternalServerError, err)
		}
	}

	return helpers.Success(c, CreateBillResponse{
		Bill: bill,
	})
}
