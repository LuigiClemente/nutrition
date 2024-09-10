package api

import (
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
		userHandler := NewHandler()
		infoRoutes.POST("/", userHandler.PostUser)
		infoRoutes.GET("/", userHandler.GetUser)
		infoRoutes.GET("/:userId", userHandler.GetUserUserId)
		infoRoutes.PUT("/:userId", userHandler.PutUserUserId)
		infoRoutes.DELETE("/:userId", userHandler.DeleteUserUserId)
	}

	mealRoutes := router.Group("api/meals")
	{
		mealHandler := NewHandler()
		mealRoutes.POST("/", mealHandler.PostMeal)
		mealRoutes.GET("/", mealHandler.GetMeal)
		mealRoutes.GET("/options", mealHandler.GetMealForOption)
		mealRoutes.GET("/:mealId", mealHandler.GetMealMealId)
		mealRoutes.PUT("/:mealId", mealHandler.PutMealMealId)
		mealRoutes.DELETE("/:mealId", mealHandler.DeleteMealMealId)
	}

	// dietaryRoutes := router.Group("api/dietary-preferences")
	// {
	// 	dietaryHandler := NewHandler()
	// 	dietaryRoutes.POST("/", dietaryHandler.PostDietaryPreferences)
	// 	dietaryRoutes.GET("/", dietaryHandler.GetDietaryPreferences)
	// 	dietaryRoutes.GET("/:userId", dietaryHandler.GetDietaryPreferencesUserId)
	// 	dietaryRoutes.PUT("/:userId", dietaryHandler.PutDietaryPreferencesUserId)
	// 	dietaryRoutes.DELETE("/:userId", dietaryHandler.DeleteDietaryPreferencesUserId)
	// }

	// environmentalFactorsRoutes := router.Group("api/environmental-factors")
	// {
	// 	environmentalHandler := NewHandler()
	// 	environmentalFactorsRoutes.POST("/", environmentalHandler.PostEnvironmentalFactors)
	// 	environmentalFactorsRoutes.GET("/", environmentalHandler.GetEnvironmentalFactors)
	// 	environmentalFactorsRoutes.GET("/:userId", environmentalHandler.GetEnvironmentalFactorsUserId)
	// 	environmentalFactorsRoutes.PUT("/:userId", environmentalHandler.PutEnvironmentalFactorsUserId)
	// 	environmentalFactorsRoutes.DELETE("/:userId", environmentalHandler.DeleteEnvironmentalFactorsUserId)
	// }

	return router
}
