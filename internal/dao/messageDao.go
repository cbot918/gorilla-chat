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
