-- +goose Up
ALTER TABLE promises DROP COLUMN details;
CREATE TABLE moderation_queue (
    id              SERIAL                  PRIMARY KEY,
    created_at      TIMESTAMP   NOT NULL,
    politician_id   TEXT                    REFERENCES politicians(id),
    promise         TEXT        NOT NULL,
    status          TEXT        NOT NULL    CHECK (status IN ('not-started', 'in-progress', 'accomplished', 'failed')),
    status_detail   TEXT        NOT NULL,
    category        TEXT        NOT NULL    CHECK (category IN ('climate', 'culture', 'economy', 'government', 'healthcare', 'immigration', 'security')),
    source_name     TEXT        NOT NULL,
    source_link     TEXT        NOT NULL
);

-- +goose Down
DROP TABLE moderation_queue;
ALTER TABLE promises ADD COLUMN details TEXT;
