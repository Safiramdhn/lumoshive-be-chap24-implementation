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