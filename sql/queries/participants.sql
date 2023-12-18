-- name: GetBillParticipants :many
SELECT p.id as id, u.id AS user_id, full_name
FROM participants p
JOIN users u ON p.user_id = u.id
WHERE p.bill_id = :bill_id;

-- name: AddParticipantToBill :one
INSERT INTO participants (bill_id, user_id, joined_at)
VALUES (:bill_id, :user_id, :joined_at)
RETURNING *;

-- name: GetParticipantOrders :many
SELECT * FROM orders
WHERE participant_id IN (sqlc.slice('participant_ids'));