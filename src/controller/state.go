package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
)

func GetStates(context *gin.Context) {
	var states []database.State
	if err := database.DB.Find(&states).Error; err != nil {
		log.Fatalf("failed to query states: %v", err)
	}

	context.JSON(http.StatusOK, states)
}

func PostState(context *gin.Context) {
	var request = database.State{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := database.DB.Create(&request).Error; err != nil {
		log.Fatalf("failed to create state: %v", err)
	}

	context.JSON(http.StatusCreated, request)
}

func PatchState(context *gin.Context) {
	var request = database.State{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := database.DB.Save(&request).Error; err != nil {
		log.Fatalf("failed to update state: %v", err)
	}

	context.Status(http.StatusOK)
}

func DeleteState(context *gin.Context) {
	var request = mod.IdRequest{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := database.DB.Delete(&database.State{}, request.ID).Error; err != nil {
		log.Fatalf("failed to delete state: %v", err)
	}

	context.Status(http.StatusOK)
}
