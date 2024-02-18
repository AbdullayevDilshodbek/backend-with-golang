-- name: CreateUser :one
INSERT INTO users (username, role)
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: UpdateUser :one
UPDATE users 
SET username = $2
WHERE id = $1 
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;