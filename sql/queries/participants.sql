-- name: GetBillParticipants :many
SELECT id, full_name, email, password
FROM users
WHERE id IN (
    SELECT user_id FROM participants
    WHERE bill_id = :bill_id
);

-- name: AddParticipantToBill :one
INSERT INTO participants (bill_id, user_id, joined_at)
VALUES (:bill_id, :user_id, :joined_at)
RETURNING *;