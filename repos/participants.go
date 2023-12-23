package repos

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mtstnt/urunan/entities"
)

func GetParticipants(ctx context.Context, db *sqlx.DB, billID int64) ([]entities.Participant, error) {
	var (
		sql = `
			SELECT
				p.id,
				p.nickname,
				u.id 		AS 'user.id',
				u.full_name AS 'user.full_name'
			FROM participants p
			JOIN users u ON u.id = p.user_id
			WHERE bill_id = ?
		`
		result = make([]entities.Participant, 0)
	)
	rows, err := db.QueryxContext(ctx, sql, billID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var row entities.Participant
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}
