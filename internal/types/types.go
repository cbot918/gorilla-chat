package types

type User struct {
	ID       string `db:"user_id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

type Channel struct {
	ChannelID   string `db:"channel_id"`
	ChannelName string `db:"channel_name"`
}

// type Invite struct {
// 	IID      string `db:"iid"`
// 	FromUser string `db:"from_user"`
// 	ToUser   string `db:"to_user"`
// }

// type Message struct {
// }
