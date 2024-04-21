-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category(
    id_category BIGINT NOT NULL,
    nama VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id_category)
)

-- +migrate StatementEnd