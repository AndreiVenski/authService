-- +goose Up
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
                   );
CREATE TABLE refresh_tokens (
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE ,
    refresh_token_id UUID PRIMARY KEY ,
    hashed_token TEXT NOT NULL,
    expires TIMESTAMP NOT NULL,
    ip_addr TEXT
);

INSERT INTO users (user_id, name, email)
VALUES (
        'b3ac3626-7e37-4026-b789-9a081b252dd1',
        'test_user1',
        'test@gmail.com'
       );

-- +goose Down
DELETE FROM users WHERE user_id = 'b3ac3626-7e37-4026-b789-9a081b252dd1';

DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS users;