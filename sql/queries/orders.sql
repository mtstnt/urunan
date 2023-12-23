-- name: GetParticipantOrders :many
SELECT
    o.id,
    o.participant_id,
    o.item_id,
    i.name AS item_name,
    i.price,
    o.qty,
    (o.qty * i.price) AS subtotal
FROM orders o
JOIN items i ON i.id = o.item_id
WHERE participant_id IN (sqlc.slice('participant_ids'));

-- name: GetAllOrders :many
SELECT
    o.id,
    o.participant_id,
    o.item_id,
    i.name AS item_name,
    i.price,
    o.qty,
    (o.qty * i.price) AS subtotal
FROM orders o
JOIN items i ON i.id = o.item_id
JOIN participants p ON p.id = o.participant_id
WHERE p.bill_id = :bill_id
ORDER BY participant_id, item_id;

-- name: CreateOrder :one
INSERT INTO orders (participant_id, item_id, qty, note)
VALUES (:participant_id, :item_id, :qty, :note)
RETURNING *;

-- name: UpdateOrder :exec
UPDATE orders
SET
    participant_id = :participant_id,
    item_id = :item_id,
    qty = :qty,
    note = :note
WHERE id = :id;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = :id;