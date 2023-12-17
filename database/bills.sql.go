// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: bills.sql

package database

import (
	"context"
)

const createBill = `-- name: CreateBill :one
INSERT INTO bills (title, code, description, host_user_id)
VALUES (?1, ?2, ?3, ?4)
RETURNING id, host_user_id, code, title, description
`

type CreateBillParams struct {
	Title       string `json:"title"`
	Code        string `json:"code"`
	Description string `json:"description"`
	HostUserID  int64  `json:"host_user_id"`
}

func (q *Queries) CreateBill(ctx context.Context, arg CreateBillParams) (Bill, error) {
	row := q.db.QueryRowContext(ctx, createBill,
		arg.Title,
		arg.Code,
		arg.Description,
		arg.HostUserID,
	)
	var i Bill
	err := row.Scan(
		&i.ID,
		&i.HostUserID,
		&i.Code,
		&i.Title,
		&i.Description,
	)
	return i, err
}

const getBillDetail = `-- name: GetBillDetail :one
SELECT b.id, b.title, b.code, b.description, b.host_user_id, u.full_name AS user_full_name
FROM bills b
JOIN users u ON host_user_id = u.id
WHERE b.code = ?1
`

type GetBillDetailRow struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Code         string `json:"code"`
	Description  string `json:"description"`
	HostUserID   int64  `json:"host_user_id"`
	UserFullName string `json:"user_full_name"`
}

func (q *Queries) GetBillDetail(ctx context.Context, code string) (GetBillDetailRow, error) {
	row := q.db.QueryRowContext(ctx, getBillDetail, code)
	var i GetBillDetailRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Code,
		&i.Description,
		&i.HostUserID,
		&i.UserFullName,
	)
	return i, err
}

const getBillsByUser = `-- name: GetBillsByUser :many
SELECT b.id, title, description, code, (b.host_user_id = ?1) AS is_host
FROM bills b
LEFT JOIN participants p ON p.bill_id = b.id
WHERE p.user_id = ?1
`

type GetBillsByUserRow struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Code        string      `json:"code"`
	IsHost      interface{} `json:"is_host"`
}

func (q *Queries) GetBillsByUser(ctx context.Context, userID int64) ([]GetBillsByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getBillsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetBillsByUserRow{}
	for rows.Next() {
		var i GetBillsByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Code,
			&i.IsHost,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
