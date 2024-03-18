CREATE TABLE IF NOT EXISTS mst_user (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    created_at VARCHAR(255) DEFAULT NULL,
    updated_at VARCHAR(255) DEFAULT NULL,
    deleted_at VARCHAR(255) DEFAULT NULL
);

INSERT INTO mst_user (name, email, password, created_at, updated_at, deleted_at)
VALUES 
    ('John Doe', 'john@example.com', 'password123', '2022-01-01', '2022-01-01', NULL),
    ('Jane Smith', 'jane@example.com', 'password456', '2022-01-02', '2022-01-02', NULL),
    ('Alice Johnson', 'alice@example.com', 'password789', '2022-01-03', '2022-01-03', NULL);
