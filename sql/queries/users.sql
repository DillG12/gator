-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
	$1,
	$2,
	$3,
	$4
)
RETURNING *;

-- name: GetUser :one
SELECT id, created_at, updated_at, name
FROM users
WHERE name = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT id, created_at, updated_at, name
FROM users;

-- name: GetUserNameByFeedID :one
SELECT u.name
FROM users u
JOIN feeds f ON u.id = f.user_id
WHERE f.id = $1;


