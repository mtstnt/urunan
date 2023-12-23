-- name: GetBillItems :many
SELECT id, bill_id, name, price, initial_qty
FROM items
WHERE bill_id = ?;

-- name: AddItemToBill :one
INSERT INTO items (name, price, initial_qty, bill_id)
VALUES (:name, :price, :initial_qty, :bill_id)
RETURNING *;

-- name: UpdateItemAtBill :one
UPDATE items
    SET
        name = :name,
        price = :price,
        initial_qty = :initial_qty
    WHERE id = :item_id
RETURNING *;

-- name: GetItemsRemainingQtyByBill :many
SELECT
    id,
    name,
    price,
    initial_qty,
    (initial_qty - purchased_qty) AS remaining_qty
FROM items
LEFT JOIN (
    SELECT o.item_id, SUM(qty) AS purchased_qty FROM orders o
    LEFT JOIN participants p ON p.id = o.participant_id
    WHERE p.bill_id = :bill_id
    GROUP BY o.item_id
) a ON a.item_id = id
WHERE bill_id = :bill_id;