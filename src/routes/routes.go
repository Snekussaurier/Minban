package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/controller"
	"github.com/snekussaurier/minban-backend/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/v1/login", controller.Login)

	// Authorisation required
	authorized := router.Group("/api/v1/:user_id")
	authorized.Use(middleware.AuthRequried())
	{
		// State routes
		authorized.GET("/states", controller.GetStates)
		authorized.POST("/state", controller.PostState)
		authorized.PATCH("/state", controller.PatchState)
		authorized.DELETE("/state", controller.DeleteState)

		// Tag routes
		authorized.GET("/tags", controller.GetTags)
		authorized.POST("/tag", controller.PostTag)
		authorized.PATCH("/tag", controller.PatchTag)
		authorized.DELETE("/tag", controller.DeleteTag)

		// Card routes
		authorized.GET("/cards", controller.GetCards)
		authorized.POST("/card", controller.PostCard)
		authorized.PATCH("/card/:card_id", controller.PatchCard)
		authorized.DELETE("/card/:card_id", controller.DeleteCard)
	}

	return router
}
