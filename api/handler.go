package api

import (
	"net/http"

	"search/song/models"
	"search/song/services"
	"search/song/utils"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	term := c.Query("term")

	if term == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Term parameter is required")
		return
	}

	songs, err := services.SearchSongs(term)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to search songs")
		return
	}

	for _, song := range songs {
		err := models.SaveSong(&song)
		if err != nil {
			// Handle error, e.g., log it
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Songs found", songs)
}
