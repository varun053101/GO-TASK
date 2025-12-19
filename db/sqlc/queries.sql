-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;