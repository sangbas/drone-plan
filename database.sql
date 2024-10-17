-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

CREATE TABLE estate (
	id UUID NOT NULL,
	length INT NOT NULL,
	width INT NOT NULL,
	CONSTRAINT estate_pkey PRIMARY KEY (id)
);

CREATE TABLE tree (
	id UUID NOT NULL,
	estate_id UUID NOT NULL,
	x_axis INT NOT NULL,
	y_axis INT NOT NULL,
	heigth INT NOT NULL,
	CONSTRAINT tree_pkey PRIMARY KEY (id),
	CONSTRAINT tree_estate_fk FOREIGN KEY (estate_id) REFERENCES estate(id)
);