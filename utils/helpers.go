package utils

import (
	"encoding/json"
	"log"
	"math"
	"nutrition/models"
	"reflect"
	"sort"
	"strconv"

	"gorm.io/datatypes"
)

// Function to generate meal recommendations
func GenerateCourseCombinations(scoredMeals []models.ScoredMeal, numStarter, numMain, numDessert int) []models.Recommendation {
	var recommendations []models.Recommendation

	if numStarter == 0 && numMain == 0 && numDessert == 0 {
		return recommendations // No course combinations to generate
	}

	// Sort the scored meals by score in descending order
	sort.SliceStable(scoredMeals, func(i, j int) bool {
		return scoredMeals[i].Score > scoredMeals[j].Score
	})

	// Limit the scoredMeals to the top 330
	if len(scoredMeals) > 330 {
		scoredMeals = scoredMeals[:330]
	}

	// Group meals by course type
	courseMeals := map[string][]models.ScoredMeal{
		"Starter":  {},
		"Main":     {},
		"Dessert":  {},
	}
	for _, scoredMeal := range scoredMeals {
		courseMeals[scoredMeal.Meal.Course] = append(courseMeals[scoredMeal.Meal.Course], scoredMeal)
	}

	// Preallocate recommendations slice based on possible combinations
	numRecommendations := 1
	if numStarter > 0 {
		numRecommendations *= len(courseMeals["Starter"]) / numStarter
	}
	if numMain > 0 {
		numRecommendations *= len(courseMeals["Main"]) / numMain
	}
	if numDessert > 0 {
		numRecommendations *= len(courseMeals["Dessert"]) / numDessert
	}
	recommendations = make([]models.Recommendation, 0, numRecommendations)

	// Generate combinations
	for i := 0; i < numRecommendations; i++ {
		courseSelection := []models.MealResponse{}
		
		// Select starters
		for j := 0; j < numStarter && j < len(courseMeals["Starter"]); j++ {
			meal := courseMeals["Starter"][i*numStarter+j]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          meal.Meal.ID,
				Course:      meal.Meal.Course,
				Name:        meal.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(meal.Meal.Ingredients),
				Score:       meal.Score,
			})
		}
		
		// Select mains
		for k := 0; k < numMain && (i*numMain+k) < len(courseMeals["Main"]); k++ {
			meal := courseMeals["Main"][i*numMain+k]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          meal.Meal.ID,
				Course:      meal.Meal.Course,
				Name:        meal.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(meal.Meal.Ingredients),
				Score:       meal.Score,
			})
		}
		
		// Select desserts
		for l := 0; l < numDessert && l < len(courseMeals["Dessert"]); l++ {
			meal := courseMeals["Dessert"][l]
			courseSelection = append(courseSelection, models.MealResponse{
				ID:          meal.Meal.ID,
				Course:      meal.Meal.Course,
				Name:        meal.Meal.Name,
				Ingredients: mapIngredientsToIngredientResponse(meal.Meal.Ingredients),
				Score:       meal.Score,
			})
		}
		
		recommendations = append(recommendations, models.Recommendation{
			Courses: courseSelection,
		})
		
		// Early exit if no mains or desserts are required
		if numMain == 0 && numDessert == 0 {
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
