BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(40) UNIQUE,
    password TEXT
);

CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(40),
    priority INT,
    is_completed BOOLEAN,
    -- user_id INT,
    category_id INT
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(19)
    -- user_id INT
);

-- For development only
-- ALTER TABLE todos
-- DROP CONSTRAINT IF EXISTS fk_user;

-- ALTER TABLE todos
-- ADD CONSTRAINT fk_user
--     FOREIGN KEY (user_id)
--         REFERENCES users (id);

-- For development only
-- ALTER TABLE categories
-- DROP CONSTRAINT IF EXISTS fk_user;

-- ALTER TABLE categories
-- ADD CONSTRAINT fk_user
--     FOREIGN KEY (user_id)
--         REFERENCES users (id);

-- For development only
ALTER TABLE todos
DROP CONSTRAINT IF EXISTS fk_category;

ALTER TABLE todos
ADD CONSTRAINT fk_category
    FOREIGN KEY (category_id)
        REFERENCES categories (id);

COMMIT;