-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at,  name, url, user_id)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
RETURNING *;

-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds;

-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE url = $1;

-- name: MarkFeedAsFetched :exec
UPDATE feeds
SET last_fetched_at = $2, updated_at = $3
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE last_fetched_at IS NULL OR last_fetched_at < $1
ORDER BY last_fetched_at ASC NULLS FIRST, updated_at ASC
LIMIT 1;