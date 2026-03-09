-- +goose Up
ALTER TABLE feeds 
ADD COLUMN last_fetched_at timestamp DEFAULT NULL;

-- +goose Down
alter table feeds
drop column last_fetchet_at;
