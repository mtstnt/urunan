package usecases

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/mtstnt/urunan/database"
	"github.com/mtstnt/urunan/helpers"
)

type GetBillDetailUserDTO struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}

type GetBillDetailItemDTO struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	InitialQty int64   `json:"initial_qty"`
}

type GetBillDetailBillDTO struct {
	Code         string                        `json:"code"`
	Title        string                        `json:"title"`
	Description  string                        `json:"description"`
	Host         GetBillDetailUserDTO          `json:"host"`
	Items        []GetBillDetailItemDTO        `json:"items"`
	Participants []GetBillDetailParticipantDTO `json:"participants"`
}

type GetBillDetailOrderDTO struct {
	ID            int64                `json:"id"`
	ParticipantID int64                `json:"participant_id"`
	Item          GetBillDetailItemDTO `json:"item"`
	Qty           int64                `json:"qty"`
}

type GetBillDetailParticipantDTO struct {
	ID     int64                   `json:"id"`
	User   GetBillDetailUserDTO    `json:"user"`
	Orders []GetBillDetailOrderDTO `json:"orders"`
}

type GetBillDetailResponse struct {
	Bill GetBillDetailBillDTO `json:"bill"`
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

	participants, err := deps.Q.GetBillParticipants(c.Context(), bill.ID)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	participantIDs := helpers.Map[database.GetBillParticipantsRow, int64](
		participants,
		func(u database.GetBillParticipantsRow) int64 {
			return u.ID
		},
	)

	orders, err := deps.Q.GetParticipantOrders(c.Context(), participantIDs)
	if err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	participantOrderMap := helpers.GroupBy[GetBillDetailOrderDTO, int64](
		helpers.Map[database.Order, GetBillDetailOrderDTO](
			orders,
			func(o database.Order) GetBillDetailOrderDTO {
				var dto GetBillDetailOrderDTO
				mapstructure.Decode(&o, &dto)
				return dto
			},
		),
		func(o GetBillDetailOrderDTO) int64 {
			return o.ParticipantID
		},
	)

	var (
		billDTO         GetBillDetailBillDTO
		itemsDTO        = make([]GetBillDetailItemDTO, 0)
		participantsDTO = make([]GetBillDetailParticipantDTO, 0)
	)

	if err := mapstructure.Decode(&bill, &billDTO); err != nil {
		return helpers.Error(c, http.StatusInternalServerError, err)
	}

	for _, item := range items {
		var itemDTO GetBillDetailItemDTO
		if err := mapstructure.Decode(&item, &itemDTO); err != nil {
			return helpers.Error(c, http.StatusInternalServerError, err)
		}
		itemsDTO = append(itemsDTO, itemDTO)
	}

	for _, participant := range participants {
		var participantDTO GetBillDetailParticipantDTO
		if err := mapstructure.Decode(&participant, &participantDTO); err != nil {
			return helpers.Error(c, http.StatusInternalServerError, err)
		}
		participantDTO.Orders = participantOrderMap[participant.ID]
		participantsDTO = append(participantsDTO, participantDTO)
	}

	userDTO := GetBillDetailUserDTO{
		ID:       bill.HostUserID,
		FullName: bill.UserFullName,
	}

	billDTO.Host = userDTO
	billDTO.Items = itemsDTO
	billDTO.Participants = participantsDTO

	response := GetBillDetailResponse{
		Bill: billDTO,
	}

	return helpers.Success(c, response)
}
