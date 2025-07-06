package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
	"gorm.io/gorm"
)

func GetTags(c *gin.Context) {
	boardID := c.Param("board_id")

	var tags []database.Tag
	if err := database.DB.Where("board_id = ?", boardID).Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}

func PostTag(c *gin.Context) {
	boardID := c.Param("board_id")

	type TagCreateRequest struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	var tagRequest TagCreateRequest

	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	var existingTag database.Tag
	result := database.DB.First(&existingTag, "board_id = ? AND name = ?", boardID, tagRequest.Name)
	if result.Error == nil {
		c.JSON(http.StatusConflict, mod.ErrorResponse{Error: "Tag with this name already exists"})
		return
	}

	tag := database.Tag{
		Name:    tagRequest.Name,
		Color:   tagRequest.Color,
		BoardID: boardID,
	}

	if err := database.DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": tag.ID})
}

func PatchTag(c *gin.Context) {
	boardID := c.Param("board_id")

	type TagUpdateRequest struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	var tagRequest TagUpdateRequest
	var tagIdStr = c.Param("tag_id")

	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	result := database.DB.Model(&database.Tag{}).Where("id = ? AND board_id = ?", tagId, boardID).Updates(map[string]interface{}{
		"name":  tagRequest.Name,
		"color": tagRequest.Color,
	})
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

func PatchTags(c *gin.Context) {
	boardID := c.Param("board_id")

	type TagUpdateRequest struct {
		ID    int    `json:"id" binding:"required"`
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	var tagRequests []TagUpdateRequest

	if err := c.ShouldBindJSON(&tagRequests); err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	var err = database.DB.Transaction(func(tx *gorm.DB) error {
		for _, tagReq := range tagRequests {
			if err := tx.Model(&database.Tag{}).Where("id = ? AND board_id = ?", tagReq.ID, boardID).Updates(map[string]interface{}{
				"name":  tagReq.Name,
				"color": tagReq.Color,
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

func DeleteTag(c *gin.Context) {
	boardID := c.Param("board_id")

	var tagIdStr = c.Param("tag_id")
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: err.Error()})
		return
	}

	result := database.DB.Where("id = ? AND board_id = ?", tagId, boardID).Delete(&database.Tag{})
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
