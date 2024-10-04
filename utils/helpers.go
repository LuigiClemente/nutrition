package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"nutrition/models"
	"reflect"
	"sort"
	"strconv"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

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

// GenerateCourseCombinations combines top-scored meals for requested courses.
func GenerateCourseCombinations(scoredMeals []models.ScoredMeal, numStarter, numMain, numDessert int) []models.ScoredMeal {
	// Group meals by course
	starters, mains, desserts := groupScoredMealsByCourse(scoredMeals)

	// Select top meals for each course
	topStarters := getTopMeals(starters, numStarter)
	topMains := getTopMeals(mains, numMain)
	topDesserts := getTopMeals(desserts, numDessert)

	// Combine the selected top meals into a flat list
	return combineMeals(topStarters, topMains, topDesserts)
}

// groupScoredMealsByCourse groups scored meals by their "Course" field.
func groupScoredMealsByCourse(scoredMeals []models.ScoredMeal) (starters, mains, desserts []models.ScoredMeal) {
	for _, scoredMeal := range scoredMeals {
		switch scoredMeal.Meal.Course {
		case "Starter":
			starters = append(starters, scoredMeal)
		case "Main":
			mains = append(mains, scoredMeal)
		case "Dessert":
			desserts = append(desserts, scoredMeal)
		}
	}
	return
}

// getTopMeals returns the top-scored meals up to the requested number.
func getTopMeals(meals []models.ScoredMeal, num int) []models.ScoredMeal {
	// Sort meals by score in descending order
	sort.Slice(meals, func(i, j int) bool {
		return meals[i].Score > meals[j].Score
	})

	// Return only the top requested number of meals
	if num > len(meals) {
		num = len(meals)
	}
	return meals[:num]
}

// combineMeals combines the selected top meals from starters, mains, and desserts into a single flat list.
func combineMeals(starters, mains, desserts []models.ScoredMeal) []models.ScoredMeal {
	return append(append(starters, mains...), desserts...)
}

// deleteOldRecentMeals deletes meal records older than 7 days from the recent meals table.
func DeleteOldRecentMeals(db *gorm.DB) error {
	// Calculate the cutoff date (7 days ago from now)
	cutoffDate := time.Now().AddDate(0, 0, -7)

	// Perform the delete operation
	if err := db.Where("timestamp < ?", cutoffDate).Delete(&models.RecentMeals{}).Error; err != nil {
		return fmt.Errorf("failed to delete old recent meals: %w", err)
	}

	log.Println("Deleted old recent meals successfully")
	return nil
}
