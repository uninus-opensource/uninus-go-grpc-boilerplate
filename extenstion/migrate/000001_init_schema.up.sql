CREATE TABLE IF NOT EXISTS mst_user (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    created_at INT DEFAULT NULL,
    updated_at INT DEFAULT NULL,
    deleted_at INT DEFAULT NULL
    );

INSERT INTO mst_user (name, email, password, created_at, updated_at, deleted_at)
VALUES
    ('John Doe', 'john@example.com', 'password123', EXTRACT(EPOCH FROM TIMESTAMP '2022-01-01'), EXTRACT(EPOCH FROM TIMESTAMP '2022-01-01'), NULL),
    ('Jane Smith', 'jane@example.com', 'password456', EXTRACT(EPOCH FROM TIMESTAMP '2022-01-02'), EXTRACT(EPOCH FROM TIMESTAMP '2022-01-02'), NULL),
    ('Alice Johnson', 'alice@example.com', 'password789', EXTRACT(EPOCH FROM TIMESTAMP '2022-01-03'), EXTRACT(EPOCH FROM TIMESTAMP '2022-01-03'), NULL);
