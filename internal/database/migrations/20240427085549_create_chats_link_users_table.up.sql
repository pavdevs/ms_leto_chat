CREATE TABLE chats_link_users (
                       id BIGSERIAL PRIMARY KEY,
                       chat_id INTEGER NOT NULL,
                       user_id INTEGER NOT NULL
);