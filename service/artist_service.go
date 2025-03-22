package service

import (
	"mirrorb.fun/hello-artist/dao"
	"mirrorb.fun/hello-artist/model"
)

func FindArtistByName(name string) (*[]model.DanbooruArtist, error) {
	return dao.FindArtistByName(name)
}
