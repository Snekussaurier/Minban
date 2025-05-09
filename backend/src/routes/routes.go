package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/snekussaurier/minban-backend/controller"
	"github.com/snekussaurier/minban-backend/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.POST("/api/v1/login", controller.Login)
	router.POST("/api/v1/logout", controller.Logout)

	// Authorisation required
	authorized := router.Group("/api/v1")
	authorized.Use(middleware.AuthRequried())
	{
		// Check auth
		authorized.GET("/check-auth", controller.CheckAuth)

		// State routes
		authorized.GET("/states", controller.GetStates)
		authorized.POST("/state", controller.PostState)
		authorized.PATCH("/state/:state_id", controller.PatchState)
		authorized.DELETE("/state/:state_id", controller.DeleteState)

		// Tag routes
		authorized.GET("/tags", controller.GetTags)
		authorized.POST("/tag", controller.PostTag)
		authorized.PATCH("/tag/:tag_id", controller.PatchTag)
		authorized.DELETE("/tag/:tag_id", controller.DeleteTag)

		// Card routes
		authorized.GET("/cards", controller.GetCards)
		authorized.POST("/card", controller.PostCard)
		authorized.PATCH("/card/:card_id", controller.PatchCard)
		authorized.DELETE("/card/:card_id", controller.DeleteCard)
	}

	return router
}
