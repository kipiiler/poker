-- Create Table for games --
CREATE TABLE games (
    BotID   VARCHAR(225)[],
    Started BOOLEAN,
    Ended   BOOLEAN,
    Winner  VARCHAR(225),
    History VARCHAR(225)[],
    Config  NULL,
    GameID  VARCHAR(225)
);

-- Create Table for game state --
CREATE TABLE gameState (
    GameStateID VARCHAR(225),
    GameID      VARCHAR(225),
    Pot         INTEGER,
    Bets        VARCHAR(225),
    SmallBlind  INTEGER,
    BigBlind    INTEGER,
    Middle      VARCHAR(225)[],
    Folded      VARCHAR(225),
    Decision    VARCHAR(225),
    Hand        VARCHAR(225),
    IsRunning   BOOLEAN,
    Turn        INTEGER,
    Around      INTEGER,
    Deck        VARCHAR(225)[]
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

