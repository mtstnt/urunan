-- name: CreateUser :one
INSERT INTO users (full_name, email, password)
VALUES (:full_name, :email, :password)
RETURNING *;

-- name: GetUserByEmail :one
SELECT id, full_name, email, password
FROM users
WHERE email = ?;

-- name: GetUserByID :one
SELECT id, full_name, email, password
FROM users
WHERE id = ?;