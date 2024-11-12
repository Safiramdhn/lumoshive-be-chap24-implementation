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
	name varchar NOT NULL,
	category_id int REFERENCES categories(id),
	price numeric(10,3) NOT NULL,
	discount int,
	author varchar,
	book_cover_url varchar,
	book_file varchar,
	quantity int
);

SELECT * FROM categories

INSERT INTO books (name, author, category_id, price, discount, quantity) VALUES 
('Marmut Merah Jambu', 'Raditya Dika', 1, 50000.000, 5, 100),
('Lima Butir Biji Jeruk', 'Arthur Conan D', 8, 105000.000, 10, 100),
('Omen', 'Lexie Xu', 8, 65000.000, 5, 100)

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
    subtotal DECIMAL(10, 2),
	rating DECIMAL(10, 2)
);


SELECT * FROM books
SELECT * FROM payment_methods
SELECT * FROM orders
SELECT * FROM admins

INSERT INTO orders (customer_id, shipping_address_id, payment_method, total_amount, final_amount, status)
VALUES 
(1, 1, 1, 50000.00, 47500.00, 'Completed'),
(2, 2, 1, 105000.00, 99750.00, 'Pending'),
(3, 3, 1, 65000.00, 63000.00, 'Shipped'),
(4, 4, 1, 50000.00, 50000.00, 'Completed'),
(5, 5, 1, 105000.00, 105000.00, 'Cancelled'),
(6, 6, 1, 65000.00, 58500.00, 'Processing'),
(7, 7, 1, 50000.00, 47500.00, 'Completed'),
(8, 8, 1, 105000.00, 99750.00, 'Pending'),
(9, 9, 1, 65000.00, 63000.00, 'Shipped'),
(10, 10, 1, 50000.00, 47500.00, 'Processing');

INSERT INTO order_items (order_id, book_id, quantity, subtotal, rating)
VALUES 
(1, 'Marmut Merah Jambu', 1, 50000.00, 4.5),
(2, 'Lima Butir Biji Jeruk', 2, 210000.00, 5.0),
(3, 'Omen', 1, 65000.00, 4.0),
(4, 'Marmut Merah Jambu', 1, 50000.00, 4.7),
(5, 'Lima Butir Biji Jeruk', 1, 105000.00, 3.8),
(6, 'Omen', 2, 130000.00, 4.5),
(7, 'Marmut Merah Jambu', 1, 50000.00, 4.0),
(8, 'Lima Butir Biji Jeruk', 1, 105000.00, 4.9),
(9, 'Omen', 1, 65000.00, 4.2),
(10, 'Marmut Merah Jambu', 2, 100000.00, 4.3);

INSERT INTO customers (customer_name, customer_phone)
VALUES 
('Alice Johnson', '123-456-7890'),
('Bob Smith', '234-567-8901'),
('Charlie Brown', '345-678-9012'),
('Daisy Miller', '456-789-0123'),
('Ethan Hunt', '567-890-1234'),
('Fiona Black', '678-901-2345'),
('George Wilson', '789-012-3456'),
('Hannah Lee', '890-123-4567'),
('Isaac Newton', '901-234-5678'),
('Julia Roberts', '012-345-6789');

INSERT INTO addresses (customer_id, street, city, postal_code, country)
VALUES 
(1, '123 Maple Street', 'Springfield', '12345', 'USA'),
(2, '456 Oak Avenue', 'Riverside', '23456', 'USA'),
(3, '789 Pine Road', 'Greenfield', '34567', 'USA'),
(4, '101 Birch Lane', 'Lakeview', '45678', 'USA'),
(5, '202 Cedar Street', 'Fairview', '56789', 'USA'),
(6, '303 Walnut Avenue', 'Brookfield', '67890', 'USA'),
(7, '404 Aspen Road', 'Hilltop', '78901', 'USA'),
(8, '505 Willow Lane', 'Elmwood', '89012', 'USA'),
(9, '606 Elm Street', 'Meadowville', '90123', 'USA'),
(10, '707 Spruce Avenue', 'Stonecrest', '01234', 'USA');