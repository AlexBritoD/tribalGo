package api

import (
	"fmt"
	"net/http"

	"search/song/models"
	"search/song/services"
	"search/song/utils"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	term := c.Query("term")

	if term == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Parameter is required")
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
			fmt.Println(err)
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to save songs")
			return
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Songs found", songs)
}
