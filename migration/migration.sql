-- Active: 1691501079762@@127.0.0.1@5432@postgres
CREATE TABLE IF NOT EXISTS users ( 
    id SERIAL PRIMARY KEY, 
    created_at TIMESTAMP DEFAULT current_timestamp 
); 
 
CREATE TABLE IF NOT EXISTS segments ( 
    id SERIAL PRIMARY KEY, 
    slug VARCHAR NOT NULL UNIQUE, 
    created_at TIMESTAMP DEFAULT current_timestamp 
); 

 
 
CREATE TABLE IF NOT EXISTS user_segments ( 
    user_id INT REFERENCES users(id) ON DELETE CASCADE, 
    segment_id INT REFERENCES segments(id) ON DELETE CASCADE, 
    added_at TIMESTAMP DEFAULT current_timestamp, 
    delete_time TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (user_id, segment_id) 
); 

INSERT INTO users(id) VALUES (DEFAULT)  RETURNING id;
 
INSERT INTO user_segments(user_id, segment_id, added_at, delete_time) VALUES (1, 1, DEFAULT, DEFAULT);

 ;

 SELECT id FROM users LIMIT(SELECT COUNT(*) FROM users) * 0.3;