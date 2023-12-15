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



insert into users (username) values('yale');
insert into users (username) values('node');
insert into users (username) values('test');
insert into users (username) values('test2');
insert into users (username) values('test3');
select * from users;

insert into channel (channel_name) values ('channel1');
insert into channel (channel_name) values ('channel2');
insert into channel (channel_name) values ('channel3');
select * from channel;

insert into user_channel (user_id, channel_id) values (1,1);
insert into user_channel (user_id, channel_id) values (1,2);
insert into user_channel (user_id, channel_id) values (1,3);
insert into user_channel (user_id, channel_id) values (2,1);

select * from user_channel;


-- select channel-users
select u.*
from users u
join user_channel uc on uc.user_id = u.user_id 
join channel c on c.channel_id = uc.channel_id
where c.channel_id = 1 ;

-- select user-channels
select c.*
from channel c
join user_channel uc on c.channel_id = uc.channel_id
join users u on u.user_id = uc.user_id
where u.user_id = 1;


	ID        int       `db:"id"`
	RoomID    int       `db:"room_id"`
	UserID    int       `db:"user_id"`
	ToUserID  int       `db:"to_user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`



CREATE TABLE messages(
	id INT AUTO_INCREMENT PRIMARY KEY,
  room_id INT NOT NULL DEFAULT 0 COMMENT '0 代表私聊訊息, 群聊會有值',
  user_id INT NOT NULL,
	to_user_id INT NOT NULL DEFAULT 0 COMMENT '0 代表群聊訊息, 私訊會有值',
  content VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


insert into messages (room_id, user_id, to_user_id, content) 
values (1,1,0,'fist message'),(1,2,0,'second message'),(1,1,0,'third message');

insert into messages (room_id, user_id, to_user_id, content) 
values (2,1,0,'fist message'),(2,2,0,'second message');


select * from messages; 



SELECT m.*, u.name
FROM messages m
INNER JOIN users u on m.user_id = u.user_id
WHERE m.room_id = 1;


CREATE TABLE rooms(
  room_id INT AUTO_INCREMENT PRIMARY KEY, 
  room_name VARCHAR(255) NOT NULL
);

INSERT INTO rooms(room_name) VALUES('大廳'),('2群'),('3群'),('4群');
INSERT INTO rooms(room_name) VALUES('程式群');
INSERT INTO rooms(room_name) VALUES('筆記群');
drop table rooms;


CREATE TABLE chatto(
  user_id INT NOT NULL, 
  chatto_id INT NOT NULL,
	PRIMARY KEY (user_id, chatto_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (chatto_id) REFERENCES users(user_id)
);

select * from messages where content='ggg' ORDER BY id DESC ;

select m.user_id, m.to_user_id, m.content, m.created_at, u.name
FROM messages m
JOIN users u ON m.to_user_id = u.user_id
WHERE m.user_id = 1 AND m.to_user_id = 2 OR m.user_id = 2 AND m.to_user_id = 1;

delete from messages where room_id = 6;

INSERT INTO chatto() VALUES(2,1);
select * from chatto;

select u.user_id, u.name
FROM chatto c
JOIN users u ON u.user_id = c.chatto_id
WHERE c.user_id = 1;


CREATE TABLE user_channel (
    user_id INT,
    channel_id INT,
    PRIMARY KEY (user_id, channel_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (channel_id) REFERENCES channel(channel_id) ON DELETE CASCADE ON UPDATE CASCADE
);
