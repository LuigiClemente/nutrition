package utils

import (
	"encoding/json"
	"fmt"
	"nutrition/models"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	/* 
	MaxAgeGenderScore           = 20.0
	MaxRecentConsumptionPenalty = 10.0
	MaxEnvironmentalAdaptability = 15.0
	MaxMicrobiomeCompatibility  = 10.0
	MaxDietaryMatch             = 25.0
	MaxHealthGoalsAlignment     = 30.0
	*/
	maxTotalScore = 110.0
)

// calculateMealScore computes scores for meals based on user preferences and goals.
func calculateMealScore(user models.User, meals []models.Meal) []models.ScoredMeal {
	var wg sync.WaitGroup
	mealsWithScores := make([]models.ScoredMeal, len(meals))
	userBMI := calculateBMI(user.BodyMetrics.Height, user.BodyMetrics.Weight)
	

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
			fitScore += calculateAgeGenderScore(user, nutritionalContent)

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

	// Create a map for quick tag lookup
	tagMap := make(map[string]bool)
	for _, tag := range meal.MealTags {
		tagMap[strings.ToLower(tag.Tag)] = true
	}

	// Score adjustments based on user season
	if strings.ToLower(user.EnvironmentalFactors.Season) == "winter" && tagMap["warm"] {
		score += 5.0 // Favor warm meals in winter
	}

	// Score adjustments based on user location
	switch strings.ToLower(user.EnvironmentalFactors.Location) {
	case "tropical":
		if tagMap["cooling"] {
			score += 3.0 // Favor cooling meals in tropical locations
		}
		if tagMap["light"] {
			score += 2.0 // Favor light meals in tropical locations
		}
	case "urban":
		if tagMap["quick"] {
			score += 3.0 // Favor quick meals for urban lifestyles
		}
	case "rural":
		if tagMap["hearty"] {
			score += 4.0 // Favor hearty meals for rural settings
		}
	}

	// Score adjustments based on user climate
	if strings.ToLower(user.EnvironmentalFactors.Climate) == "hot" && tagMap["light"] {
		score += 4.0 // Favor light meals in hot climates
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
func calculateAgeGenderScore(user models.User, nutritionalContent map[string]float64) float64 {
	score := 0.0

	// Define thresholds for nutrients based on user age
	var calorieThreshold float64
	var proteinThreshold float64

	// Set thresholds based on age
	switch {
	case user.Age < 18:
		calorieThreshold = 600.0
		proteinThreshold = 0 // Not emphasized for teens in this logic
	case user.Age <= 50:
		proteinThreshold = 20.0
		// No specific calorie threshold for this age range
	default: // user.Age > 50
		calorieThreshold = 500.0
	}

	// Age-based scoring
	if user.Age < 18 && nutritionalContent["Calories"] >= calorieThreshold {
		score += 10.0
	} else if user.Age > 50 && nutritionalContent["Calories"] <= calorieThreshold {
		score += 15.0
	} else if user.Age >= 18 && user.Age <= 50 && nutritionalContent["Protein"] >= proteinThreshold {
		score += 10.0
	}

	// Gender-based scoring
	switch user.Gender {
	case "Male":
		if nutritionalContent["Protein"] >= 25.0 {
			score += 15.0
		}
		if nutritionalContent["Zinc"] >= 11.0 {
			score += 5.0
		}
	case "Female":
		if nutritionalContent["Iron"] >= 18.0 {
			score += 15.0
		}
		if nutritionalContent["Calcium"] >= 1000.0 {
			score += 5.0
		}
	}

	return score
}
