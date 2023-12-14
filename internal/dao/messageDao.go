package dao

import (
	"gorilla-chat/internal/types"
	"gorilla-chat/internal/util"
)

func (d *Dao) GetRoomHistoryMessages(roomID int) ([]types.GetRoomMessagesParam, error) {

	var results []types.GetRoomMessagesParam

	query := `
SELECT m.*, u.name
FROM messages m
INNER JOIN users u on m.user_id = u.user_id
WHERE m.room_id = ?;
`

	err := d.DB.Select(&results, query, roomID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (d *Dao) GetUserHistoryMessages(userID int, toUserID int) ([]types.GetUserMessageHistoryParam, error) {
	var userMHistory []types.GetUserMessageHistoryParam
	query := `
select m.user_id, m.to_user_id, m.content, m.created_at, u.name
FROM messages m
JOIN users u ON m.to_user_id = u.user_id
WHERE m.user_id = ? AND m.to_user_id = ? OR m.user_id = ? AND m.to_user_id = ?;
`
	err := d.DB.Select(&userMHistory, query, userID, toUserID, toUserID, userID)
	if err != nil {
		return nil, err
	}

	util.PrintJSON(userMHistory)

	return userMHistory, nil
}

func (d *Dao) AddMessageHistory(roomID int, userID int, toUserID int, content string) error {

	query := `INSERT INTO messages (room_id, user_id, to_user_id, content) VALUES (?,?,?,?)`
	_, err := d.DB.Exec(query, roomID, userID, toUserID, content)
	if err != nil {
		return err
	}

	return nil

}
