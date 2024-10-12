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
		MaxEnvironmentalAdaptability = 22.0
		MaxMicrobiomeCompatibility  = 10.0
		MaxDietaryMatch             = 25.0
		MaxHealthGoalsAlignment     = 30.0
	*/
	maxTotalScore = 117.0
)

// calculateMealScore computes scores for meals based on user preferences and goals.
func CalculateMealScore(user models.User, meals []models.Meal) []models.ScoredMeal {
	var wg sync.WaitGroup
	mealsWithScores := make([]models.ScoredMeal, len(meals))
	userBMI := calculateBMI(user.BodyMetrics.Height, user.BodyMetrics.Weight)

	for i, meal := range meals {
		nutritionalContent := map[string]float64{}
		// Unmarshal JSON before launching goroutine
		if err := json.Unmarshal(meal.NutritionalContent, &nutritionalContent); err != nil {
			handleError(fmt.Errorf("error unmarshaling nutritional content for meal %d: %w", meal.ID, err))
			continue
		}

		wg.Add(1)
		go func(i int, meal models.Meal) {
			defer wg.Done()

			fitScore := 0.0
			fitScore += calculateDietaryMatch(user, meal)
			fitScore += calculateHealthGoalsAlignment(user, nutritionalContent, userBMI)
			fitScore += calculateMicrobiomeCompatibility(user, meal)
			fitScore += calculateEnvironmentalAdaptability(user, meal)
			fitScore += calculateRecentConsumptionPenalty(user, meal)
			fitScore += calculateAgeGenderScore(user, nutritionalContent)

			// Normalize the score
			normalizedScore := normalizeScore(fitScore, maxTotalScore)
			mealsWithScores[i] = models.ScoredMeal{Meal: &meal, Score: normalizedScore}
		}(i, meal)
	}

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

	// Create a map of dietary preferences in lowercase
	preferences := map[string]bool{
		"vegetarian":  user.DietaryPreferences.Vegetarian,
		"vegan":       user.DietaryPreferences.Vegan,
		"gluten-free": user.DietaryPreferences.GlutenFree,
		"dairy-free":  user.DietaryPreferences.DairyFree,
	}

	// Iterate over the tags of the meal
	for _, tag := range meal.MealTags {
		// Convert the tag to lowercase and check if it matches any user preference
		tagLower := strings.ToLower(tag.Tag)
		if match, exists := preferences[tagLower]; exists && match {
			score = 10.0
			break
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

		switch strings.ToLower(goalType) {
		case "weight loss":
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

		case "muscle gain":
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

		case "heart health":

			if cholesterol, hasCholesterol := nutritionalContent["Cholesterol"]; hasCholesterol && cholesterol <= 20.0 {
				score += 10.0
			}
			if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 10.0 {
				score += 10.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 10.0
			}

		case "blood sugar management":

			if hasSugar && sugar <= 5.0 {
				score += 20.0
			}
			if hasFiber && fiber >= 8.0 {
				score += 10.0
			}

		case "gut health":
			if hasFiber && fiber >= 8.0 {
				score += 30.0
			}

		case "metabolic health":
			if hasSugar && sugar <= 5.0 && hasCalories && calories <= 500.0 {
				score += 30.0
			}
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 {
				if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 15.0 {
					score += 20.0
				}
			}

		case "cholesterol reduction":
			if cholesterol, hasCholesterol := nutritionalContent["Cholesterol"]; hasCholesterol && cholesterol <= 20.0 {
				score += 30.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 20.0
			}

		case "weight maintenance":
			if hasCalories && calories >= 400.0 && calories <= 600.0 {
				score += 30.0
			}
			if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 15.0 && hasSugar && sugar <= 5.0 {
				score += 20.0
			}
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 {
				score += 10.0
			}

		case "lean muscle maintenance":
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 {
				score += 30.0
			}
			if hasFiber && fiber >= 5.0 {
				score += 10.0
			}

		case "improved energy levels":
			if carbs, hasCarbs := nutritionalContent["Carbs"]; hasCarbs && carbs >= 40.0 {
				score += 30.0
			}
			if hasCalories && calories >= 400.0 && calories <= 600.0 {
				score += 20.0
			}
			if hasSugar && sugar <= 5.0 {
				score += 10.0
			}

		case "endurance training support":
			if carbs, hasCarbs := nutritionalContent["Carbs"]; hasCarbs && carbs >= 40.0 {
				score += 30.0
			}
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 {
				score += 20.0
			}
			if fat, hasFat := nutritionalContent["Fat"]; hasFat && fat <= 15.0 {
				score += 10.0
			}

		case "cardiovascular fitness":
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

		case "detoxification":
			if hasFiber && fiber >= 8.0 {
				score += 30.0
			}

		case "bone health":
			if calcium, hasCalcium := nutritionalContent["Calcium"]; hasCalcium && calcium >= 200.0 {
				score += 30.0
			}

		case "skin health":
			if vitaminA, hasVitaminA := nutritionalContent["Vitamin A"]; hasVitaminA {
				if vitaminC, hasVitaminC := nutritionalContent["Vitamin C"]; hasVitaminC && vitaminA >= 10.0 && vitaminC >= 20.0 {
					score += 30.0
				}
			}
			if healthyFats, hasHealthyFats := nutritionalContent["healthy fats"]; hasHealthyFats && healthyFats >= 10.0 {
				score += 20.0
			}

		case "anti-inflammatory diet":
			if antioxidants, hasAntioxidants := nutritionalContent["Antioxidants"]; hasAntioxidants && antioxidants >= 5.0 {
				score += 30.0
			}

		case "hormonal balance":
			if omega3, hasOmega3 := nutritionalContent["Omega-3"]; hasOmega3 && omega3 >= 10.0 {
				score += 30.0
			}
			if protein, hasProtein := nutritionalContent["Protein"]; hasProtein && protein >= 15.0 && hasFiber && fiber >= 5.0 {
				score += 20.0
			}

		case "improved mental clarity":
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
	// Define mappings of gut health recommendations to meal tags
	recommendationToTag := map[string]string{
		"Increase Fiber":             "High Fiber",
		"Reduce Fat Intake":          "Low Fat",
		"Add Probiotics":             "Probiotic",
		"Eat More Vegetables":        "Plant-Based",
		"Increase Water Intake":      "Hydrating",
		"Reduce Sugar":               "Sugar-Free",
		"Add Prebiotics":             "Prebiotic",
		"Reduce Meat Consumption":    "Vegetarian",
		"Avoid Gluten":               "Gluten-Free",
		"Increase Protein":           "High Protein",
		"Eat Fermented Foods":        "Fermented",
		"Increase Calcium":           "High Calcium",
		"Eat Leafy Greens":           "Plant-Based",
		"Reduce Salt":                "Low Sodium",
		"Increase Fruits":            "Plant-Based",
		"Add Whole Grains":           "Whole Grain",
		"Eat More Yogurt":            "Probiotic",
		"Avoid Allergens":            "Allergen-Free",
		"Increase Magnesium":         "High Magnesium",
		"Eat Omega-3 Rich Foods":     "Heart-Healthy",
		"Reduce Cholesterol":         "Low Cholesterol",
		"Increase Iron Intake":       "Iron-Rich",
		"Reduce Processed Foods":     "Whole30",
		"Add Antioxidants":           "Anti-Inflammatory",
		"Increase Healthy Fats":      "Keto",
		"Eat Plant-Based Proteins":   "Plant-Based",
		"Increase Omega-3":           "Heart-Healthy",
		"Avoid Processed Sugar":      "Sugar-Free",
		"Eat More Fiber":             "High Fiber",
		"Reduce Dairy Intake":        "Dairy-Free",
		"Eat More Nuts and Seeds":    "Nut-Free",
		"Avoid Refined Carbs":        "Low Carb",
		"Eat More Legumes":           "Plant-Based",
		"Reduce Alcohol Consumption": "Moderate",
		"Eat More Whole Foods":       "Whole30",
	}

	// Create a set of meal tags for faster lookup
	mealTagSet := make(map[string]struct{})
	for _, tag := range meal.MealTags {
		mealTagSet[strings.ToLower(tag.Tag)] = struct{}{}
	}

	score := 0.0

	// Iterate over the gut health recommendations of the user
	for _, recommendation := range user.MicrobiomeData.GutHealthRecommendations {
		// Trim spaces and convert recommendation to lowercase for comparison
		trimmedRecommendation := strings.TrimSpace(strings.ToLower(recommendation))

		// Check if the recommendation maps to a tag
		for key, tag := range recommendationToTag {
			if strings.ToLower(strings.TrimSpace(key)) == trimmedRecommendation {
				// Convert the tag to lowercase and check against the meal tag set
				lowerTag := strings.ToLower(tag)
				if _, exists := mealTagSet[lowerTag]; exists {
					score = 10.0

					break // Exit loop after finding the match
				}
			}
		}

	}

	return score
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
	if strings.ToLower(user.EnvironmentalFactors.Season) == "winter" && tagMap["healthy"] {
		score += 5.0 // Favor healthy meals in winter
	}
	if strings.ToLower(user.EnvironmentalFactors.Season) == "summer" && tagMap["light"] {
		score += 5.0 // Favor light meals in summer
	}

	// Score adjustments based on user climate
	switch strings.ToLower(user.EnvironmentalFactors.Climate) {
	case "tropical":
		if tagMap["cooling"] {
			score += 3.0 // Favor cooling meals in tropical climates
		}
		if tagMap["light"] {
			score += 2.0 // Favor light meals in tropical climates
		}
	case "urban":
		if tagMap["quick"] {
			score += 3.0 // Favor quick meals for urban lifestyles
		}
		if tagMap["healthy"] {
			score += 2.0 // Favor healthy meals in urban settings
		}
	case "rural":
		if tagMap["hearty"] {
			score += 4.0 // Favor hearty meals for rural settings
		}
		if tagMap["family-friendly"] {
			score += 3.0 // Favor family-friendly meals in rural areas
		}
	case "temperate":
		if tagMap["organic"] {
			score += 2.0 // Favor organic meals in temperate climates
		}
	case "hot":
		if tagMap["light"] {
			score += 4.0 // Favor light meals in hot climates
		}
		if tagMap["low calorie"] {
			score += 3.0 // Favor low-calorie meals in hot climates
		}
	}

	// General tags that can enhance adaptability
	if tagMap["plant-based"] {
		score += 3.0 // Favor plant-based meals
	}
	if tagMap["heart-healthy"] {
		score += 2.0 // Favor heart-healthy meals
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

	case "Other":
		// Implement scoring criteria for "Other" gender
		if nutritionalContent["Iron"] >= 15.0 { // Example criteria
			score += 10.0
		}
		if nutritionalContent["Calcium"] >= 800.0 { // Example criteria
			score += 5.0
		}
		if nutritionalContent["Vitamin D"] >= 20.0 { // Example criteria for general health
			score += 10.0
		}
	default:
		handleError(fmt.Errorf("unknown gender type: %s", user.Gender))
	}

	return score
}
