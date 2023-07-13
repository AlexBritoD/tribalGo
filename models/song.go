package models

import (
	"search/song/db"
)

type Song struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Duration string `json:"duration"`
	Album    string `json:"album"`
	Artwork  string `json:"artwork"`
	Price    string `json:"price"`
	Origin   string `json:"origin"`
}

func SaveSong(song *Song) error {
	_, err := db.DB.Model(song).OnConflict("DO NOTHING").Insert()
	if err != nil {
		return err
	}
	return nil
}
