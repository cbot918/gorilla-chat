CREATE TABLE users (
  user_id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
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
insert into user_channel (user_id, channel_id) values (2,1);
insert into user_channel (user_id, channel_id) values (1,2);
select * from user_channel;


-- select users from channel
select u.*
from users u
join user_channel uc on uc.user_id = u.user_id 
join channel c on c.channel_id = uc.channel_id
where c.channel_id = 1 ;

-- select channels from user
select c.*
from channel c
join user_channel uc on c.channel_id = uc.channel_id
join users u on u.user_id = uc.user_id
where u.user_id = 1;