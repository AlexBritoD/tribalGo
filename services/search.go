package services

import (
	"search/song/models"
	"search/song/resources"
)

func SearchSongs(term string) ([]models.Song, error) {
	songs, _ := resources.SearchSongsDatabase(term)

	return songs, nil
}
