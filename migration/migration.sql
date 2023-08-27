-- Active: 1675774118937@@localhost@5433@postgres
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

INSERT INTO segments(id,slug,created_at) VALUES (DEFAULT, 'hello111' ,DEFAULT)  RETURNING id;
 