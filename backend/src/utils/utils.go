package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/mod"
)

func GetAuthenticatedUserID(c *gin.Context) (string, bool) {
	authenticatedUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusForbidden, mod.ErrorResponse{Error: "User not authenticated"})
		return "", false
	}

	userIDStr, ok := authenticatedUserID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: "Invalid user ID"})
		return "", false
	}

	return userIDStr, true
}
