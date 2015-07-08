
--Create DB
CREATE DATABASE gastt_web WITH ENCODYNG ='utf8';

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name varchar(99) NOT NULL,
	image_url varchar(255) NOT NULL,
	email varchar(99) NOT NULL UNIQUE,
	id_token varchar(999) NOT NULL,
	created_at timestamp NOT NULL DEFAULT NOW(),
	updated_at timestamp DEFAULT NULL,
	deleted_at timestamp DEFAULT NULL
);

CREATE TABLE user_type(
	id SERIAL PRIMARY KEY,
	name varchar(50) NOT NULL UNIQUE
);

CREATE TABLE projects (
	id SERIAL PRIMARY KEY,
	name varchar(99) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp DEFAULT NULL,
  deleted_at timestamp DEFAULT NULL
);

CREATE TABLE projects_users (
	id SERIAL PRIMARY KEY,
	user_id integer NOT NULL REFERENCES users (id),
	project_id integer NOT NULL REFERENCES projects (id),
	user_type_id integer NOT NULL REFERENCES user_type (id)
);

CREATE TABLE languages (
  id SERIAL NOT NULL,
  code varchar(2) NOT NULL,
  language varchar(30),
  available boolean DEFAULT FALSE,
  CONSTRAINT languages_pkey
    PRIMARY KEY (id)
) WITH (OIDS = FALSE);
  
CREATE TABLE translations (
  id SERIAL PRIMARY KEY,
  language_code integer NOT NULL REFERENCES languages (code),
  data TEXT,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp DEFAULT NULL,
  deleted_at timestamp DEFAULT NULL
);