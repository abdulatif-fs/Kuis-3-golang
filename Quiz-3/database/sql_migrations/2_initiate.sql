-- +migrate Up
-- +migrate StatementBegin

-- CREATE TABLE category(
--     id_category BIGINT NOT NULL,
--     nama VARCHAR(256),
--     created_at date,
--     updated_at date,
--     PRIMARY KEY(id_category)
-- )

CREATE TABLE book(
    id_book BIGINT NOT NULL,
	id_category BIGINT,
    title VARCHAR(256),
    description VARCHAR(256),
    image_url VARCHAR(256),
    release_year int,
    price VARCHAR(256),
    total_page int,
    thickness VARCHAR(256),
    created_at date,
    PRIMARY KEY(id_book),
    CONSTRAINT fk_category
      FOREIGN KEY(id_category) 
	    REFERENCES category(id_category)
	    ON DELETE SET NULL
)

-- +migrate StatementEnd