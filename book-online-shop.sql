CREATE TABLE admins (
	id serial PRIMARY KEY NOT NULL UNIQUE,
	username varchar NOT NULL UNIQUE,
	password varchar NOT NULL
);

CREATE TABLE categories (
	id serial PRIMARY KEY NOT NULL UNIQUE,
	name varchar
);

CREATE SEQUENCE custom_id_seq START 1;

CREATE OR REPLACE FUNCTION generate_custom_id() RETURNS TEXT AS $$
DECLARE
    new_id TEXT;
BEGIN
    -- Get the next value from the sequence and format it with leading zeros
    new_id := 'B' || LPAD(nextval('custom_id_seq')::TEXT, 4, '0');
    RETURN new_id;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE books (
	id text PRIMARY KEY DEFAULT generate_custom_id(),
	title varchar NOT NULL,
	category_id int REFERENCES categories(id),
	price numeric(10,3) NOT NULL,
	discount int
);

CREATE TABLE addresses (
    address_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    street VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    customer_phone VARCHAR(20)
);

CREATE TABLE payment_methods (
	id SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
	is_active BOOL NOT NULL DEFAULT true,
	photo VARCHAR,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    shipping_address_id INT,
    payment_method int REFERENCES payment_methods(id) ON DELETE CASCADE,
    total_amount DECIMAL(10, 2),
    final_amount DECIMAL(10, 2),
    order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50),
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    FOREIGN KEY (shipping_address_id) REFERENCES addresses(address_id) ON DELETE SET NULL
);

CREATE TABLE order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    book_id VARCHAR(50),
    quantity INT,
    subtotal DECIMAL(10, 2)
);

SELECT * FROM books
SELECT * FROM payment_methods
SELECT * FROM orders