-- +goose Up
-- +goose StatementBegin
SELECT 'CREATE TABLE IF NOT EXISTS baskets (
        id 			BIGSERIAL PRIMARY KEY,
        user_id 	BIGINT NOT NULL,
        item_id 	BIGINT,
        quantity 	SMALLINT,
    	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
      	FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
    );';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop table baskets';
-- +goose StatementEnd
