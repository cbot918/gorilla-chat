package dao

import (
	"gorilla-chat/internal/types"
)

func (d *Dao) GetRoomMessages(RoomID int) ([]types.GetRoomMessagesParam, error) {

	var results []types.GetRoomMessagesParam

	query := `
SELECT m.*, u.name
FROM messages m
INNER JOIN users u on m.user_id = u.user_id
WHERE m.room_id = ?;
`

	err := d.DB.Select(&results, query, RoomID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *Dao) AddMessageHistory(roomID int, userID int, toUserID int, content string) error {

	query := `INSERT INTO messages (room_id, user_id, to_user_id, content) VALUES (?,?,?,?)`
	_, err := d.DB.Exec(query, roomID, userID, toUserID, content)
	if err != nil {
		return err
	}

	return nil

}
