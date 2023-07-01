CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    title VARCHAR(40),
    category_id INT,
    priority INT,
    is_completed BOOLEAN,
    CONSTRAINT fk_category
        FOREIGN KEY (category_id)
            REFERENCES category(id)
);

CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(20)
)