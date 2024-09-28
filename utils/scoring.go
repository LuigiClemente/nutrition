package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"nutrition/models"
	"sort"
	"strings"
	"sync"
	"time"
)

// calculateMealScore computes scores for meals based on user preferences and goals.
func calculateMealScore(user models.User, meals []models.Meal) []models.ScoredMeal {
	var wg sync.WaitGroup
	mealsWithScores := make([]models.ScoredMeal, len(meals))
	userBMI := calculateBMI(user.BodyMetrics.Height, user.BodyMetrics.Weight)
	const maxTotalScore = 120.0

	// Use a WaitGroup to score meals concurrently
	for i, meal := range meals {
		wg.Add(1)
		go func(i int, meal models.Meal) {
			defer wg.Done()

			fitScore := 0.0
			nutritionalContent := map[string]float64{}

			// Unmarshal JSON concurrently
			if err := json.Unmarshal(meal.NutritionalContent, &nutritionalContent); err != nil {
				handleError(fmt.Errorf("error unmarshaling nutritional content for meal %d: %w", meal.ID, err))
				return
			}

			// Calculate the scores concurrently
			fitScore += calculateDietaryMatch(user, meal)

			fitScore += calculateHealthGoalsAlignment(user, nutritionalContent, userBMI)

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

	// Create a map of dietary preferences
	preferences := map[string]bool{
		"Vegetarian":  user.DietaryPreferences.Vegetarian,
		"Vegan":       user.DietaryPreferences.Vegan,
		"Gluten-Free": user.DietaryPreferences.GlutenFree,
		"Dairy-Free":  user.DietaryPreferences.DairyFree,
	}

	// Iterate over the tags of the meal
	for _, tag := range meal.MealTags {
		// Check if the tag's name matches any of the user preferences
		if match, exists := preferences[tag.Tag]; exists && match {
			score += 10.0 // Increase score for each matching tag
		}
	}

	return score
}

// calculateHealthGoalsAlignment scores meals based on how well they align with user's health goals.
func calculateHealthGoalsAlignment(user models.User, nutritionalContent map[string]float64, bmi float64) float64 {
	var wg sync.WaitGroup
	scoreMap := sync.Map{}

	// Cache frequently used values
	calories, hasCalories := nutritionalContent["Calories"]
	sugar, hasSugar := nutritionalContent["Sugar"]
	fiber, hasFiber := nutritionalContent["Fiber"]

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
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 25.0 {
				score += 30.0
			}
			if hasCalories && calories >= 500.0 {
				score += 20.0
			}
			if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat >= 10.0 {
				score += 10.0
			}
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 25.0 && calories <= 600.0 {
				score += 30.0
			}

		case "Heart health":

			if cholesterol, hasCholesterol := nutritionalContent["Cholesterol"]; hasCholesterol && cholesterol <= 20.0 {
				score += 10.0
			}
			if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 10.0 {
				score += 10.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 10.0
			}

		case "Blood sugar management":

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
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 {
				if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 15.0 {
					score += 20.0
				}
			}

		case "Cholesterol reduction":
			if cholesterol, hasCholesterol := nutritionalContent["Cholesterol"]; hasCholesterol && cholesterol <= 20.0 {
				score += 30.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 20.0
			}

		case "Weight maintenance":
			if hasCalories && calories >= 400.0 && calories <= 600.0 {
				score += 30.0
			}
			if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 15.0 && hasSugar && sugar <= 5.0 {
				score += 20.0
			}
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 {
				score += 10.0
			}

		case "Lean muscle maintenance":
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 {
				score += 30.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 10.0
			}

		case "Improved energy levels":
			if carbs, hasCarbs := nutritionalContent["Carbs"]; hasCarbs && carbs >= 40.0 {
				score += 30.0
			}
			if hasCalories && calories >= 400.0 && calories <= 600.0 {
				score += 20.0
			}
			if hasSugar && sugar <= 5.0 {
				score += 10.0
			}

		case "Endurance training support":
			if carbs, hasCarbs := nutritionalContent["Carbs"]; hasCarbs && carbs >= 40.0 {
				score += 30.0
			}
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 {
				score += 20.0
			}
			if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 15.0 {
				score += 10.0
			}

		case "Cardiovascular fitness":
			if sodium, hasSodium := nutritionalContent["Sodium"]; hasSodium && sodium <= 200.0 {
				score += 30.0
			}
			if cholesterol, hasCholesterol := nutritionalContent["Cholesterol"]; hasCholesterol && cholesterol <= 20.0 {
				if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 10.0 {
					score += 20.0
				}
			}
			if antioxidants, hasAntioxidants := nutritionalContent["Antioxidants"]; hasAntioxidants && antioxidants >= 5.0 {
				score += 10.0
			}

		case "Detoxification":
			if hasFiber && fiber >= 8.0 {
				score += 30.0
			}

		case "Bone health":
			if calcium, hasCalcium := nutritionalContent["Calcium"]; hasCalcium && calcium >= 200.0 {
				score += 30.0
			}

		case "Skin health":
			if vitaminA, hasVitaminA := nutritionalContent["Vitamin A"]; hasVitaminA {
				if vitaminC, hasVitaminC := nutritionalContent["Vitamin C"]; hasVitaminC && vitaminA >= 10.0 && vitaminC >= 20.0 {
					score += 30.0
				}
			}
			if healthyFats, hasHealthyFats := nutritionalContent["healthy fats"]; hasHealthyFats && healthyFats >= 10.0 {
				score += 20.0
			}

		case "Anti-inflammatory diet":
			if antioxidants, hasAntioxidants := nutritionalContent["Antioxidants"]; hasAntioxidants && antioxidants >= 5.0 {
				score += 30.0
			}

		case "Hormonal balance":
			if omega3, hasOmega3 := nutritionalContent["Omega-3"]; hasOmega3 && omega3 >= 10.0 {
				score += 30.0
			}
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 && hasFiber && fiber >= 5.0 {
				score += 20.0
			}

		case "Improved mental clarity":
			if omega3, hasOmega3 := nutritionalContent["Omega-3"]; hasOmega3 && omega3 >= 10.0 {
				score += 30.0
			}
			if antioxidants, hasAntioxidants := nutritionalContent["Antioxidants"]; hasAntioxidants && antioxidants >= 5.0 {
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

// calculateMicrobiomeCompatibility scores meals based on how well they support the user's microbiome.
func calculateMicrobiomeCompatibility(user models.User, meal models.Meal) float64 {
	// Iterate over the gut health recommendations of the user
	for _, recommendation := range user.MicrobiomeData.GutHealthRecommendations {
		// Check if the meal contains the recommendation as a tag
		if containsRecommendationTag(meal.MealTags, recommendation) {
			return 10.0 // Increase score for gut-health-promoting meals
		}
	}
	return 0.0
}

// Helper function to check if a recommendation is in the meal tags
func containsRecommendationTag(tags []models.MealTag, recommendation string) bool {
	for _, tag := range tags {
		if tag.Tag == recommendation {
			return true
		}
	}
	return false
}

// calculateEnvironmentalAdaptability scores meals based on their environmental impact and user’s preferences.
func calculateEnvironmentalAdaptability(user models.User, meal models.Meal) float64 {
	score := 0.0

	// If the user's current season is Winter, increase score for warm meals
	if user.EnvironmentalFactors.Season == "Winter" && containsTag(meal.MealTags, "Warm") {
		score += 5.0
	}

	// Adapt the meal based on user location (example logic)
	if user.EnvironmentalFactors.Location == "Tropical" && containsTag(meal.MealTags, "Cooling") {
		score += 3.0
	}

	// Adapt to the user's climate (example: if the climate is hot, avoid heavy or warm meals)
	if user.EnvironmentalFactors.Climate == "Hot" && containsTag(meal.MealTags, "Light") {
		score += 4.0
	}

	return score
}

// Utility function to check if a meal has a specific tag (case-insensitive)
func containsTag(tags []models.MealTag, targetTag string) bool {
	for _, tag := range tags {
		if strings.Contains(tag.Tag, targetTag) {
			return true
		}
	}
	return false
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
		if nutritionalContent["Calories"] <= 500.0 {
			score += 15.0
		}

	}

	// Gender factor: Protein for men, Iron for women
	if user.Gender == "Male" && nutritionalContent["Protein"] >= 25.0 {
		score += 15.0
	} else if user.Gender == "Female" {
		score += 10.0
	}

	return score
}
