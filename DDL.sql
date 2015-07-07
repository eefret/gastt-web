
--Create DB
CREATE DATABASE gastt_web WITH ENCODYNG ='utf8';

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name varchar(99) NOT NULL,
	image_url varchar(255) NOT NULL,
	email varchar(99) NOT NULL UNIQUE,
	id_token varchar(999) NOT NULL
);