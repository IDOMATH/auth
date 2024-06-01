CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR(50),
    password_hash VARCHAR(50),
    email VARCHAR(50)
);