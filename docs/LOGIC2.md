[Here’s a Go function for scoring and recommending meals based on user data:

```go
package main

import (
	"fmt"
	"sort"
)

// User, Meal, and related structures
type User struct {
	ID                      int
	Name                    string
	Age                     int
	Gender                  string
	BodyMetrics             BodyMetrics
	ActivityLevel           string
	DietaryPreferences      DietaryPreferences
	HealthConditions        []HealthCondition
	NutritionalDeficiencies []string
	Allergies               []string
	RecentMeals             []Meal
	Goals                   []Goal
	MicrobiomeData          MicrobiomeData
	BloodGlucose            float64
	LipidProfile            LipidProfile
	EnvironmentalFactors    EnvironmentalFactors
	MealHistory             []MealHistory
	HealthScore             float64
	Preferences             UserPreferences
}

type BodyMetrics struct {
	Weight float64
	Height float64
	BMI    float64
}

type DietaryPreferences struct {
	Vegetarian         bool
	Vegan              bool
	GlutenFree         bool
	DairyFree          bool
	SpecificAvoidances []string
}

type HealthCondition struct {
	Name     string
	Severity string
}

type Goal struct {
	Type     string
	Target   float64
	Duration int
}

type MicrobiomeData struct {
	DiversityScore           float64
	GutHealthRecommendations []string
}

type LipidProfile struct {
	Cholesterol   float64
	HDL           float64
	LDL           float64
	Triglycerides float64
}

type EnvironmentalFactors struct {
	Location string
	Climate  string
	Season   string
}

type MealHistory struct {
	Meal      Meal
	Timestamp string
}

type UserPreferences struct {
	PreferredCuisines   []string
	MealTimings         []string
	FavoriteIngredients []string
}

// Meal and related structures
type Meal struct {
	ID                 int
	Name               string
	Ingredients        []Ingredient
	NutritionalContent map[string]float64
	Category           string
	MealType           []string
	Cuisine            string
	Tags               []string
	HealthScores       map[string]float64
	PreparationTime    int
	Difficulty         string
	ServingSize        int
	Instructions       string
}

type Ingredient struct {
	Name        string
	Amount      float64
	Unit        string
	Nutritional map[string]float64
}

// ScoredMeal to store meal and its scores
type ScoredMeal struct {
	Meal  Meal
	Score float64
}

func CalculateMealScore(user User, meals []Meal) []ScoredMeal {
	mealsWithScores := []ScoredMeal{}

	for _, meal := range meals {
		fitScore := 0.0

		// Check dietary preferences and avoidances
		for _, tag := range meal.Tags {
			if user.DietaryPreferences.Vegetarian && tag == "vegetarian" {
				fitScore += 10.0
			}
			if user.DietaryPreferences.Vegan && tag == "vegan" {
				fitScore += 10.0
			}
			if user.DietaryPreferences.GlutenFree && tag == "gluten-free" {
				fitScore += 10.0
			}
			if user.DietaryPreferences.DairyFree && tag == "dairy-free" {
				fitScore += 10.0
			}
		}

		// Match Nutritional Deficiencies
		for _, deficiency := range user.NutritionalDeficiencies {
			if deficiency == "Iron" && meal.NutritionalContent["iron"] > 0 {
				fitScore += 15.0
			}
			if deficiency == "Vitamin D" && meal.NutritionalContent["vitamin_d"] > 0 {
				fitScore += 15.0
			}
		}

		// Match Health Conditions
		for _, condition := range user.HealthConditions {
			if condition.Name == "Diabetes" && meal.HealthScores["diabetes_friendly"] >= 8.0 {
				fitScore += 20.0
			}
			if condition.Name == "Hypertension" && meal.HealthScores["heart_healthy"] >= 8.0 {
				fitScore += 20.0
			}
		}

		// Match Health Goals
		for _, goal := range user.Goals {
			switch goal.Type {
			case "Weight loss":
				if meal.NutritionalContent["calories"] <= 400.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["sugar"] <= 5.0 {
					fitScore += 20.0
				}
				if meal.NutritionalContent["fiber"] >= 5.0 {
					fitScore += 10.0
				}

			case "Muscle gain":
				if meal.NutritionalContent["protein"] >= 25.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["calories"] >= 500.0 {
					fitScore += 20.0
				}
				if meal.NutritionalContent["fat"] >= 10.0 {
					fitScore += 10.0
				}

			case "Heart health":
				if meal.HealthScores["heart_healthy"] >= 8.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["cholesterol"] <= 20.0 {
					fitScore += 10.0
				}
				if meal.NutritionalContent["fat"] <= 10.0 {
					fitScore += 10.0
				}
				if meal.NutritionalContent["fiber"] >= 5.0 {
					fitScore += 10.0
				}

			case "Blood sugar management":
				if meal.HealthScores["diabetes_friendly"] >= 8.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["sugar"] <= 5.0 {
					fitScore += 20.0
				}
				if meal.NutritionalContent["fiber"] >= 8.0 {
					fitScore += 10.0
				}

			case "Gut health":
				if meal.NutritionalContent["fiber"] >= 8.0 {
					fitScore += 30.0
				}
				// Check for probiotic or prebiotic ingredients, if available in your meal structure

			case "Metabolic health":
				if meal.NutritionalContent["sugar"] <= 5.0 && meal.NutritionalContent["calories"] <= 500.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["protein"] >= 15.0 && meal.NutritionalContent["fat"] <= 15.0 {
					fitScore += 20.0
				}

			case "Cholesterol reduction":
				if meal.NutritionalContent["cholesterol"] <= 20.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["fiber"] >= 5.0 {
					fitScore += 20.0
				}
				// Include other ingredients that reduce LDL cholesterol, like oats

			case "Weight maintenance":
				if meal.NutritionalContent["calories"] >= 400.0 && meal.NutritionalContent["calories"] <= 600.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["fat"] <= 15.0 && meal.NutritionalContent["sugar"] <= 5.0 {
					fitScore += 20.0
				}
				if meal.NutritionalContent["protein"] >= 15.0 {
					fitScore += 10.0
				}

			case "Lean muscle maintenance":
				if meal.NutritionalContent["protein"] >= 15.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["fiber"] >= 5.0 {
					fitScore += 10.0
				}

			case "Improved energy levels":
				if meal.NutritionalContent["carbs"] >= 40.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["calories"] >= 400.0 && meal.NutritionalContent["calories"] <= 600.0 {
					fitScore += 20.0
				}
				if meal.NutritionalContent["sugar"] <= 5.0 {
					fitScore += 10.0
				}

			case "Endurance training support":
				if meal.NutritionalContent["carbs"] >= 40.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["protein"] >= 15.0 {
					fitScore += 20.0
				}
				if meal.NutritionalContent["fat"] <= 15.0 {
					fitScore += 10.0
				}

			case "Cardiovascular fitness":
				if meal.NutritionalContent["sodium"] <= 200.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["cholesterol"] <= 20.0 && meal.NutritionalContent["fat"] <= 10.0 {
					fitScore += 20.0
				}
				if meal.NutritionalContent["antioxidants"] >= 5.0 { // Assume some score for antioxidants
					fitScore += 10.0
				}

			case "Detoxification":
				if meal.NutritionalContent["fiber"] >= 8.0 {
					fitScore += 30.0
				}
				// Include ingredients like lemon, ginger, greens that promote detox

			case "Bone health":
				if meal.NutritionalContent["calcium"] >= 200.0 {
					fitScore += 30.0
				}
				// Include other nutrients like vitamin D and magnesium

			case "Skin health":
				if meal.NutritionalContent["vitamin A"] >= 10.0 && meal.NutritionalContent["vitamin C"] >= 20.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["healthy fats"] >= 10.0 {
					fitScore += 20.0
				}

			case "Anti-inflammatory diet":
				if meal.NutritionalContent["antioxidants"] >= 5.0 {
					fitScore += 30.0
				}
				// Include anti-inflammatory ingredients like turmeric, ginger

			case "Hormonal balance":
				if meal.NutritionalContent["omega-3"] >= 10.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["protein"] >= 15.0 && meal.NutritionalContent["fiber"] >= 5.0 {
					fitScore += 20.0
				}

			case "Improved mental clarity":
				if meal.NutritionalContent["omega-3"] >= 10.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["antioxidants"] >= 5.0 {
					fitScore += 20.0
				}

			case "Hydration and fluid balance":
				if meal.NutritionalContent["water content"] >= 50.0 { // Assuming some % of water content
					fitScore += 30.0
				}
				if meal.NutritionalContent["sodium"] <= 150.0 {
					fitScore += 20.0
				}
				if meal.NutritionalContent["potassium"] >= 300.0 {
					fitScore += 10.0
				}

			case "Immune support":
				if meal.NutritionalContent["vitamin C"] >= 30.0 && meal.NutritionalContent["vitamin D"] >= 15.0 {
					fitScore += 30.0
				}
				if meal.NutritionalContent["zinc"] >= 10.0 {
					fitScore += 20.0
				}
				// Include immune-boosting ingredients like garlic, ginger, citrus
			}
		}

		// Consider environmental factors
		if user.EnvironmentalFactors.Season == "Winter" && contains(meal.Tags, "warm") {
			fitScore += 5.0
		}
		if user.EnvironmentalFactors.Climate == "Cold" && contains(meal.Tags, "warm") {
			fitScore += 5.0
		}

		// Append meal with score to the result
		mealsWithScores = append(mealsWithScores, ScoredMeal{Meal: meal, Score: fitScore})
	}

	// Sort meals by score in descending order
	sort.Slice(mealsWithScores, func(i, j int) bool {
		return mealsWithScores[i].Score > mealsWithScores[j].Score
	})

	return mealsWithScores
}

// calculateNutritionalMatching scores meals based on how well they match the user's nutritional deficiencies.
func calculateNutritionalMatching(user User, meal Meal) float64 {
	score := 0.0
	for _, deficiency := range user.NutritionalDeficiencies {
		if containsNutrient(meal.Ingredients, deficiency) {
			score += 10.0 // Increase score for every deficiency matched
		}
	}
	return score
}

// Helper function to extract ingredient names from a meal
func extractIngredientNames(ingredients []Ingredient) []string {
	var names []string
	for _, ingredient := range ingredients {
		names = append(names, ingredient.Name)
	}
	return names
}

// calculateHealthGoalsAlignment scores meals based on how well they align with the user's health goals.
func calculateHealthGoalsAlignment(user User, meal Meal) float64 {
	score := 0.0
	// Match Health Goals
	for _, goal := range user.Goals {
		switch goal.Type {
		case "Weight loss":
			if meal.NutritionalContent["calories"] <= 400.0 {
				score += 30.0
			}
			if meal.NutritionalContent["sugar"] <= 5.0 {
				score += 20.0
			}
			if meal.NutritionalContent["fiber"] >= 5.0 {
				score += 10.0
			}

		case "Muscle gain":
			if meal.NutritionalContent["protein"] >= 25.0 {
				score += 30.0
			}
			if meal.NutritionalContent["calories"] >= 500.0 {
				score += 20.0
			}
			if meal.NutritionalContent["fat"] >= 10.0 {
				score += 10.0
			}

		case "Heart health":
			if meal.HealthScores["heart_healthy"] >= 8.0 {
				score += 30.0
			}
			if meal.NutritionalContent["cholesterol"] <= 20.0 {
				score += 10.0
			}
			if meal.NutritionalContent["fat"] <= 10.0 {
				score += 10.0
			}
			if meal.NutritionalContent["fiber"] >= 5.0 {
				score += 10.0
			}

		case "Blood sugar management":
			if meal.HealthScores["diabetes_friendly"] >= 8.0 {
				score += 30.0
			}
			if meal.NutritionalContent["sugar"] <= 5.0 {
				score += 20.0
			}
			if meal.NutritionalContent["fiber"] >= 8.0 {
				score += 10.0
			}

		case "Gut health":
			if meal.NutritionalContent["fiber"] >= 8.0 {
				score += 30.0
			}
			// Check for probiotic or prebiotic ingredients, if available in your meal structure

		case "Metabolic health":
			if meal.NutritionalContent["sugar"] <= 5.0 && meal.NutritionalContent["calories"] <= 500.0 {
				score += 30.0
			}
			if meal.NutritionalContent["protein"] >= 15.0 && meal.NutritionalContent["fat"] <= 15.0 {
				score += 20.0
			}

		case "Cholesterol reduction":
			if meal.NutritionalContent["cholesterol"] <= 20.0 {
				score += 30.0
			}
			if meal.NutritionalContent["fiber"] >= 5.0 {
				score += 20.0
			}
			// Include other ingredients that reduce LDL cholesterol, like oats

		case "Weight maintenance":
			if meal.NutritionalContent["calories"] >= 400.0 && meal.NutritionalContent["calories"] <= 600.0 {
				score += 30.0
			}
			if meal.NutritionalContent["fat"] <= 15.0 && meal.NutritionalContent["sugar"] <= 5.0 {
				score += 20.0
			}
			if meal.NutritionalContent["protein"] >= 15.0 {
				score += 10.0
			}

		case "Lean muscle maintenance":
			if meal.NutritionalContent["protein"] >= 15.0 {
				score += 30.0
			}
			if meal.NutritionalContent["fiber"] >= 5.0 {
				score += 10.0
			}

		case "Improved energy levels":
			if meal.NutritionalContent["carbs"] >= 40.0 {
				score += 30.0
			}
			if meal.NutritionalContent["calories"] >= 400.0 && meal.NutritionalContent["calories"] <= 600.0 {
				score += 20.0
			}
			if meal.NutritionalContent["sugar"] <= 5.0 {
				score += 10.0
			}

		case "Endurance training support":
			if meal.NutritionalContent["carbs"] >= 40.0 {
				score += 30.0
			}
			if meal.NutritionalContent["protein"] >= 15.0 {
				score += 20.0
			}
			if meal.NutritionalContent["fat"] <= 15.0 {
				score += 10.0
			}

		case "Cardiovascular fitness":
			if meal.NutritionalContent["sodium"] <= 200.0 {
				score += 30.0
			}
			if meal.NutritionalContent["cholesterol"] <= 20.0 && meal.NutritionalContent["fat"] <= 10.0 {
				score += 20.0
			}
			if meal.NutritionalContent["antioxidants"] >= 5.0 { // Assume some score for antioxidants
				score += 10.0
			}

		case "Detoxification":
			if meal.NutritionalContent["fiber"] >= 8.0 {
				score += 30.0
			}
			// Include ingredients like lemon, ginger, greens that promote detox

		case "Bone health":
			if meal.NutritionalContent["calcium"] >= 200.0 {
				score += 30.0
			}
			// Include other nutrients like vitamin D and magnesium

		case "Skin health":
			if meal.NutritionalContent["vitamin A"] >= 10.0 && meal.NutritionalContent["vitamin C"] >= 20.0 {
				score += 30.0
			}
			if meal.NutritionalContent["healthy fats"] >= 10.0 {
				score += 20.0
			}

		case "Anti-inflammatory diet":
			if meal.NutritionalContent["antioxidants"] >= 5.0 {
				score += 30.0
			}
			// Include anti-inflammatory ingredients like turmeric, ginger

		case "Hormonal balance":
			if meal.NutritionalContent["omega-3"] >= 10.0 {
				score += 30.0
			}
			if meal.NutritionalContent["protein"] >= 15.0 && meal.NutritionalContent["fiber"] >= 5.0 {
				score += 20.0
			}

		case "Improved mental clarity":
			if meal.NutritionalContent["omega-3"] >= 10.0 {
				score += 30.0
			}
			if meal.NutritionalContent["antioxidants"] >= 5.0 {
				score += 20.0
			}

		case "Hydration and fluid balance":
			if meal.NutritionalContent["water content"] >= 50.0 { // Assuming some % of water content
				score += 30.0
			}
			if meal.NutritionalContent["sodium"] <= 150.0 {
				score += 20.0
			}
			if meal.NutritionalContent["potassium"] >= 300.0 {
				score += 10.0
			}

		case "Immune support":
			if meal.NutritionalContent["vitamin C"] >= 30.0 && meal.NutritionalContent["vitamin D"] >= 15.0 {
				score += 30.0
			}
			if meal.NutritionalContent["zinc"] >= 10.0 {
				score += 20.0
			}
			// Include immune-boosting ingredients like garlic, ginger, citrus
		}
	}
	return score
}

// calculateRecentConsumptionPenalty penalizes meals that have been recently consumed to promote variety.
func calculateRecentConsumptionPenalty(user User, meal Meal) float64 {
	for _, recentMeal := range user.RecentMeals {
		if recentMeal.Name == meal.Name {
			return -15.0 // Penalize for recent consumption
		}
	}
	return 0.0
}

// calculateNutritionalImpact adjusts the score based on how meals impact the user's health monitoring (e.g., blood sugar, cholesterol).
func calculateNutritionalImpact(user User, meal Meal) float64 {
	score := 0.0
	if user.BloodGlucose > 100 && meal.HealthScores["diabetes_friendly"] >= 8.0 {
		score += 15.0 // Prioritize meals that help with blood sugar management
	}
	if user.LipidProfile.LDL > 100 && meal.HealthScores["heart_healthy"] >= 8.0 {
		score += 15.0 // Prioritize heart-healthy meals for users with high cholesterol
	}
	return score
}

// calculateMicrobiomeCompatibility scores meals based on how well they support the user's gut microbiome.
func calculateMicrobiomeCompatibility(user User, meal Meal) float64 {
	for _, recommendation := range user.MicrobiomeData.GutHealthRecommendations {
		if contains(meal.Tags, recommendation) {
			return 10.0 // Increase score for gut-health-promoting meals
		}
	}
	return 0.0
}

// calculateEnvironmentalAdaptability adjusts the score based on climate and season.
func calculateEnvironmentalAdaptability(user User, meal Meal) float64 {
	score := 0.0
	if user.EnvironmentalFactors.Season == "Winter" && contains(meal.Tags, "warm") {
		score += 5.0 // Increase score for warm foods in cold seasons
	}
	return score
}

// Utility functions to check if a slice contains a string or nutrient
func contains(slice []string, item string) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

func containsNutrient(ingredients []Ingredient, nutrient string) bool {
	for _, ingredient := range ingredients {
		if _, ok := ingredient.Nutritional[nutrient]; ok {
			return true
		}
	}
	return false
}

// calculateAgeGenderScore adjusts the meal score based on age and gender.
func calculateAgeGenderScore(user User, meal Meal) float64 {
	score := 0.0

	// Age factor: Older users get lower-calorie and nutrient-specific recommendations
	if user.Age > 50 {
		if meal.NutritionalContent["calories"] <= 500.0 {
			score += 15.0
		}
		if containsNutrient(meal.Ingredients, "Calcium") || containsNutrient(meal.Ingredients, "Vitamin D") {
			score += 10.0
		}
	}

	// Gender factor: Protein for men, Iron for women
	if user.Gender == "Male" && meal.NutritionalContent["protein"] >= 25.0 {
		score += 15.0
	} else if user.Gender == "Female" && containsNutrient(meal.Ingredients, "Iron") {
		score += 10.0
	}

	return score
}

// GetTopMeals recommends the top N meals based on their calculated scores.
func GetTopMeals(user User, meals []Meal, topN int) []ScoredMeal {
	// Calculate scores for all meals
	mealsWithScores := CalculateMealScore(user, meals)

	// Return the top N meals
	if topN > len(mealsWithScores) {
		topN = len(mealsWithScores) // Handle the case where there are fewer meals than topN
	}
	return mealsWithScores[:topN]
}

// Example usage
func start() {
	user := User{
		Name:          "Zara",
		Age:           20,
		Gender:        "Female",
		BodyMetrics:   BodyMetrics{Weight: 70.5, Height: 160.0, BMI: 12.8},
		ActivityLevel: "Moderately Active",
		DietaryPreferences: DietaryPreferences{
			Vegetarian:         false,
			Vegan:              false,
			GlutenFree:         true,
			DairyFree:          true,
			SpecificAvoidances: []string{"Red meat", "Processed sugar"},
		},
		HealthConditions: []HealthCondition{
			{Name: "Diabetes", Severity: "Moderate"},
		},
		NutritionalDeficiencies: []string{"Iron", "Vitamin C"},
		Allergies:               []string{"Peanuts"},
		RecentMeals:             []Meal{},
		Goals: []Goal{
			{Type: "Weight loss", Target: 10.0, Duration: 6},
		},
		MicrobiomeData: MicrobiomeData{
			DiversityScore:           7.2,
			GutHealthRecommendations: []string{"Increase fiber intake", "Add more fermented foods"},
		},
		BloodGlucose: 110.5,
		LipidProfile: LipidProfile{
			Cholesterol:   190.0,
			HDL:           55.0,
			LDL:           120.0,
			Triglycerides: 150.0,
		},
		EnvironmentalFactors: EnvironmentalFactors{
			Location: "New York",
			Climate:  "Temperate",
			Season:   "Winter",
		},
		MealHistory: []MealHistory{},
		HealthScore: 75.5,
	}

	meals := []Meal{
		{
			Name: "Avocado Toast",
			Tags: []string{"vegetarian", "gluten-free", "diabetes-friendly"},
			NutritionalContent: map[string]float64{
				"calories": 200, "protein": 6, "fat": 10, "fiber": 9, "sugar": 1,
			},
			Cuisine:      "American",
			HealthScores: map[string]float64{"diabetes_friendly": 9.0, "heart_healthy": 8.5},
		},
		{
			Name: "Chicken Salad",
			Tags: []string{"low-carb", "heart-healthy"},
			NutritionalContent: map[string]float64{
				"calories": 300, "protein": 25, "fat": 15, "fiber": 5, "sugar": 2,
			},
			Cuisine:      "Mediterranean",
			HealthScores: map[string]float64{"diabetes_friendly": 6.0, "heart_healthy": 9.0},
		},
		{
			Name: "Grilled Salmon",
			Tags: []string{"low-carb", "heart-healthy", "gluten-free"},
			NutritionalContent: map[string]float64{
				"calories": 400, "protein": 30, "fat": 20, "fiber": 0, "sugar": 0,
			},
			Cuisine:      "Mediterranean",
			HealthScores: map[string]float64{"diabetes_friendly": 8.0, "heart_healthy": 9.5},
		},
		{
			Name: "Vegan Buddha Bowl",
			Tags: []string{"vegan", "gluten-free", "diabetes-friendly"},
			NutritionalContent: map[string]float64{
				"calories": 350, "protein": 15, "fat": 10, "fiber": 12, "sugar": 3,
			},
			Cuisine:      "Asian",
			HealthScores: map[string]float64{"diabetes_friendly": 9.0, "heart_healthy": 8.0},
		},
		{
			Name: "Turkey Sandwich",
			Tags: []string{"high-protein", "low-fat"},
			NutritionalContent: map[string]float64{
				"calories": 320, "protein": 28, "fat": 8, "fiber": 4, "sugar": 3,
			},
			Cuisine:      "American",
			HealthScores: map[string]float64{"diabetes_friendly": 7.5, "heart_healthy": 8.0},
		},
		{
			Name: "Quinoa Salad",
			Tags: []string{"vegetarian", "gluten-free", "diabetes-friendly"},
			NutritionalContent: map[string]float64{
				"calories": 280, "protein": 10, "fat": 9, "fiber": 6, "sugar": 2,
			},
			Cuisine:      "Mediterranean",
			HealthScores: map[string]float64{"diabetes_friendly": 8.5, "heart_healthy": 8.0},
		},
		{
			Name: "Lentil Soup",
			Tags: []string{"vegetarian", "gluten-free", "high-fiber"},
			NutritionalContent: map[string]float64{
				"calories": 250, "protein": 12, "fat": 5, "fiber": 10, "sugar": 2,
			},
			Cuisine:      "Middle Eastern",
			HealthScores: map[string]float64{"diabetes_friendly": 7.0, "heart_healthy": 8.0},
		},
		{
			Name: "Beef Stir Fry",
			Tags: []string{"high-protein", "low-carb"},
			NutritionalContent: map[string]float64{
				"calories": 450, "protein": 35, "fat": 18, "fiber": 3, "sugar": 6,
			},
			Cuisine:      "Asian",
			HealthScores: map[string]float64{"diabetes_friendly": 6.5, "heart_healthy": 7.5},
		},
		{
			Name: "Tofu Scramble",
			Tags: []string{"vegan", "high-protein", "low-carb"},
			NutritionalContent: map[string]float64{
				"calories": 220, "protein": 14, "fat": 12, "fiber": 5, "sugar": 2,
			},
			Cuisine:      "American",
			HealthScores: map[string]float64{"diabetes_friendly": 8.0, "heart_healthy": 7.5},
		},
		{
			Name: "Eggplant Parmesan",
			Tags: []string{"vegetarian", "gluten-free", "low-carb"},
			NutritionalContent: map[string]float64{
				"calories": 360, "protein": 15, "fat": 18, "fiber": 7, "sugar": 5,
			},
			Cuisine:      "Italian",
			HealthScores: map[string]float64{"diabetes_friendly": 6.5, "heart_healthy": 7.0},
		},
	}

	// Get top 2 recommended meals for the user
	topMeals := GetTopMeals(user, meals, 2)

	// Display the top meals with their scores
	fmt.Println("Top Recommended Meals for", user.Name)
	fmt.Println("----------------------------")
	for _, mealWithScore := range topMeals {
		meal := mealWithScore.Meal
		fmt.Printf("Meal: %s | Score: %.2f\n", meal.Name, mealWithScore.Score)
		fmt.Printf("Tags: %v\n", meal.Tags)
		fmt.Printf("Nutritional Content: %v\n", meal.NutritionalContent)
		fmt.Printf("Health Scores: %v\n", meal.HealthScores)
		fmt.Println("----------------------------")
	}
}


```

