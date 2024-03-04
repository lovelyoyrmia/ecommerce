-- name: CreateOrder :one
INSERT INTO orders (
    oid,
    uid,
    order_status,
    total_amount
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CreateOrderItems :exec
INSERT INTO order_items (
    oid,
    pid,
    quantity,
    amount
) VALUES (
    $1, $2, $3, $4
);

-- name: GetCart :many
SELECT 
    oid, uid, total_amount, ordered_at
FROM orders
WHERE uid = $1 AND order_status = 0
LIMIT $2
OFFSET $3;

-- name: GetCartUser :one
SELECT 
    oid, uid, total_amount, ordered_at
FROM orders
WHERE uid = $1 AND order_status = 0
LIMIT 1;

-- name: GetCartProducts :many
SELECT 
    oid, pid, quantity, amount
FROM order_items
WHERE oid = $1;

-- name: GetCartProductDetail :one
SELECT 
    oid,
    pid,
    quantity,
    amount 
FROM order_items 
WHERE oid = $1 AND pid = $2;

-- name: UpdateCartProductDetail :exec
UPDATE order_items
SET
    quantity = $1,
    amount = $2
WHERE oid = $3 AND pid = $4;

-- name: GetOrderDetails :one
SELECT 
    oid, uid, total_amount, ordered_at
FROM orders
WHERE oid = $1 AND uid = $2
LIMIT 1;

-- name: DeleteOrderItems :exec
DELETE FROM order_items 
WHERE oid = $1 AND pid = $2 AND order_status <> 0;

-- name: GetOrderItems :many
SELECT
    order_items.oid, products.pid, products.name, 
    products.price, products.stock, 
    order_items.quantity, order_items.amount
FROM order_items
LEFT JOIN products
ON order_items.pid = products.pid
WHERE order_items.oid = $1
LIMIT $2
OFFSET $3;

-- name: UpdateCart :exec
UPDATE orders
SET
    total_amount = $1
WHERE oid = $2;


-- name: DeleteOrderItemByProduct :exec
DELETE FROM order_items
WHERE oid = $1 AND pid = $2;

-- name: DeleteOrder :exec
DELETE FROM orders as O
WHERE EXISTS (
    SELECT oid FROM order_items
    WHERE O.oid = oid
);