CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(40) NOT NULL,
    last_name VARCHAR(40) NOT NULL,
    username VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(20) NOT NULL,
    created TIMESTAMP NOT NULL,
    salt VARCHAR(36) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    content TEXT NOT NULL,
    author BIGINT UNSIGNED NOT NULL,
    created TIMESTAMP NOT NULL,
    CONSTRAINT fk_post_author FOREIGN KEY (author) REFERENCES accounts(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    author_id BIGINT UNSIGNED NULL DEFAULT NULL,
    parent_id BIGINT UNSIGNED NULL DEFAULT NULL,
    post_id BIGINT UNSIGNED NOT NULL,
    CONSTRAINT fk_comment_author FOREIGN KEY (author_id) REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_comment_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_comment_parent FOREIGN KEY (parent_id) REFERENCES comments(id) ON DELETE CASCADE
);
