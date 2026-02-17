-- +goose Up
CREATE TABLE users (
    id uuid NOT NULL PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name text NOT NULL,
    UNIQUE(name)
);

-- +goose Down
DROP TABLE users;
