package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
	"github.com/snekussaurier/minban-backend/utils"
	"gorm.io/gorm"
)

func GetStates(c *gin.Context) {
	boardID := c.Param("board_id")

	var states []database.State
	if err := database.DB.Where("board_id = ?", boardID).Find(&states).Error; err != nil {
		log.Fatalf("failed to query states: %v", err)
	}

	c.JSON(http.StatusOK, states)
}

func PostState(c *gin.Context) {
	boardID := c.Param("board_id")

	var state = database.State{}

	if err := c.ShouldBindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	state.BoardID = boardID

	var existingState database.State

	result := database.DB.First(&existingState, "position = ? AND board_id = ? AND name = ?", state.Position, boardID, state.Name)
	if result.Error == nil {
		c.JSON(http.StatusConflict, mod.ErrorResponse{Error: "State with this position already exists"})
		return
	}

	if err := database.DB.Create(&state).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": state.ID})
}

func PatchState(c *gin.Context) {
	boardID := c.Param("board_id")

	var state = database.State{}
	var stateIdStr = c.Param("state_id")

	if err := c.ShouldBindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	state.BoardID = boardID
	stateId, err := strconv.Atoi(stateIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	state.ID = stateId

	result := database.DB.Model(&database.State{}).Where("id = ? AND board_id = ?", state.ID, boardID).Updates(state)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "State not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func PatchStates(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	type StateUpdateRequest struct {
		ID       int    `json:"id" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Position int    `json:"position" binding:"required"`
		Color    string `json:"color" binding:"required"`
	}

	var stateRequests []StateUpdateRequest

	if err := c.ShouldBindJSON(&stateRequests); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	// Get the selected board for the user
	var board database.Board
	if err := database.DB.First(&board, "user_id = ? AND selected = ?", userIDStr, true).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "No selected board found"})
			return
		}
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	var err = database.DB.Transaction(func(tx *gorm.DB) error {
		for _, stateReq := range stateRequests {
			if err := tx.Model(&database.State{}).Where("id = ? AND board_id = ?", stateReq.ID, board.ID).Updates(map[string]interface{}{
				"name":     stateReq.Name,
				"position": stateReq.Position,
				"color":    stateReq.Color,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func DeleteState(c *gin.Context) {
	boardID := c.Param("board_id")

	var stateIdStr = c.Param("state_id")
	stateId, err := strconv.Atoi(stateIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	result := database.DB.Where("id = ? AND board_id = ?", stateId, boardID).Delete(&database.State{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "State not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
