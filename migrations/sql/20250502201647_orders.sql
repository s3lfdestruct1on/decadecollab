-- +goose Up
-- +goose StatementBegin
SELECT 'DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_type
            WHERE typname = 'status'
            AND   typtype = 'e'
        ) THEN
            CREATE TYPE status AS ENUM ('created', 'payed', 'shipping','delivered');
        END IF;
    END $$;

    CREATE TABLE IF NOT EXISTS orders (
        id BIGSERIAL PRIMARY KEY,
        user_id BIGINT NOT NULL,
        total_price DECIMAL(10, 2) NOT NULL,
        shipping_address TEXT NOT NULL,
        delivery_date TIMESTAMP,
        status status DEFAULT 'created',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop table orders;drop type status;';
-- +goose StatementEnd
