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

		// Board routes
		authorized.GET("/boards", controller.GetBoards)
		authorized.GET("/board", controller.GetBoard)
		authorized.PATCH("/board/:board_id", controller.UpdateBoard)

		// State routes
		authorized.GET("/:board_id/states", controller.GetStates)
		authorized.POST("/:board_id/state", controller.PostState)
		authorized.PATCH("/:board_id/state/:state_id", controller.PatchState)
		authorized.PATCH("/:board_id/states", controller.PatchStates)
		authorized.DELETE("/:board_id/state/:state_id", controller.DeleteState)

		// Tag routes
		authorized.GET("/:board_id/tags", controller.GetTags)
		authorized.POST("/:board_id/tag", controller.PostTag)
		authorized.PATCH("/:board_id/tag/:tag_id", controller.PatchTag)
		authorized.PATCH("/:board_id/tags", controller.PatchTags)
		authorized.DELETE("/:board_id/tag/:tag_id", controller.DeleteTag)

		// Card routes
		authorized.GET("/:board_id/cards", controller.GetCards)
		authorized.POST("/:board_id/card", controller.PostCard)
		authorized.PATCH("/:board_id/card/:card_id", controller.PatchCard)
		authorized.DELETE("/:board_id/card/:card_id", controller.DeleteCard)
	}

	return router
}
