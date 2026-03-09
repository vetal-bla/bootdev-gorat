-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
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
select users.name as username, feeds.name, feeds.url from feeds
join users on users.id = feeds.user_id
limit 20;

-- name: GetFeedByUrl :one
select id, created_at, updated_at, name, url, user_id
from feeds
where url = $1;

-- name: MarkFeedFetched :exec
update feeds set updated_at = now(), last_fetched_at = now()
where id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
