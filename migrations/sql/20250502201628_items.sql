-- +goose Up
-- +goose StatementBegin
SELECT 'CREATE TABLE IF NOT EXISTS items (
        id 				BIGSERIAL PRIMARY KEY,
        title 			varchar(255) NOT NULL,
        price 			DECIMAL(10,2),
    	sale_percent 	SMALLINT,
        stock 			SMALLINT,
    	description 	TEXT,
      	tags 			TEXT
    );';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop table items';
-- +goose StatementEnd
