package dao

import (
	"fmt"

	"mirrorb.fun/hello-artist/db"
	"mirrorb.fun/hello-artist/model"
)

func FindArtistByName(name string) (*[]model.DanbooruArtist, error) {
	var artists []model.DanbooruArtist
	result := db.DB.Find(&artists, "artist LIKE ?", "%"+name+"%")
	if result.Error != nil {
		return nil, fmt.Errorf("查询数据库失败: %w", result.Error)
	}
	return &artists, nil
}
