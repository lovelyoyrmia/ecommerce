-- name: CreateProduct :one
INSERT INTO products (
    pid,
    name,
    description,
    category,
    price,
    stock
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetProducts :many
SELECT
    products.pid, products.name, products.description,
    product_categories.name as category_name, products.stock,
    products.price
FROM products
LEFT JOIN product_categories
ON products.category = product_categories.id
LIMIT $1
OFFSET $2;

-- name: GetCountProducts :one
SELECT COUNT(*) as count_product FROM products;

-- name: GetProductDetails :one
SELECT
    products.pid, products.name, products.description,
    product_categories.name as category_name, products.stock,
    products.price
FROM products
LEFT JOIN product_categories
ON products.category = product_categories.id
WHERE pid = $1;