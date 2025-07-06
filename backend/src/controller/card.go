package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
	"gorm.io/gorm"
)

func GetCards(c *gin.Context) {
	boardID := c.Param("board_id")

	var cards []database.Card
	if err := database.DB.Preload("Tags").Where("board_id = ?", boardID).Find(&cards).Error; err != nil {
		log.Fatalf("failed to query cards: %v", err)
	}

	// Create a response structure with tag IDs instead of full tag objects
	type CardResponse struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Position    int    `json:"position"`
		StateID     int    `json:"state_id"`
		TagIDs      []int  `json:"tags"`
	}

	var response []CardResponse

	response = make([]CardResponse, 0)

	for _, card := range cards {
		cardResp := CardResponse{
			ID:          card.ID,
			Title:       card.Title,
			Description: card.Description,
			Position:    card.Position,
			StateID:     card.StateID,
			TagIDs:      make([]int, 0, len(card.Tags)),
		}

		// Extract just the tag IDs
		for _, tag := range card.Tags {
			cardResp.TagIDs = append(cardResp.TagIDs, tag.ID)
		}

		response = append(response, cardResp)
	}

	c.JSON(http.StatusOK, response)
}

func PostCard(c *gin.Context) {
	boardID := c.Param("board_id")

	type CardRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Position    int    `json:"position"`
		StateID     int    `json:"state_id"`
		TagIDs      []int  `json:"tags"` // <- das passt zu deinem JSON
	}

	var req CardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Validate state
	var state database.State
	if err := database.DB.First(&state, "id = ? AND board_id = ?", req.StateID, boardID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "State with ID: " + strconv.Itoa(req.StateID) + " not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Validate and load tags
	var tags []database.Tag
	for _, tagID := range req.TagIDs {
		var tag database.Tag
		if err := database.DB.First(&tag, "id = ? AND board_id = ?", tagID, boardID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "Tag with ID: " + strconv.Itoa(tagID) + " not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
			return
		}
		tags = append(tags, tag)
	}

	// Build Card from request + validated data
	card := database.Card{
		Title:       req.Title,
		BoardID:     boardID,
		Description: req.Description,
		Position:    req.Position,
		StateID:     req.StateID,
		Tags:        tags,
	}

	if err := database.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": card.ID})
}

func PatchCard(c *gin.Context) {
	boardID := c.Param("board_id")
	cardID := c.Param("card_id")

	type CardRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Position    int    `json:"position"`
		StateID     int    `json:"state_id"`
		TagIDs      []int  `json:"tags"`
	}

	var req CardRequest

	var card database.Card
	if err := database.DB.First(&card, "id = ? AND board_id = ?", cardID, boardID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "Card not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Validate state
	var state database.State
	if err := database.DB.First(&state, "id = ? AND board_id = ?", card.StateID, boardID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "State not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Validate tags
	var tags []database.Tag
	for _, tagID := range req.TagIDs {
		var tag database.Tag
		if err := database.DB.First(&tag, "id = ? AND board_id = ?", tagID, boardID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "Tag with ID: " + strconv.Itoa(tagID) + " not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
			return
		}
		tags = append(tags, tag)
	}

	// Build Card from request + validated data
	// Convert cardID from string to int
	cardIDInt, err := strconv.Atoi(cardID)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "Invalid card ID"})
		return
	}
	card = database.Card{
		ID:          cardIDInt,
		Title:       req.Title,
		BoardID:     boardID,
		Description: req.Description,
		Position:    req.Position,
		StateID:     req.StateID,
		Tags:        tags,
	}

	// Replacing Tags with the new ones
	if err := database.DB.Model(&card).Association("Tags").Replace(&card.Tags); err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Save the card
	if err := database.DB.Save(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func DeleteCard(c *gin.Context) {
	boardID := c.Param("board_id")
	cardIDStr := c.Param("card_id")

	var card database.Card
	if err := database.DB.First(&card, "id = ? AND board_id = ?", cardIDStr, boardID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "Card not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := database.DB.Delete(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
