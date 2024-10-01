package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"nutrition/models"
	"reflect"
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

// Helper function to determine the course type based on the number of courses
func GetCourseType(numCourses int) string {
	switch numCourses {
	case 1:
		return "single-course"
	case 2:
		return "2-course"
	case 3:
		return "3-course"
	default:
		return "multi-course" // For any other number of courses
	}
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

// GetTopMeals recommends the top N meals based on their calculated scores.
func GetTopMeals(user models.User, meals []models.Meal, topN int) []models.ScoredMeal {
	// Calculate scores for all meals
	mealsWithScores := calculateMealScore(user, meals)

	// Return the top N meals
	if topN > len(mealsWithScores) {
		topN = len(mealsWithScores) // Handle the case where there are fewer meals than topN
	}
	return mealsWithScores[:topN]
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
