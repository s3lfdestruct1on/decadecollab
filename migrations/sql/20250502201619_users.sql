-- +goose Up
-- +goose StatementBegin
SELECT 'DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_type
            WHERE typname = 'rolename'
            AND   typtype = 'e'
        ) THEN
            CREATE TYPE rolename AS ENUM ('customer', 'manager', 'admin');
        END IF;
    END $$;
	
	CREATE TABLE IF NOT EXISTS users (
        id 			BIGSERIAL PRIMARY KEY,
        name 		VARCHAR(255) NOT NULL,
        password	VARCHAR(255) NOT NULL,
        email 		VARCHAR(255) UNIQUE NOT NULL,
		last_active TIMESTAMP,
		created_at 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		role 		ROLENAME DEFAULT 'customer'
    );`';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop table users';
-- +goose StatementEnd
