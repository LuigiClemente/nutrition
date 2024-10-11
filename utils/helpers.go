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
		"Starter": {},
		"Main":    {},
		"Dessert": {},
	}
	for _, scoredMeal := range scoredMeals {
		courseMeals[scoredMeal.Meal.Course] = append(courseMeals[scoredMeal.Meal.Course], scoredMeal)
	}

	// Generate all possible combinations of starters, mains, and desserts
	starters := courseMeals["Starter"]
	mains := courseMeals["Main"]
	desserts := courseMeals["Dessert"]

	// Nested loops to generate all combinations
	for s := 0; s+numStarter <= len(starters); s++ {
		for m := 0; m+numMain <= len(mains); m++ {
			for d := 0; d+numDessert <= len(desserts); d++ {
				courseSelection := []models.MealResponse{}

				// Add starters
				for i := s; i < s+numStarter; i++ {
					courseSelection = append(courseSelection, models.MealResponse{
						ID:          starters[i].Meal.ID,
						Course:      starters[i].Meal.Course,
						Name:        starters[i].Meal.Name,
						Ingredients: mapIngredientsToIngredientResponse(starters[i].Meal.Ingredients),
						Score:       starters[i].Score,
					})
				}

				// Add mains
				for i := m; i < m+numMain; i++ {
					courseSelection = append(courseSelection, models.MealResponse{
						ID:          mains[i].Meal.ID,
						Course:      mains[i].Meal.Course,
						Name:        mains[i].Meal.Name,
						Ingredients: mapIngredientsToIngredientResponse(mains[i].Meal.Ingredients),
						Score:       mains[i].Score,
					})
				}

				// Add desserts
				for i := d; i < d+numDessert; i++ {
					courseSelection = append(courseSelection, models.MealResponse{
						ID:          desserts[i].Meal.ID,
						Course:      desserts[i].Meal.Course,
						Name:        desserts[i].Meal.Name,
						Ingredients: mapIngredientsToIngredientResponse(desserts[i].Meal.Ingredients),
						Score:       desserts[i].Score,
					})
				}

				// Add the new combination to recommendations
				recommendations = append(recommendations, models.Recommendation{
					Courses: courseSelection,
				})
			}
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
