package repos

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mtstnt/urunan/entities"
	"github.com/mtstnt/urunan/helpers"
)

type ParticipantOrderMap map[int64][]entities.Order

func GetOrdersByParticipantID(ctx context.Context, db *sqlx.DB, participantIDs []int64) (ParticipantOrderMap, error) {
	var (
		sql = `
			SELECT
				o.id,
				i.id AS 'item.id',
				i.name AS 'item.name',
				i.price AS 'item.price',
				i.price AS 'item.qty',
				o.qty,
				(o.qty * i.price) AS subtotal
			FROM orders o
			JOIN items i ON i.id = o.item_id
			WHERE participant_id IN (?);
		`
		result = make([]entities.Order, 0)
	)
	query, args, err := sqlx.In(sql, participantIDs)
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	rows, err := db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var row entities.Order
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	// Data transform
	ordersMap := helpers.GroupBy[entities.Order, int64](result, func(o entities.Order) int64 {
		return o.Participant.ID
	})
	return ordersMap, nil
}
