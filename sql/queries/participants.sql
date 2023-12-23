-- name: GetBillParticipants :many
SELECT p.id as id, u.id AS user_id, full_name, nickname
FROM participants p
LEFT JOIN users u ON p.user_id = u.id
WHERE p.bill_id = :bill_id;

-- name: AddParticipantToBill :one
INSERT INTO participants (bill_id, nickname, user_id, joined_at)
VALUES (:bill_id, :nickname, :user_id, :joined_at)
RETURNING *;