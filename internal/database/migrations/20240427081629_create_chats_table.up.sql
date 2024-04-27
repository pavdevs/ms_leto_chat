CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255),
    owner_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);