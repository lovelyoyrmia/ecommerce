-- name: CreateUser :one
INSERT INTO users (
    uid,
    email,
    first_name,
    last_name,
    password
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT 
    uid, first_name, last_name, email, password
FROM users
WHERE 
    email = $1 
LIMIT 1;