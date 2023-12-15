// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: queries.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (full_name, email, password)
VALUES (?1, ?2, ?3)
RETURNING id, full_name, email, password
`

type CreateUserParams struct {
	FullName string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.FullName, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
    id,
    full_name,
    email,
    password
FROM users
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT
    id,
    full_name,
    email,
    password
FROM users
WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
	)
	return i, err
}
