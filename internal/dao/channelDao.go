package dao

import (
	"fmt"
	"gorilla-chat/internal/types"
)

func (d *Dao) GetUserChannels(userID int) error {

	query := `
SELECT c.*
FROM channel c
JOIN user_channel uc ON c.channel_id = uc.channel_id
JOIN users u ON u.user_id = uc.user_id
WHERE u.user_id = ?
`

	var channels []types.Channel
	err := d.DB.Select(&channels, query, userID)
	if err != nil {
		return err
	}

	fmt.Println(channels)

	return nil
}
