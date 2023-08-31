
CREATE TABLE IF NOT EXISTS users ( 
    id SERIAL PRIMARY KEY, 
    created_at TIMESTAMP DEFAULT current_timestamp 
); 
 
CREATE TABLE IF NOT EXISTS segments ( 
    id SERIAL PRIMARY KEY, 
    slug VARCHAR NOT NULL UNIQUE, 
    deleted INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT current_timestamp 
); 

 
 
CREATE TABLE IF NOT EXISTS user_segments ( 
    user_id INT REFERENCES users(id) ON DELETE CASCADE, 
    segment_id INT REFERENCES segments(id) ON DELETE CASCADE, 
    added_at TIMESTAMP DEFAULT current_timestamp, 
    delete_time TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (user_id, segment_id) 
); 
