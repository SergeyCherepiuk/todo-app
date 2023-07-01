BEGIN;

CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    title VARCHAR(40),
    category_id INT,
    priority INT,
    is_completed BOOLEAN
);

CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(19)
);

-- For development only
ALTER TABLE todo
DROP CONSTRAINT IF EXISTS fk_category;

ALTER TABLE todo
ADD CONSTRAINT fk_category
    FOREIGN KEY (category_id)
        REFERENCES category (id);

COMMIT;