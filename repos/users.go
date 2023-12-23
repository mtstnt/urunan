package repos

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mtstnt/urunan/entities"
)

type CreateUserParams struct {
	FullName string
	Email    string
	Password string
}

func CreateUser(ctx context.Context, db *sqlx.DB, params CreateUserParams) (entities.User, error) {
	var (
		sql = `
			INSERT INTO users (full_name, email, password)
			VALUES (:full_name, :email, :password)
			RETURNING *;
		`
		result entities.User
	)

	row, err := db.NamedQueryContext(ctx, sql, params)
	if err != nil {
		return entities.User{}, err
	}
	row.Next()
	err = row.StructScan(&result)
	return result, err
}

func GetUserByEmail(ctx context.Context, db *sqlx.DB, email string) (entities.User, error) {
	var (
		sql = `
			SELECT id, full_name, email, password
			FROM users
			WHERE email = ?;
		`
		result entities.User
	)
	row := db.QueryRowxContext(ctx, sql, email)
	err := row.StructScan(&result)
	return result, err
}

func GetUserByID(ctx context.Context, db *sqlx.DB, id int64) (entities.User, error) {
	var (
		sql = `
			SELECT id, full_name, email, password
			FROM users
			WHERE id = ?;
		`
		result entities.User
	)
	row := db.QueryRowxContext(ctx, sql, id)
	err := row.StructScan(&result)
	return result, err
}
