-- name: CreateBill :one
INSERT INTO bills (title, code, description, host_user_id)
VALUES (:title, :code, :description, :host_user_id)
RETURNING *;

-- name: GetBillDetail :one
SELECT b.id, b.title, b.code, b.description, b.host_user_id, u.full_name AS user_full_name
FROM bills b
JOIN users u ON host_user_id = u.id
WHERE b.code = :code;

-- name: GetBillsByUser :many
SELECT b.id, title, description, code, (b.host_user_id = :user_id) AS is_host
FROM bills b
LEFT JOIN participants p ON p.bill_id = b.id
WHERE p.user_id = :user_id