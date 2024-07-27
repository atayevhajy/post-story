CREATE DATABASE postgres;


CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE followers (
                           user_id INTEGER NOT NULL,
                           follower_id INTEGER NOT NULL,
                           PRIMARY KEY (user_id, follower_id),
                           FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                           FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE stories (
                         id SERIAL PRIMARY KEY,
                         user_id INTEGER NOT NULL,
                         content TEXT NOT NULL,
                         timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


INSERT INTO users (username, email) VALUES
                                        ('user123', 'user123@example.com'),
                                        ('user456', 'user456@example.com');

INSERT INTO followers (user_id, follower_id) VALUES
                                                 (1, 2),
                                                 (2, 1);

INSERT INTO stories (user_id, content) VALUES
    (1, 'This is user123\'s first story'),
(2, 'This is user456\'s first story');

