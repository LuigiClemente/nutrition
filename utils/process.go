package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"nutrition/models"
	"sort"
	"time"
)

// CalculateMealScore computes scores for meals based on user preferences and goals.
func CalculateMealScore(user models.User, meals []models.Meal) []models.ScoredMeal {
	mealsWithScores := make([]models.ScoredMeal, len(meals))

	// Pre-calculate user's BMI for health goal-related calculations.
	userBMI := user.BodyMetrics.Weight / ((user.BodyMetrics.Height / 100) * (user.BodyMetrics.Height / 100))

	// Constants for max possible scores for normalization.
	const maxTotalScore = 155.0

	for i, meal := range meals {
		fitScore := 0.0
		var nutritionalContent, healthScores map[string]float64

		// Unmarshal JSONB fields once per meal for nutritional content and health scores.
		if err := json.Unmarshal(meal.NutritionalContent, &nutritionalContent); err != nil {
			fmt.Println("Error unmarshaling nutritional content:", err)
			continue
		}
		if err := json.Unmarshal(meal.HealthScores, &healthScores); err != nil {
			fmt.Println("Error unmarshaling health scores:", err)
			continue
		}

		// Add the calculated score from each component.
		fitScore += calculateDietaryMatch(user, meal)
		fitScore += calculateNutritionalMatch(user, meal)
		fitScore += calculateHealthGoalsAlignment(user, nutritionalContent, healthScores, userBMI)
		fitScore += calculateNutritionalImpact(user, healthScores)
		fitScore += calculateMicrobiomeCompatibility(user, meal)
		fitScore += calculateEnvironmentalAdaptability(user, meal)
		fitScore += calculateRecentConsumptionPenalty(user, meal)
		fitScore += calculateAgeGenderScore(user, meal, nutritionalContent)

		// Normalize the score to be out of 100 and round to 2 decimal places.
		normalizedScore := normalizeScore(fitScore, maxTotalScore)
		mealsWithScores[i] = models.ScoredMeal{Meal: meal, Score: normalizedScore}
	}

	// Sort meals by normalized score in descending order.
	sort.Slice(mealsWithScores, func(i, j int) bool {
		return mealsWithScores[i].Score > mealsWithScores[j].Score
	})

	return mealsWithScores
}

// Normalize the score to a 0-100 scale.
func normalizeScore(fitScore, maxPossibleScore float64) float64 {
	return math.Max(0, math.Min(100, math.Round((fitScore/maxPossibleScore)*100*100)/100))
}

// calculateDietaryMatch scores meals based on user's dietary preferences and avoidances.
func calculateDietaryMatch(user models.User, meal models.Meal) float64 {
	score := 0.0
	for _, tag := range meal.Tags {
		switch tag {
		case "vegetarian":
			if user.DietaryPreferences.Vegetarian {
				score += 10.0
			}
		case "vegan":
			if user.DietaryPreferences.Vegan {
				score += 10.0
			}
		case "gluten-free":
			if user.DietaryPreferences.GlutenFree {
				score += 10.0
			}
		case "dairy-free":
			if user.DietaryPreferences.DairyFree {
				score += 10.0
			}
		}
	}
	return score
}

// calculateNutritionalMatch scores meals based on how well they match the user's nutritional deficiencies.
func calculateNutritionalMatch(user models.User, meal models.Meal) float64 {
	score := 0.0
	for _, deficiency := range user.NutritionalDeficiencies {
		if containsNutrient(meal.Ingredients, deficiency) {
			score += 10.0 // Increase score for every deficiency matched
		}
	}
	return score
}

// calculateHealthGoalsAlignment scores meals based on how well they align with user's health goals.
func calculateHealthGoalsAlignment(user models.User, nutritionalContent, healthScores map[string]float64, bmi float64) float64 {
	score := 0.0

	for _, goal := range user.Goals {
		switch goal.Type {
		case "Weight loss":
			if nutritionalContent["calories"] <= 400.0 {
				score += 30.0
			}
			if bmi >= 25 && nutritionalContent["calories"] <= 400.0 {
				score += 30.0
			} else if bmi < 25 && nutritionalContent["calories"] <= 500.0 {
				score += 20.0
			}
			if nutritionalContent["sugar"] <= 5.0 {
				score += 20.0
			}
			if nutritionalContent["fiber"] >= 5.0 {
				score += 10.0
			}

		case "Muscle gain":
			if nutritionalContent["protein"] >= 25.0 {
				score += 30.0
			}
			if nutritionalContent["calories"] >= 500.0 {
				score += 20.0
			}
			if nutritionalContent["fat"] >= 10.0 {
				score += 10.0
			}
			if nutritionalContent["protein"] >= 25.0 && nutritionalContent["calories"] <= 600.0 {
				score += 30.0
			}

		case "Heart health":
			if healthScores["heart_healthy"] >= 8.0 {
				score += 30.0
			}
			if nutritionalContent["cholesterol"] <= 20.0 {
				score += 10.0
			}
			if nutritionalContent["fat"] <= 10.0 {
				score += 10.0
			}
			if nutritionalContent["fiber"] >= 5.0 {
				score += 10.0
			}

		case "Blood sugar management":
			if healthScores["diabetes_friendly"] >= 8.0 {
				score += 30.0
			}
			if nutritionalContent["sugar"] <= 5.0 {
				score += 20.0
			}
			if nutritionalContent["fiber"] >= 8.0 {
				score += 10.0
			}

		case "Gut health":
			if nutritionalContent["fiber"] >= 8.0 {
				score += 30.0
			}

		case "Metabolic health":
			if nutritionalContent["sugar"] <= 5.0 && nutritionalContent["calories"] <= 500.0 {
				score += 30.0
			}
			if nutritionalContent["protein"] >= 15.0 && nutritionalContent["fat"] <= 15.0 {
				score += 20.0
			}

		case "Cholesterol reduction":
			if nutritionalContent["cholesterol"] <= 20.0 {
				score += 30.0
			}
			if nutritionalContent["fiber"] >= 5.0 {
				score += 20.0
			}

		case "Weight maintenance":
			if nutritionalContent["calories"] >= 400.0 && nutritionalContent["calories"] <= 600.0 {
				score += 30.0
			}
			if nutritionalContent["fat"] <= 15.0 && nutritionalContent["sugar"] <= 5.0 {
				score += 20.0
			}
			if nutritionalContent["protein"] >= 15.0 {
				score += 10.0
			}

		case "Lean muscle maintenance":
			if nutritionalContent["protein"] >= 15.0 {
				score += 30.0
			}
			if nutritionalContent["fiber"] >= 5.0 {
				score += 10.0
			}

		case "Improved energy levels":
			if nutritionalContent["carbs"] >= 40.0 {
				score += 30.0
			}
			if nutritionalContent["calories"] >= 400.0 && nutritionalContent["calories"] <= 600.0 {
				score += 20.0
			}
			if nutritionalContent["sugar"] <= 5.0 {
				score += 10.0
			}

		case "Endurance training support":
			if nutritionalContent["carbs"] >= 40.0 {
				score += 30.0
			}
			if nutritionalContent["protein"] >= 15.0 {
				score += 20.0
			}
			if nutritionalContent["fat"] <= 15.0 {
				score += 10.0
			}

		case "Cardiovascular fitness":
			if nutritionalContent["sodium"] <= 200.0 {
				score += 30.0
			}
			if nutritionalContent["cholesterol"] <= 20.0 && nutritionalContent["fat"] <= 10.0 {
				score += 20.0
			}
			if nutritionalContent["antioxidants"] >= 5.0 {
				score += 10.0
			}

		case "Detoxification":
			if nutritionalContent["fiber"] >= 8.0 {
				score += 30.0
			}

		case "Bone health":
			if nutritionalContent["calcium"] >= 200.0 {
				score += 30.0
			}

		case "Skin health":
			if nutritionalContent["vitamin A"] >= 10.0 && nutritionalContent["vitamin C"] >= 20.0 {
				score += 30.0
			}
			if nutritionalContent["healthy fats"] >= 10.0 {
				score += 20.0
			}

		case "Anti-inflammatory diet":
			if nutritionalContent["antioxidants"] >= 5.0 {
				score += 30.0
			}

		case "Hormonal balance":
			if nutritionalContent["omega-3"] >= 10.0 {
				score += 30.0
			}
			if nutritionalContent["protein"] >= 15.0 && nutritionalContent["fiber"] >= 5.0 {
				score += 20.0
			}

		case "Improved mental clarity":
			if nutritionalContent["omega-3"] >= 10.0 {
				score += 30.0
			}
			if nutritionalContent["antioxidants"] >= 5.0 {
				score += 20.0
			}

		}
	}
	return score
}

// calculateNutritionalImpact adjusts the score based on how meals impact the user's health monitoring (e.g., blood sugar, cholesterol).
func calculateNutritionalImpact(user models.User, healthScores map[string]float64) float64 {

	score := 0.0
	if user.BloodGlucose > 100 && healthScores["diabetes_friendly"] >= 8.0 {
		score += 15.0 // Prioritize meals that help with blood sugar management
	}
	if user.LipidProfile.LDL > 100 && healthScores["heart_healthy"] >= 8.0 {
		score += 15.0 // Prioritize heart-healthy meals for users with high cholesterol
	}
	return score
}

// calculateMicrobiomeCompatibility scores meals based on how well they support the user's microbiome.
func calculateMicrobiomeCompatibility(user models.User, meal models.Meal) float64 {
	for _, recommendation := range user.MicrobiomeData.GutHealthRecommendations {
		if contains(meal.Tags, recommendation) {
			return 10.0 // Increase score for gut-health-promoting meals
		}
	}
	return 0.0
}

// calculateEnvironmentalAdaptability scores meals based on their environmental impact and user’s preferences.
func calculateEnvironmentalAdaptability(user models.User, meal models.Meal) float64 {
	score := 0.0
	if user.EnvironmentalFactors.Season == "Winter" && contains(meal.Tags, "warm") {
		score += 5.0 // Increase score for warm foods in cold seasons
	}
	return score
}

// calculateRecentConsumptionPenalty penalizes meals based on recent user consumption history.
func calculateRecentConsumptionPenalty(user models.User, meal models.Meal) float64 {
	for _, recentMeal := range user.RecentMeals {
		if recentMeal.ID == meal.ID {
			parsedTime, _ := ParseTime(recentMeal.Timestamp)
			return CalculateRecentMealPenalty(parsedTime)

		}
	}
	return 0.0
}

func CalculateRecentMealPenalty(mealDate time.Time) float64 {
	// Define constants for maximum and minimum penalty
	const maxPenalty = -15.0
	const minPenalty = -5.0

	currentDate := time.Now()
	// Calculate the number of days since the meal was consumed
	daysSinceMeal := currentDate.Sub(mealDate).Hours() / 24

	// Apply a linear scaling for the penalty
	// Scale the penalty to decrease from maxPenalty to minPenalty as daysSinceMeal increases
	penalty := maxPenalty
	if daysSinceMeal > 7 { // Assuming 7 days is the threshold for minimum penalty
		penalty = minPenalty
	} else {
		penalty = maxPenalty + (minPenalty-maxPenalty)*(daysSinceMeal/7)
	}

	return penalty
}

// calculateAgeGenderScore adjusts the meal score based on age and gender.
func calculateAgeGenderScore(user models.User, meal models.Meal, nutritionalContent map[string]float64) float64 {
	score := 0.0

	// Age factor: Older users get lower-calorie and nutrient-specific recommendations
	if user.Age > 50 {
		if nutritionalContent["calories"] <= 500.0 {
			score += 15.0
		}
		if containsNutrient(meal.Ingredients, "Calcium") || containsNutrient(meal.Ingredients, "Vitamin D") {
			score += 10.0
		}
	}

	// Gender factor: Protein for men, Iron for women
	if user.Gender == "Male" && nutritionalContent["protein"] >= 25.0 {
		score += 15.0
	} else if user.Gender == "Female" && containsNutrient(meal.Ingredients, "Iron") {
		score += 10.0
	}

	return score
}

// Utility function to check if a meal contains a nutrient
func containsNutrient(ingredients []models.Ingredient, nutrient string) bool {
	for _, ingredient := range ingredients {
		var nutritionalContent map[string]float64
		if err := json.Unmarshal(ingredient.Nutritional, &nutritionalContent); err != nil {
			// Handle error, possibly log it
			continue
		}
		if _, ok := nutritionalContent[nutrient]; ok {
			return true
		}
	}
	return false
}

// Utility functions to check if a slice contains a string or nutrient
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// GetTopMeals recommends the top N meals based on their calculated scores.
func GetTopMeals(user models.User, meals []models.Meal, topN int) []models.ScoredMeal {
	// Calculate scores for all meals
	mealsWithScores := CalculateMealScore(user, meals)

	// Return the top N meals
	if topN > len(mealsWithScores) {
		topN = len(mealsWithScores) // Handle the case where there are fewer meals than topN
	}
	return mealsWithScores[:topN]
}
