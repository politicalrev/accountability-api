-- +goose Up
CREATE TABLE api_clients (
    key             TEXT        NOT NULL    PRIMARY KEY,
    secret          TEXT        NOT NULL,
    name            TEXT        NOT NULL
);

CREATE TABLE api_requests (
    key             TEXT        NOT NULL    REFERENCES api_clients(key),
    method          TEXT        NOT NULL,
    resource        TEXT        NOT NULL,
    requested_at    TIMESTAMP   NOT NULL    DEFAULT NOW(),
    requested_by    TEXT        NOT NULL
);

-- +goose Down
DROP TABLE api_requests;
DROP TABLE api_clients;
