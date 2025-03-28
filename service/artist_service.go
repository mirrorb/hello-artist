package service

import (
	"mirrorb.fun/hello-artist/dao"
	"mirrorb.fun/hello-artist/model"
)

func FindArtistByName(name string) (*[]model.DanbooruArtist, error) {
	return dao.FindArtistByName(name)
}

func FindCharacterByName(name string) (*[]model.DanbooruCharacter, error) {
	return dao.FindCharacterByName(name)
}

func FindResultByName(name string) ([]model.DanbooruResult, error) {
	var results []model.DanbooruResult

	// 查询画师
	artists, err := FindArtistByName(name)
	if err != nil {
		return nil, err
	}

	// 查询角色
	characters, err := FindCharacterByName(name)
	if err != nil {
		return nil, err
	}

	// 将画师数据转换为混合结果
	for _, artist := range *artists {
		results = append(results, model.DanbooruResult{
			Artist:  artist.Artist,
			Trigger: artist.Trigger,
			Count:   artist.Count,
			Url:     artist.Url,
		})
	}

	// 将角色数据转换为混合结果
	for _, character := range *characters {
		results = append(results, model.DanbooruResult{
			Character: character.Character,
			Copyright: character.Copyright,
			Trigger:   character.Trigger,
			CoreTags:  character.CoreTags,
			Count:     character.Count,
			SoloCount: character.SoloCount,
			Url:       character.Url,
		})
	}

	return results, nil
}
