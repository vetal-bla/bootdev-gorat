-- +goose Up
CREATE TABLE feeds (
    id uuid NOT NULL PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name text NOT NULL,
    url text NOT NULL,
    user_id uuid NOT NULL REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(id),
    UNIQUE(url)
);

-- +goose Down
DROP TABLE feeds;
