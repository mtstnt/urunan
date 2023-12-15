-- name: CreateBill :one
INSERT INTO bills (title, description, host_user_id)
VALUES (:title, :description, :host_user_id)
RETURNING *;

-- name: GetBillDetail :one
SELECT b.id, b.title, b.description, b.host_user_id, u.full_name AS user_full_name
FROM bills b
JOIN users u ON host_user_id = :user_id
WHERE b.id = :user_id;

-- name: GetBillsByUser :many
SELECT id, title, description, (b.host_user_id = :user_id) AS is_host
FROM bills b
WHERE host_user_id = :user_id
OR :user_id IN (
    SELECT user_id FROM participants
    WHERE bill_id = b.id
);