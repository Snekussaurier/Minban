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

func GetTags(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var tags []database.Tag
	if err := database.DB.Where("user_id = ?", userIDStr).Find(&tags).Error; err != nil {
		log.Fatalf("failed to query tags: %v", err)
	}

	c.JSON(http.StatusOK, tags)
}

func PostTag(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var tag = database.Tag{}

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	tag.UserID = userIDStr

	var existingTag database.Tag

	result := database.DB.First(&existingTag, "user_id = ? AND name = ?", userIDStr, tag.Name)
	if result.Error == nil {
		c.JSON(http.StatusConflict, mod.ErrorResponse{Error: "Tag with this name already exists"})
		return
	}

	if err := database.DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": tag.ID})
}

func PatchTag(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var tag = database.Tag{}
	var tagIdStr = c.Param("tag_id")

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	tag.UserID = userIDStr
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	tag.ID = tagId

	result := database.DB.Model(&database.Tag{}).Where("id = ? AND user_id = ?", tag.ID, userIDStr).Updates(tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "Tag not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func DeleteTag(c *gin.Context) {
	userIDStr, ok := utils.GetAuthenticatedUserID(c)
	if !ok {
		return
	}

	var tagIdStr = c.Param("tag_id")
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	result := database.DB.Where("id = ? AND user_id = ?", tagId, userIDStr).Delete(&database.Tag{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, mod.ErrorResponse{Error: "Tag not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
