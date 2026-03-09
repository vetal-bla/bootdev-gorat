-- +goose Up
create table posts (
    id uuid not null primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title text not null,
    url text not null,
    description text,
    published_at timestamp not null,
    feed_id uuid no null references on delete cascade,
    foreign key(feed_id) references feeds(id)
    UNIQUE(url)
);

-- +goose Down
drop table posts;
