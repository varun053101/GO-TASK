-- name: CreateUser :one
INSERT INTO
    users (name, dob)
VALUES
    ($1, $2) RETURNING *;

-- name: GetUserByID :one
SELECT
    *
FROM
    users
WHERE
    id = $1;

-- name: UpdateUser :one
UPDATE
    users
SET
    name = $2,
    dob = $3,
    updated_at = NOW()
WHERE
    id = $1 RETURNING id,
    name,
    dob,
    created_at,
    updated_at;

-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = $1;

-- name: ListUsers :many
SELECT
    id,
    name,
    dob,
    created_at,
    updated_at
FROM
    users
ORDER BY
    id;