package repos

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mtstnt/urunan/entities"
)

type CreateBillParams struct {
	Title       string
	Code        string
	Description string
	HostUserID  int64
}

func CreateBill(ctx context.Context, db *sqlx.DB, params CreateBillParams) (int64, error) {
	var (
		sql = `
			INSERT INTO bills (title, code, description, host_user_id)
			VALUES (:title, :code, :description, :host_user_id)
			RETURNING id;
		`
		result int64
	)

	row, err := db.NamedQueryContext(ctx, sql, params)
	if err != nil {
		return result, err
	}
	err = row.Scan(&result)
	return result, err
}

func GetBillDetail(ctx context.Context, db *sqlx.DB, code string) (entities.Bill, error) {
	var (
		sql = `
			SELECT
				b.id,
				b.title,
				b.description,
				b.code,

				u.id 		AS 'host.id',
				u.full_name AS 'host.full_name',
				u.email 	AS 'host.email'
			FROM bills b
			JOIN users u ON host_user_id = u.id
			WHERE b.code = ?;
		`
		result entities.Bill
	)

	row := db.QueryRowxContext(ctx, sql, code)
	err := row.StructScan(&result)
	return result, err
}

func GetBillsByUser(ctx context.Context, db *sqlx.DB, userID int64) ([]entities.Bill, error) {
	var (
		sql = `
			SELECT
				b.id,
				b.title,
				b.description,
				b.code,

				u.id 		AS 'host.id',
				u.full_name AS 'host.full_name',
				u.email 	AS 'host.email'
			FROM bills b
			JOIN users u ON u.id = b.host_user_id
			JOIN participants p ON p.bill_id = b.id
			WHERE p.user_id = :user_id
		`
		result []entities.Bill
	)

	rows, err := db.NamedQueryContext(ctx, sql, map[string]int64{"user_id": userID})
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var b entities.Bill
		if err := rows.StructScan(&b); err != nil {
			return result, err
		}
		result = append(result, b)
	}
	return result, err
}
