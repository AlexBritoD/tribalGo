package resources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"search/song/db"
	"search/song/models"
	"time"
)

func SearchSongsApi(term string) ([]models.Song, error) {
	baseURL := "https://itunes.apple.com/search"
	urlParams := url.Values{}
	urlParams.Set("term", term)
	urlParams.Set("entity", "song")

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = urlParams.Encode()
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Results []struct {
			TrackID   int     `json:"trackId"`
			TrackName string  `json:"trackName"`
			Artist    string  `json:"artistName"`
			Duration  int64   `json:"trackTimeMillis"`
			Album     string  `json:"collectionName"`
			Artwork   string  `json:"artworkUrl60"`
			Price     float64 `json:"collectionPrice"`
			Origin    string  `json:"country"`
			Currency  string  `json:"currency"`
		} `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	songs := make([]models.Song, len(result.Results))
	for i, item := range result.Results {
		duration, _ := time.ParseDuration(fmt.Sprintf("%dms", item.Duration))
		songs[i] = models.Song{
			ID:       item.TrackID,
			Name:     item.TrackName,
			Artist:   item.Artist,
			Duration: fmt.Sprintf("%d:%02d", int(duration.Minutes()), int(duration.Seconds())%60),
			Album:    item.Album,
			Artwork:  item.Artwork,
			Price:    fmt.Sprintf("%.2f %s", float64(item.Price), item.Currency),
			Origin:   "apple",
		}
	}

	return songs, nil
}

func SearchSongsDatabase(term string) ([]models.Song, error) {
	var songs []models.Song
	db.DB.Model(&songs).
		Where("name like ?", "%"+term+"%").
		WhereOr("album like ?", "%"+term+"%").
		WhereOr("artist like ?", "%"+term+"%").
		Select()
	return songs, nil
}
