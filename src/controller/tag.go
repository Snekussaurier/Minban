package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
)

func GetTags(context *gin.Context) {
	var tags []database.Tag
	if err := database.DB.Find(&tags).Error; err != nil {
		log.Fatalf("failed to query tags: %v", err)
	}

	context.JSON(http.StatusOK, tags)
}

func PostTag(context *gin.Context) {
	var request = database.Tag{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := database.DB.Create(&request).Error; err != nil {
		log.Fatalf("failed to create tag: %v", err)
	}

	context.JSON(http.StatusCreated, request)
}

func PatchTag(context *gin.Context) {
	var request = database.Tag{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := database.DB.Save(&request).Error; err != nil {
		log.Fatalf("failed to update tag: %v", err)
	}

	context.Status(http.StatusOK)
}

func DeleteTag(context *gin.Context) {
	var request = mod.NameRequest{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	if err := database.DB.Delete(&database.Tag{}, request.Name).Error; err != nil {
		log.Fatalf("failed to delete tag: %v", err)
	}
}
