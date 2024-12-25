package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
	"github.com/snekussaurier/minban-backend/utils"
	"gorm.io/gorm"
)

func GetCards(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	query := database.DB.Preload("Tags").Where("user_id = ?", userIDStr)

	var cards []database.Card
	if err := query.Where("user_id = ?", userIDStr).Find(&cards).Error; err != nil {
		log.Fatalf("failed to query tags: %v", err)
	}

	c.JSON(http.StatusOK, cards)
}

func PostCard(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var card database.Card

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Validate state
	var state database.State
	if err := database.DB.First(&state, "id = ? AND user_id = ?", card.StateID, userIDStr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "State with ID: " + strconv.Itoa(card.StateID) + " not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Validate tags
	for _, tagID := range card.Tags {
		var tag database.Tag
		if err := database.DB.First(&tag, "id = ? AND user_id = ?", tagID.ID, userIDStr).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "Tag with ID: " + strconv.Itoa(tagID.ID) + " not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
			return
		}
	}

	card.ID = uuid.New().String()
	card.UserID = userIDStr

	if err := database.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": card.ID})
}

func PatchCard(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	cardID := c.Param("card_id")

	var card database.Card
	if err := database.DB.First(&card, "id = ? AND user_id = ?", cardID, userIDStr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "Card not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	card.ID = cardID
	card.UserID = userIDStr

	// Validate state
	var state database.State
	if err := database.DB.First(&state, "id = ? AND user_id = ?", card.StateID, userIDStr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "State not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Validate tags
	for _, tagID := range card.Tags {
		var tag database.Tag
		if err := database.DB.First(&tag, "id = ? AND user_id = ?", tagID.ID, userIDStr).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "Tag with ID: " + strconv.Itoa(tagID.ID) + " not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
			return
		}
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
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	cardIDStr := c.Param("card_id")

	var card database.Card
	if err := database.DB.First(&card, "id = ? AND user_id = ?", cardIDStr, userIDStr).Error; err != nil {
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
