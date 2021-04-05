CREATE TABLE public.user (
	id serial PRIMARY KEY,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	is_active BOOLEAN NOT NULL
);

CREATE TABLE public.task (
	id serial PRIMARY KEY,
    user_id NUMERIC NOT NULL,
	title VARCHAR ( 255 ) NOT NULL,
	description VARCHAR ( 255 ) NOT NULL,
    deadline NUMERIC NOT NULL
);