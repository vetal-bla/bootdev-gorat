-- +goose Up
CREATE TABLE feed_follows (
  id uuid NOT NULL PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  user_id uuid not null REFERENCES users ON DELETE CASCADE,
  feed_id uuid not null REFERENCES feeds ON DELETE CASCADE,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(feed_id) REFERENCES feeds(id),
  UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
