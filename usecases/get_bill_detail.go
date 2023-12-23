package usecases

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtstnt/urunan/entities"
	"github.com/mtstnt/urunan/helpers"
	"github.com/mtstnt/urunan/repos"
)

type GetBillDetailResponse struct {
	Bill entities.Bill `json:"bill"`
}

// GetBillDetailHandler returns a bill's detail based on the code.
// Will be called on joining a bill.
func GetBillDetailHandler(c *fiber.Ctx, deps *Dependencies) error {
	code := c.Params("code", "")
	if len(code) < 6 {
		return helpers.Error(c, http.StatusBadRequest, fmt.Errorf("invalid code"))
	}

	bill, err := repos.GetBillDetail(c.Context(), deps.DB, code)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	participants, err := repos.GetParticipants(c.Context(), deps.DB, bill.ID)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	participantIDs := helpers.Map[entities.Participant, int64](participants, func(p entities.Participant) int64 {
		return p.ID
	})
	participantOrders, err := repos.GetOrdersByParticipantID(c.Context(), deps.DB, participantIDs)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	for i, p := range participants {
		orders, isExist := participantOrders[p.ID]
		if isExist {
			participants[i].Orders = orders
		} else {
			participants[i].Orders = make([]entities.Order, 0)
		}
	}
	bill.Participants = participants

	return helpers.Success(c, GetBillDetailResponse{
		Bill: bill,
	})
}
