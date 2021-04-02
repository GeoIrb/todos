CREATE TABLE public.user (
	id serial PRIMARY KEY,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	is_active BOOLEAN NOT NULL
);

CREATE TABLE public.todo (
	id serial PRIMARY KEY,
    id_user NUMERIC NOT NULL,
	title VARCHAR ( 255 ) NOT NULL,
    deadline TIME NOT NULL
);