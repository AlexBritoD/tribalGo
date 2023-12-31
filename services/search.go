package services

import (
	"fmt"
	"search/song/models"
	"search/song/resources"
)

func SearchSongs(term string) ([]models.Song, error) {
	songs, _ := resources.SearchSongsDatabase(term)
	if len(songs) == 0 {
		songs, _ = resources.SearchSongsApi(term)
		fmt.Println("viene del api")
	}

	return songs, nil
}
