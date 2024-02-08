CREATE TABLE bots (
    bot_id VARCHAR(255) UNIQUE PRIMARY KEY,
    bot_name VARCHAR(255) NOT NULL,
    img_url VARCHAR(255),
    user_id VARCHAR(255) NOT NULL,
    bot_tokens TEXT[],
    keys TEXT[],
    FOREIGN KEY (user_id) REFERENCES users(email)
);

-- Add new bot to bots table --
INSERT INTO bots (bot_id, bot_name, img_url, bot_tokens, user_id, keys)
VALUES ('testste', 'testbot', 'https://www.google.com', '{}', 'testemail@gmail.com', '{}');