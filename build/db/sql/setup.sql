CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30),
    email VARCHAR(30)
);

INSERT INTO users (name, email) VALUES ('user1', 'user1@sample.com');
INSERT INTO users (name, email) VALUES ('user2', 'user2@sample.com');
INSERT INTO users (name, email) VALUES ('user3', 'user3@sample.com');
INSERT INTO users (name, email) VALUES ('user4', 'user4@sample.com');
INSERT INTO users (name, email) VALUES ('user5', 'user5@sample.com');
