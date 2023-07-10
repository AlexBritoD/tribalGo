package models

import (
	"fmt"
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
	db.Connect()
	_, err := db.DB.Model(song).OnConflict("DO NOTHING").Insert()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
