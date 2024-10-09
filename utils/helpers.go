package utils

import (
	"encoding/json"
	"log"
	"math"
	"nutrition/models"
	"reflect"
	"strconv"

	"gorm.io/datatypes"
)

// Function to generate meal recommendations
func GenerateCourseCombinations(scoredMeals []models.ScoredMeal, numStarter, numMain, numDessert int) []models.Recommendation {
	var recommendations []models.Recommendation

	// Group meals by course type
	courseMeals := make(map[string][]models.ScoredMeal)
	for _, scoredMeal := range scoredMeals {
		courseMeals[scoredMeal.Meal.Course] = append(courseMeals[scoredMeal.Meal.Course], scoredMeal)
	}

	// Cache lengths of available courses
	numStartersAvailable := len(courseMeals["Starter"])
	numMainsAvailable := len(courseMeals["Main"])
	numDessertsAvailable := len(courseMeals["Dessert"])

	// Handle case when only starters should be listed (numMain == 0 and numDessert == 0)
	if numMain == 0 && numDessert == 0 && numStarter > 0 {
		courseSelection := make([]models.MealResponse, 0)
		for starterIndex := 0; starterIndex < numStartersAvailable && starterIndex < numStarter; starterIndex++ {
			starterMeals := courseMeals["Starter"][starterIndex]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          starterMeals.Meal.ID,
				Course:      starterMeals.Meal.Course,
				Name:        starterMeals.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(starterMeals.Meal.Ingredients),
				Score:       starterMeals.Score,
			})
		}
		// Append the combination to recommendations
		recommendations = append(recommendations, models.Recommendation{
			Courses: courseSelection,
		})
		return recommendations
	}

	// Handle case when only desserts should be listed (numMain == 0 and numStarter == 0)
	if numMain == 0 && numStarter == 0 && numDessert > 0 {
		courseSelection := make([]models.MealResponse, 0)
		for dessertIndex := 0; dessertIndex < numDessertsAvailable && dessertIndex < numDessert; dessertIndex++ {
			dessertMeals := courseMeals["Dessert"][dessertIndex]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          dessertMeals.Meal.ID,
				Course:      dessertMeals.Meal.Course,
				Name:        dessertMeals.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(dessertMeals.Meal.Ingredients),
				Score:       dessertMeals.Score,
			})
		}
		// Append the combination to recommendations
		recommendations = append(recommendations, models.Recommendation{
			Courses: courseSelection,
		})
		return recommendations
	}

	// Handle case when numStarter and numDessert are 0 (only mains)
	if numStarter == 0 && numDessert == 0 && numMain > 0 {
		for mainIndex := 0; mainIndex < numMainsAvailable && mainIndex < numMain; mainIndex++ {
			courseSelection := make([]models.MealResponse, 0)
			mainMeals := courseMeals["Main"][mainIndex]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          mainMeals.Meal.ID,
				Course:      mainMeals.Meal.Course,
				Name:        mainMeals.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(mainMeals.Meal.Ingredients),
				Score:       mainMeals.Score,
			})
			// Append the combination to recommendations
			recommendations = append(recommendations, models.Recommendation{
				Courses: courseSelection,
			})
		}
		return recommendations
	}

	// Handle combinations of starters, mains, and desserts
	for mainIndex := 0; mainIndex < numMainsAvailable || numMain == 0; mainIndex++ {
		courseSelection := make([]models.MealResponse, 0)

		// Add starters
		for j := 0; j < numStarter && j < numStartersAvailable; j++ {
			meals := courseMeals["Starter"][j]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          meals.Meal.ID,
				Course:      meals.Meal.Course,
				Name:        meals.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(meals.Meal.Ingredients),
				Score:       meals.Score,
			})
		}

		// Add the main course if numMain > 0
		if numMain > 0 && mainIndex < numMainsAvailable {
			mainMeals := courseMeals["Main"][mainIndex]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          mainMeals.Meal.ID,
				Course:      mainMeals.Meal.Course,
				Name:        mainMeals.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(mainMeals.Meal.Ingredients),
				Score:       mainMeals.Score,
			})
		}

		// Add desserts
		for dessertIndex := 0; dessertIndex < numDessert && dessertIndex < numDessertsAvailable; dessertIndex++ {
			dessertMeals := courseMeals["Dessert"][dessertIndex]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          dessertMeals.Meal.ID,
				Course:      dessertMeals.Meal.Course,
				Name:        dessertMeals.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(dessertMeals.Meal.Ingredients),
				Score:       dessertMeals.Score,
			})
		}

		// Append the combination to recommendations
		recommendations = append(recommendations, models.Recommendation{
			Courses: courseSelection,
		})

		// If there are no mains and we've processed all available starters and desserts, we break the loop
		if numMain == 0 {
			break
		}
	}

	return recommendations
}

// mapIngredientsToIngredientResponse converts ingredients into IngredientResponse.
func mapIngredientsToIngredientResponse(ingredients []models.Ingredient) []models.IngredientResponse {
	var ingredientResponses []models.IngredientResponse
	for _, ingredient := range ingredients {
		ingredientResponses = append(ingredientResponses, models.IngredientResponse{
			Name:    ingredient.Name,
			Amount:  ingredient.Amount,
			Unit:    ingredient.Unit,
			Portion: ingredient.Portion,
			Ounces:  FloatToString(GramsToOunces(ingredient.Amount)),
		})
	}
	return ingredientResponses
}

// Normalize the score to a 0-100 scale.
func normalizeScore(fitScore, maxPossibleScore float64) float64 {
	return math.Max(0, math.Min(100, math.Round((fitScore/maxPossibleScore)*100*100)/100))
}

// calculateBMI calculates the BMI from height and weight.
func calculateBMI(height, weight float64) float64 {
	heightInMeters := height / 100
	return weight / (heightInMeters * heightInMeters)
}

// Helper function to compare JSON fields
func CompareJSONFields(a, b datatypes.JSON) bool {
	var aMap, bMap map[string]interface{}
	_ = json.Unmarshal(a, &aMap)
	_ = json.Unmarshal(b, &bMap)
	return reflect.DeepEqual(aMap, bMap)
}

func GramsToOunces(grams float64) float64 {
	ounces := grams / 28.35
	return math.Round(ounces*100) / 100
}

func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func handleError(err error) {
	log.Printf("Error: %v", err)
}
