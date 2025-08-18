CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(64) PRIMARY KEY,
    customer_id VARCHAR(64) NOT NULL,
    customer_name VARCHAR(128) NOT NULL,
    status INT NOT NULL,
    total_amount NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- You can add a separate table for order items if needed
-- CREATE TABLE IF NOT EXISTS order_items (
--     id SERIAL PRIMARY KEY,
--     order_id VARCHAR(64) REFERENCES orders(id),
--     product_id VARCHAR(64) NOT NULL,
--     product_name VARCHAR(128) NOT NULL,
--     quantity INT NOT NULL,
--     unit_price NUMERIC(12,2) NOT NULL,
--     total_price NUMERIC(12,2) NOT NULL
-- );
