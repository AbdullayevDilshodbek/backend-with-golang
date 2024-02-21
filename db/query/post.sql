-- name: CreatePost :one
INSERT INTO posts (title, body, user_id, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1;

-- name: GetPostsList :many
SELECT * FROM posts
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET title = $2, body = $3, user_id = $4, status = $5
where id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE from posts
WHERE id = $1;