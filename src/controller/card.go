package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
	"gorm.io/gorm"
)

func GetCards(c *gin.Context) {
	var cards []database.Card

	userID := c.Param("user_id")
	query := database.DB.Preload("Tags")

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cards)
}

func PostCard(c *gin.Context) {
	var card database.Card

	userID := c.Param("user_id")

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	card.ID = uuid.New().String()
	card.UserID = userID

	if err := database.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": card.ID})
}

func PatchCard(c *gin.Context) {
	userId := c.Param("user_id")
	cardId := c.Param("card_id")

	var card database.Card

	if err := database.DB.First(&card, "id = ? AND user_id = ?", cardId, userId).Error; err != nil {
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

	if err := database.DB.Save(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteCard(c *gin.Context) {
	userId := c.Param("user_id")
	cardId := c.Param("card_id")

	result := database.DB.Delete(&database.Card{}, "id = ? AND user_id = ?", cardId, userId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "Card not found"})
		return
	}

	c.Status(http.StatusOK)
}
