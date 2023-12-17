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