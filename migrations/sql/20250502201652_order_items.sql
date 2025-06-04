-- +goose Up
-- +goose StatementBegin
SELECT 'CREATE TABLE IF NOT EXISTS order_items (
        id 			BIGSERIAL PRIMARY KEY,
    	order_id 	BIGINT NOT NULL,
    	item_id 	BIGINT NOT NULL,
    	quantity 	INTEGER NOT NULL,
    	total_price DECIMAL(10, 2) NOT NULL,
    	FOREIGN KEY (order_id) REFERENCES orders(id),
    	FOREIGN KEY (item_id) REFERENCES items(id)
    );';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop table order_items;';
-- +goose StatementEnd
