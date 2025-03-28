package dao

import (
	"fmt"

	"mirrorb.fun/hello-artist/db"
	"mirrorb.fun/hello-artist/model"
)

func FindCharacterByName(name string) (*[]model.DanbooruCharacter, error) {
	var characters []model.DanbooruCharacter
	result := db.DB.Find(&characters, "character LIKE ?", "%"+name+"%")
	if result.Error != nil {
		return nil, fmt.Errorf("查询数据库失败: %w", result.Error)
	}
	return &characters, nil
}
