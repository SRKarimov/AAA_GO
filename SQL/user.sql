CREATE TABLE users
(
	id SERIAL PRIMARY KEY,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	ip TEXT
);

INSERT INTO users
	(email, password, ip)
VALUES
	('john@mail.mail', '1Password23', '127.0.0.1');
