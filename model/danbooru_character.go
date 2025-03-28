package model

type DanbooruCharacter struct {
	Character string
	Copyright string
	Trigger   string
	CoreTags  string
	Count     int
	SoloCount int
	Url       string
}

func (character *DanbooruCharacter) TableName() string {
	return "danbooru_character"
}
