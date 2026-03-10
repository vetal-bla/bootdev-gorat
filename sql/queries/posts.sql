-- name: AddPost :one
insert into posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
returning *;

-- name: GetPostsByuser :many
select posts.* from posts
join feed_follows on posts.feed_id = feed_follows.feed_id
join users on feed_follows.user_id = users.id
where users.id = $1
order by posts.published_at asc
limit $2;
