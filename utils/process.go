package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"nutrition/models"
	"sort"
	"sync"
	"time"
)

// CalculateMealScore computes scores for meals based on user preferences and goals.
func CalculateMealScore(user models.User, meals []models.Meal) []models.ScoredMeal {
	var wg sync.WaitGroup
	mealsWithScores := make([]models.ScoredMeal, len(meals))
	userBMI := calculateBMI(user.BodyMetrics.Height, user.BodyMetrics.Weight)
	const maxTotalScore = 155.0

	// Use a WaitGroup to score meals concurrently
	for i, meal := range meals {
		wg.Add(1)
		go func(i int, meal models.Meal) {
			defer wg.Done()

			fitScore := 0.0
			nutritionalContent := map[string]float64{}
			healthScores := map[string]float64{}

			// Unmarshal JSON concurrently
			if err := json.Unmarshal(meal.NutritionalContent, &nutritionalContent); err != nil {
				handleError(fmt.Errorf("error unmarshaling nutritional content for meal %d: %w", meal.ID, err))
				return
			}
			if err := json.Unmarshal(meal.HealthScores, &healthScores); err != nil {
				handleError(fmt.Errorf("error unmarshaling health scores for meal %d: %w", meal.ID, err))
				return
			}

			// Calculate the scores concurrently
			fitScore += calculateDietaryMatch(user, meal)
			fitScore += calculateNutritionalMatch(user, meal)
			fitScore += calculateHealthGoalsAlignment(user, nutritionalContent, healthScores, userBMI)
			fitScore += calculateNutritionalImpact(user, healthScores)
			fitScore += calculateMicrobiomeCompatibility(user, meal)
			fitScore += calculateEnvironmentalAdaptability(user, meal)
			fitScore += calculateRecentConsumptionPenalty(user, meal)
			fitScore += calculateAgeGenderScore(user, meal, nutritionalContent)

			// Normalize the score
			normalizedScore := normalizeScore(fitScore, maxTotalScore)
			mealsWithScores[i] = models.ScoredMeal{Meal: meal, Score: normalizedScore}
		}(i, meal)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Sort meals by normalized score in descending order
	sort.Slice(mealsWithScores, func(i, j int) bool {
		return mealsWithScores[i].Score > mealsWithScores[j].Score
	})

	return mealsWithScores
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

// calculateDietaryMatch scores meals based on user's dietary preferences and avoidances.
func calculateDietaryMatch(user models.User, meal models.Meal) float64 {
	score := 0.0
	preferences := map[string]bool{
		"vegetarian":  user.DietaryPreferences.Vegetarian,
		"vegan":       user.DietaryPreferences.Vegan,
		"gluten-free": user.DietaryPreferences.GlutenFree,
		"dairy-free":  user.DietaryPreferences.DairyFree,
	}

	for _, tag := range meal.Tags {
		if preferences[tag] {
			score += 10.0
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
	var wg sync.WaitGroup
	scoreMap := sync.Map{}

	// Cache frequently used values
	calories, hasCalories := nutritionalContent["calories"]
	sugar, hasSugar := nutritionalContent["sugar"]
	fiber, hasFiber := nutritionalContent["fiber"]

	// Function to calculate score for a specific goal
	calculateScore := func(goalType string) {
		defer wg.Done()
		score := 0.0

		switch goalType {
		case "Weight loss":
			if hasCalories && calories <= 400.0 {
				score += 30.0
			}
			if hasCalories {
				if bmi >= 25 && calories <= 400.0 {
					score += 30.0
				} else if bmi < 25 && calories <= 500.0 {
					score += 20.0
				}
			}
			if hasSugar && sugar <= 5.0 {
				score += 20.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 10.0
			}

		case "Muscle gain":
			if protein, hasProtein := nutritionalContent["protein"]; hasProtein && protein >= 25.0 {
				score += 30.0
			}
			if hasCalories && calories >= 500.0 {
				score += 20.0
			}
			if fat, hasFat := nutritionalContent["fat"]; hasFat && fat >= 10.0 {
				score += 10.0
			}
			if protein, hasProtein := nutritionalContent["protein"]; hasProtein && protein >= 25.0 && calories <= 600.0 {
				score += 30.0
			}

		case "Heart health":
			if heartHealthy, ok := healthScores["heart_healthy"]; ok && heartHealthy >= 8.0 {
				score += 30.0
			}
			if cholesterol, hasCholesterol := nutritionalContent["cholesterol"]; hasCholesterol && cholesterol <= 20.0 {
				score += 10.0
			}
			if fat, hasFat := nutritionalContent["fat"]; hasFat && fat <= 10.0 {
				score += 10.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 10.0
			}

		case "Blood sugar management":
			if diabetesFriendly, ok := healthScores["diabetes_friendly"]; ok && diabetesFriendly >= 8.0 {
				score += 30.0
			}
			if hasSugar && sugar <= 5.0 {
				score += 20.0
			}
			if hasFiber && fiber >= 8.0 {
				score += 10.0
			}

		case "Gut health":
			if hasFiber && fiber >= 8.0 {
				score += 30.0
			}

		case "Metabolic health":
			if hasSugar && sugar <= 5.0 && hasCalories && calories <= 500.0 {
				score += 30.0
			}
			if protein, hasProtein := nutritionalContent["protein"]; hasProtein && protein >= 15.0 {
				if fat, hasFat := nutritionalContent["fat"]; hasFat && fat <= 15.0 {
					score += 20.0
				}
			}

		case "Cholesterol reduction":
			if cholesterol, hasCholesterol := nutritionalContent["cholesterol"]; hasCholesterol && cholesterol <= 20.0 {
				score += 30.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 20.0
			}

		case "Weight maintenance":
			if hasCalories && calories >= 400.0 && calories <= 600.0 {
				score += 30.0
			}
			if fat, hasFat := nutritionalContent["fat"]; hasFat && fat <= 15.0 && hasSugar && sugar <= 5.0 {
				score += 20.0
			}
			if protein, hasProtein := nutritionalContent["protein"]; hasProtein && protein >= 15.0 {
				score += 10.0
			}

		case "Lean muscle maintenance":
			if protein, hasProtein := nutritionalContent["protein"]; hasProtein && protein >= 15.0 {
				score += 30.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 10.0
			}

		case "Improved energy levels":
			if carbs, hasCarbs := nutritionalContent["carbs"]; hasCarbs && carbs >= 40.0 {
				score += 30.0
			}
			if hasCalories && calories >= 400.0 && calories <= 600.0 {
				score += 20.0
			}
			if hasSugar && sugar <= 5.0 {
				score += 10.0
			}

		case "Endurance training support":
			if carbs, hasCarbs := nutritionalContent["carbs"]; hasCarbs && carbs >= 40.0 {
				score += 30.0
			}
			if protein, hasProtein := nutritionalContent["protein"]; hasProtein && protein >= 15.0 {
				score += 20.0
			}
			if fat, hasFat := nutritionalContent["fat"]; hasFat && fat <= 15.0 {
				score += 10.0
			}

		case "Cardiovascular fitness":
			if sodium, hasSodium := nutritionalContent["sodium"]; hasSodium && sodium <= 200.0 {
				score += 30.0
			}
			if cholesterol, hasCholesterol := nutritionalContent["cholesterol"]; hasCholesterol && cholesterol <= 20.0 {
				if fat, hasFat := nutritionalContent["fat"]; hasFat && fat <= 10.0 {
					score += 20.0
				}
			}
			if antioxidants, hasAntioxidants := nutritionalContent["antioxidants"]; hasAntioxidants && antioxidants >= 5.0 {
				score += 10.0
			}

		case "Detoxification":
			if hasFiber && fiber >= 8.0 {
				score += 30.0
			}

		case "Bone health":
			if calcium, hasCalcium := nutritionalContent["calcium"]; hasCalcium && calcium >= 200.0 {
				score += 30.0
			}

		case "Skin health":
			if vitaminA, hasVitaminA := nutritionalContent["vitamin A"]; hasVitaminA {
				if vitaminC, hasVitaminC := nutritionalContent["vitamin C"]; hasVitaminC && vitaminA >= 10.0 && vitaminC >= 20.0 {
					score += 30.0
				}
			}
			if healthyFats, hasHealthyFats := nutritionalContent["healthy fats"]; hasHealthyFats && healthyFats >= 10.0 {
				score += 20.0
			}

		case "Anti-inflammatory diet":
			if antioxidants, hasAntioxidants := nutritionalContent["antioxidants"]; hasAntioxidants && antioxidants >= 5.0 {
				score += 30.0
			}

		case "Hormonal balance":
			if omega3, hasOmega3 := nutritionalContent["omega-3"]; hasOmega3 && omega3 >= 10.0 {
				score += 30.0
			}
			if protein, hasProtein := nutritionalContent["protein"]; hasProtein && protein >= 15.0 && hasFiber && fiber >= 5.0 {
				score += 20.0
			}

		case "Improved mental clarity":
			if omega3, hasOmega3 := nutritionalContent["omega-3"]; hasOmega3 && omega3 >= 10.0 {
				score += 30.0
			}
			if antioxidants, hasAntioxidants := nutritionalContent["antioxidants"]; hasAntioxidants && antioxidants >= 5.0 {
				score += 20.0
			}

		default:
			handleError(fmt.Errorf("unknown health goal type: %s", goalType))

		}

		// Store the calculated score in the thread-safe map
		scoreMap.Store(goalType, score)
	}

	// Launch concurrent calculation for each goal
	for _, goal := range user.Goals {
		wg.Add(1)
		go calculateScore(goal.Type)
	}
	// Wait for all calculations to finish
	wg.Wait()

	// Sum up all scores from the thread-safe map
	totalScore := 0.0
	scoreMap.Range(func(_, value interface{}) bool {
		totalScore += value.(float64)
		return true
	})

	return totalScore
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
			
			return CalculateRecentMealPenalty(recentMeal.Timestamp)

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

// handleError logs and handles errors.
func handleError(err error) {
	// Here you can log the error to a file, monitoring system, or simply print it
	fmt.Printf("Error: %v\n", err)
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
