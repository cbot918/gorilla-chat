package types

import "time"

type User struct {
	ID       string `db:"user_id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

type Message struct {
	ID        int     `db:"id"`
	RoomID    int     `db:"room_id"`
	UserID    int     `db:"user_id"`
	ToUserID  int     `db:"to_user_id"`
	Content   string  `db:"content"`
	CreatedAt []uint8 `db:"created_at"`
}

type Room struct {
	RoomID   int    `db:"room_id" json:"room_id"`
	RoomName string `db:"room_name" json:"room_name"`
}

type ChatTo struct {
	UserID int `db:"user_id"`
	ChatTo int `db:"chatto_id"`
}

type GetChatToParam struct {
	UserID int    `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
}

type GetRoomMessagesParam struct {
	ID        int       `db:"id"`
	RoomID    int       `db:"room_id"`
	UserID    int       `db:"user_id"`
	ToUserID  int       `db:"to_user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	Name      string    `db:"name"`
}

// type Channel struct {
// 	ChannelID   string `db:"channel_id"`
// 	ChannelName string `db:"channel_name"`
// }

// type Invite struct {
// 	IID      string `db:"iid"`
// 	FromUser string `db:"from_user"`
// 	ToUser   string `db:"to_user"`
// }

// type Message struct {
// }

type DefaultRoom struct {
	RoomID   int    `json:"room_id"`
	RoomName string `json:"room_name"`
}
