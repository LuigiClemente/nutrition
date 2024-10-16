package api

import (
	"nutrition/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// Register routes
	infoRoutes := router.Group("api/user-health-info")
	{
		userHandler := handlers.NewHandler()
		infoRoutes.POST("/", userHandler.PostUser)
		infoRoutes.GET("/", userHandler.GetUser)
		infoRoutes.GET("/:userId", userHandler.GetUserUserId)
		infoRoutes.PUT("/:userId", userHandler.PutUserUserId)
		infoRoutes.DELETE("/:userId", userHandler.DeleteUserUserId)
	}

	mealRoutes := router.Group("api/meals")
	{
		mealHandler := handlers.NewHandler()
		mealRoutes.POST("/", mealHandler.PostMeal)
		mealRoutes.GET("/", mealHandler.GetMeal)
		mealRoutes.GET("/options", mealHandler.GetMealForOption)
		mealRoutes.GET("/:mealId", mealHandler.GetMealMealId)
		mealRoutes.GET("/types", mealHandler.GetMealTypes)
		mealRoutes.PUT("/:mealId", mealHandler.PutMealMealId)
		mealRoutes.DELETE("/:mealId", mealHandler.DeleteMealMealId)
	}

	return router
}
