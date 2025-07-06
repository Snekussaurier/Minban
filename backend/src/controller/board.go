package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
	"github.com/snekussaurier/minban-backend/utils"
	"gorm.io/gorm"
)

func GetBoards(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	type BoardsRequest struct {
		Id       string `json:"id" binding:"required"`
		Title    string `json:"title" binding:"required"`
		Selected bool   `json:"selected" binding:"required"`
	}

	var boards []BoardsRequest
	if err := database.DB.Model(&database.Board{}).Where("user_id = ?", userIDStr).Find(&boards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, boards)
}

func GetBoard(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var board database.Board
	if err := database.DB.Preload("Cards.Tags").Preload("Tags").Preload("States").First(&board, "user_id = ? AND selected = ?", userIDStr, true).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "No selected board found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Transform the response to include only tag IDs for cards
	type CardResponse struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Position    int    `json:"position"`
		StateID     int    `json:"state_id"`
		Tags        []int  `json:"tags"`
	}

	type BoardResponse struct {
		ID          string           `json:"id"`
		Title       string           `json:"name"`
		Description string           `json:"description"`
		Token       string           `json:"token"`
		Selected    bool             `json:"selected"`
		Cards       []CardResponse   `json:"cards"`
		Tags        []database.Tag   `json:"tags"`
		States      []database.State `json:"states"`
	}

	// Transform cards to include only tag IDs
	var cardResponses []CardResponse
	for _, card := range board.Cards {
		var tagIDs []int
		for _, tag := range card.Tags {
			tagIDs = append(tagIDs, tag.ID)
		}

		cardResponses = append(cardResponses, CardResponse{
			ID:          card.ID,
			Title:       card.Title,
			Description: card.Description,
			Position:    card.Position,
			StateID:     card.StateID,
			Tags:        tagIDs,
		})
	}

	response := BoardResponse{
		ID:          board.ID,
		Title:       board.Title,
		Description: board.Description,
		Token:       board.Token,
		Selected:    board.Selected,
		Cards:       cardResponses,
		Tags:        board.Tags,
		States:      board.States,
	}

	c.JSON(http.StatusOK, response)
}

func UpdateBoard(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	boardID := c.Param("board_id")

	var board database.Board
	if err := database.DB.First(&board, "id = ? AND user_id = ?", boardID, userIDStr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "Board not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	type BoardRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		Token       string `json:"token" binding:"required"`
		Selected    bool   `json:"selected" binding:"required"`
	}

	var req BoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	board.Title = req.Name
	board.Token = req.Token
	board.Selected = req.Selected
	board.Description = req.Description

	if err := database.DB.Save(&board).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
