CREATE TABLE users (
  user_id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE channel (
    channel_id INT AUTO_INCREMENT PRIMARY KEY,
    channel_name VARCHAR(255) NOT NULL
);

CREATE TABLE user_channel (
    user_id INT,
    channel_id INT,
    PRIMARY KEY (user_id, channel_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (channel_id) REFERENCES channel(channel_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE messages(
	id INT AUTO_INCREMENT PRIMARY KEY,
  room_id INT NOT NULL DEFAULT 0 COMMENT '0 代表私聊訊息, 群聊會有值',
  user_id INT NOT NULL,
	to_user_id INT NOT NULL DEFAULT 0 COMMENT '0 代表群聊訊息, 私訊會有值',
  content VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rooms(
  room_id INT AUTO_INCREMENT PRIMARY KEY, 
  room_name VARCHAR(255) NOT NULL
);

CREATE TABLE chatto(
  user_id INT NOT NULL, 
  chatto_id INT NOT NULL,
	PRIMARY KEY (user_id, chatto_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (chatto_id) REFERENCES users(user_id)
);

CREATE TABLE user_channel (
    user_id INT,
    channel_id INT,
    PRIMARY KEY (user_id, channel_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (channel_id) REFERENCES channel(channel_id) ON DELETE CASCADE ON UPDATE CASCADE
);