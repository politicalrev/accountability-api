-- +goose Up
CREATE TABLE politicians (
    id              TEXT        NOT NULL    PRIMARY KEY,
    name            TEXT        NOT NULL,
    title           TEXT        NOT NULL,
    country         TEXT        NOT NULL
);

CREATE TABLE sources (
    id              SERIAL                   PRIMARY KEY,
    name            TEXT        NOT NULL,
    link            TEXT        NOT NULL
);

CREATE TABLE promises (
    id              SERIAL                  PRIMARY KEY,
    politician_id   TEXT      NOT NULL      REFERENCES politicians(id),
    name            TEXT      NOT NULL,
    details         TEXT      NOT NULL,
    category        TEXT      NOT NULL      CHECK (category IN ('climate', 'culture', 'economy', 'government', 'healthcare', 'immigration', 'security'))
);

CREATE TABLE promise_status (
    id              SERIAL                  PRIMARY KEY,
    promise_id      INT                     REFERENCES promises(id),
    name            TEXT        NOT NULL    CHECK (name IN ('not-started', 'in-progress', 'accomplished', 'failed')),
    updated_on      TIMESTAMP   NOT NULL    DEFAULT NOW(),
    detail          TEXT
);

CREATE TABLE promise_status_sources (
    source_id       INT                     REFERENCES sources(id),
    status_id       INT                     REFERENCES promise_status(id),
    PRIMARY KEY (source_id, status_id)
);

CREATE TABLE promise_history (
    promise_id      INT                     REFERENCES promises(id),
    status_id       INT                     REFERENCES promise_status(id),
    PRIMARY KEY (promise_id, status_id)
);

CREATE TABLE promise_sources (
    promise_id      INT                     REFERENCES promises(id),
    source_id       INT                     REFERENCES sources(id),
    PRIMARY KEY (promise_id, source_id)
);

-- +goose Down
DROP TABLE promise_status_sources;
DROP TABLE promise_status;
DROP TABLE promise_sources;
DROP TABLE promises;
DROP TABLE sources;
DROP TABLE politicians;
