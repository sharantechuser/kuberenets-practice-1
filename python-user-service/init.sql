CREATE DATABASE IF NOT EXISTS  userdb;

USE userdb;

CREATE TABLE user (
  user_id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50),
  password VARCHAR(50)
);

INSERT INTO user
  (username, password)
VALUES
  ('sharan-1', 'pwd123'); 