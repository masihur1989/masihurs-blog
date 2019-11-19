-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS comments (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
    post_id INT UNSIGNED NOT NULL,
    guest BOOLEAN DEFAULT FALSE,
    name VARCHAR(25) DEFAULT NULL,
    email VARCHAR(40) DEFAULT NULL,
    user_id INT UNSIGNED NOT NULL,
    comment TEXT NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE comments;
