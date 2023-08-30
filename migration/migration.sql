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

 SELECT 2, id FROM users LIMIT(SELECT COUNT(*) FROM users) * 0.3;




 UPDATE user_segments SET delete_time = CURRENT_TIMESTAMP::timestamp  where segment_id = (SELECT id from segments where slug = 'start')
 
;
SELECT id from segments where slug = 'start'

;
SELECT segments.id, segments.slug, segments.created_at FROM segments JOIN user_segments ON user_segments.segment_id = segments.id  WHERE user_id = 1 AND (delete_time > CURRENT_TIMESTAMP::timestamp or delete_time is NULL); 


;
 SELECT * FROM segments where slug in 


;
 INSERT INTO user_segments(user_id, segment_id)  (SELECT 10, id from segments where slug = 'start')


 SELECT 10, id from segments where slug ='start'