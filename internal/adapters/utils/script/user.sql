CREATE TABLE users (
    email VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    auth_tokens TEXT[],
    bot_tokens TEXT[]
);

-- Add new user to users table --
INSERT INTO users (email, password, auth_tokens, bot_tokens)
VALUES ('testemail@gmail.com', 'root', '{}', '{}');

-- Add new token to auth_tokens array --
UPDATE users
SET auth_tokens = auth_tokens || '{"newAuthToken"}'
WHERE email = 'testemail@gmail.com';

-- Remove token from auth_tokens array --
UPDATE users
SET auth_tokens = ARRAY_REMOVE(auth_tokens, 'newAuthToken')
WHERE email = 'testemail@gmail.com';

-- Drop table --
DROP TABLE users;

