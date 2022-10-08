CREATE TABLE IF NOT EXISTS users
(
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(15) NOT NULL,
    mail VARCHAR(30) UNIQUE
);

CREATE TABLE IF NOT EXISTS todo_lists
(
    todo_id SERIAL PRIMARY KEY,
    title VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS user_lists
(
    user_lists_id SERIAL PRIMARY KEY,
    user_id INT not null,
    todo_id INT not null,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    FOREIGN KEY (todo_id) REFERENCES todo_lists (todo_id) ON DELETE CASCADE
);