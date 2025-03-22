package model

type DanbooruArtist struct {
	Artist  string
	Trigger string
	Count   int
	Url     string
}

func (artist *DanbooruArtist) TableName() string {
	return "danbooru_artist"
}
