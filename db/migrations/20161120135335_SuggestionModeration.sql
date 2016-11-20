-- +goose Up
ALTER TABLE moderation_queue
ADD COLUMN accepted_at TIMESTAMP,
ADD COLUMN accepted_by TEXT,
ADD COLUMN deleted_at TIMESTAMP,
ADD COLUMN deleted_by TEXT;

DROP TABLE promise_history; -- This table was unused

-- +goose Down
ALTER TABLE moderation_queue
DROP COLUMN accepted_at,
DROP COLUMN accepted_by,
DROP COLUMN deleted_at,
DROP COLUMN deleted_by;
