CREATE TABLE users
(
    user_id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    username VARCHAR(100) NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    email VARCHAR(255),
    phone VARCHAR(15),
    profile_image VARCHAR(255),
    account_status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT null
);

CREATE TABLE todos
(
    todo_id SERIAL PRIMARY KEY,
    label VARCHAR(255) NOT NULL,
    description TEXT DEFAULT '',
    is_done BOOLEAN DEFAULT FALSE,
    creator_id INT REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT null
);