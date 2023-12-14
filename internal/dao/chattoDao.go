package dao

import (
	"gorilla-chat/internal/types"
)

func (d *Dao) GetChatTo(user int) ([]types.GetChatToParam, error) {

	var chatto []types.GetChatToParam

	query := `
select u.user_id, u.name
FROM chatto c
JOIN users u ON u.user_id = c.chatto_id
WHERE c.user_id = ?;
`

	err := d.DB.Select(&chatto, query, user)
	if err != nil {
		return nil, err
	}

	return chatto, nil
}
