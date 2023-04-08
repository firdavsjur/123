CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    login VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    access_token VARCHAR not null
);
ALTER TABLE users
ADD CONSTRAINT login_unique UNIQUE (login);