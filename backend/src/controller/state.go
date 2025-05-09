package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
	"github.com/snekussaurier/minban-backend/utils"
)

func GetStates(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var states []database.State
	if err := database.DB.Where("user_id = ?", userIDStr).Find(&states).Error; err != nil {
		log.Fatalf("failed to query states: %v", err)
	}

	c.JSON(http.StatusOK, states)
}

func PostState(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var state = database.State{}

	if err := c.ShouldBindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	state.UserID = userIDStr

	var existingState database.State

	result := database.DB.First(&existingState, "position = ? AND user_id = ? AND name = ?", state.Position, userIDStr, state.Name)
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
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var state = database.State{}
	var stateIdStr = c.Param("state_id")

	if err := c.ShouldBindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	state.UserID = userIDStr
	stateId, err := strconv.Atoi(stateIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	state.ID = stateId

	result := database.DB.Model(&database.State{}).Where("id = ? AND user_id = ?", state.ID, userIDStr).Updates(state)
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

func DeleteState(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var stateIdStr = c.Param("state_id")
	stateId, err := strconv.Atoi(stateIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	result := database.DB.Where("id = ? AND user_id = ?", stateId, userIDStr).Delete(&database.State{})
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
