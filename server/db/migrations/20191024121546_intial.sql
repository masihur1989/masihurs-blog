-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
    name VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    remember_token VARCHAR(256) NOT NULL,
    login_type VARCHAR(10) NOT NULL,
    active BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

-- CREATE TABLE IF NOT EXISTS roles (
--     id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
--     name VARCHAR(25) NOT NULL,
--     description VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
-- )

-- CREATE TABLE IF NOT EXISTS permissions (
--     id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
--     name VARCHAR(25) NOT NULL,
--     description VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     PRIMARY KEY(id)
-- )

-- CREATE TABLE IF NOT EXISTS user_role (
--     user_id INT UNSIGNED NOT NULL,
--     role_id INT UNSIGNED NOT NULL,
--     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
--     FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
-- )

-- CREATE TABLE IF NOT EXISTS user_permission (
--     user_id INT UNSIGNED NOT NULL,
--     permission_id INT UNSIGNED NOT NULL,
--     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
--     FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
-- )

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- DROP TABLE roles;
-- DROP TABLE permissions;
-- DROP TABLE user_role;
-- DROP TABLE user_permission;
-- +goose StatementEnd
