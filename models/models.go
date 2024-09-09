package models

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type User struct {
	ID                      uint                 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name                    string               `json:"name" binding:"required"`
	Age                     int                  `json:"age" binding:"required"`
	Gender                  string               `json:"gender" binding:"required"`
	BodyMetrics             BodyMetrics          `json:"body_metrics" gorm:"foreignKey:UserID"`
	ActivityLevel           string               `json:"activity_level"`
	DietaryPreferences      DietaryPreferences   `json:"dietary_preferences" gorm:"foreignKey:UserID"`
	HealthConditions        []HealthCondition    `json:"health_conditions" gorm:"foreignKey:UserID"`
	NutritionalDeficiencies pq.StringArray       `gorm:"type:text[]" json:"nutritional_deficiencies"`
	Allergies               pq.StringArray       `gorm:"type:text[]" json:"allergies"`
	RecentMeals             []RecentMeals        `json:"recent_meals" gorm:"foreignKey:UserID"`
	Goals                   []Goal               `json:"goals" gorm:"foreignKey:UserID"`
	MicrobiomeData          MicrobiomeData       `json:"microbiome_data" gorm:"foreignKey:UserID"`
	BloodGlucose            float64              `json:"blood_glucose"`
	LipidProfile            LipidProfile         `json:"lipid_profiles" gorm:"foreignKey:UserID"`
	EnvironmentalFactors    EnvironmentalFactors `json:"environmental_factors" gorm:"foreignKey:UserID"`
	MealHistory             []MealHistory        `json:"meal_histories" gorm:"foreignKey:UserID"`
	HealthScore             float64              `json:"health_score"`
	Preferences             UserPreferences      `json:"preferences" gorm:"foreignKey:UserID"`
}

type BodyMetrics struct {
	UserID uint    `json:"user_id"`
	Weight float64 `json:"weight"` // kg
	Height float64 `json:"height"` // cm
	BMI    float64 `json:"bmi"`
}

type DietaryPreferences struct {
	UserID             uint           `json:"user_id"`
	Vegetarian         bool           `json:"vegetarian"`
	Vegan              bool           `json:"vegan"`
	GlutenFree         bool           `json:"gluten_free"`
	DairyFree          bool           `json:"dairy_free"`
	SpecificAvoidances pq.StringArray `gorm:"type:text[]" json:"specific_avoidances"`
}

type HealthCondition struct {
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Severity string `json:"severity"`
}

type Goal struct {
	ID       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID   uint    `json:"user_id"`
	Type     string  `json:"type"`
	Target   float64 `json:"target"`
	Duration int     `json:"duration"`
}

type MicrobiomeData struct {
	ID                       uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID                   uint           `json:"user_id"`
	DiversityScore           float64        `json:"diversity_score"`
	GutHealthRecommendations pq.StringArray `gorm:"type:text[]" json:"gut_health_recommendations"`
}

type LipidProfile struct {
	ID            uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID        uint    `json:"user_id"`
	Cholesterol   float64 `json:"cholesterol"`
	HDL           float64 `json:"hdl"`
	LDL           float64 `json:"ldl"`
	Triglycerides float64 `json:"triglycerides"`
}

type EnvironmentalFactors struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID   uint   `json:"user_id"`
	Location string `json:"location"`
	Climate  string `json:"climate"`
	Season   string `json:"season"`
}

type RecentMeals struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"user_id"`
	MealID    uint   `json:"meal_id"` // Foreign key for Meal
	Timestamp string `json:"timestamp"`
}

type MealHistory struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"user_id"`
	MealID    uint   `json:"meal_id"` // Foreign key for Meal
	Timestamp string `json:"timestamp"`
}

type UserPreferences struct {
	ID                  uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID              uint           `json:"user_id"`
	PreferredCuisines   pq.StringArray `gorm:"type:text[]" json:"preferred_cuisines"`
	MealTimings         pq.StringArray `gorm:"type:text[]" json:"meal_timings"`
	FavoriteIngredients pq.StringArray `gorm:"type:text[]" json:"favorite_ingredients"`
}

type Meal struct {
	ID                 uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name               string         `json:"name" binding:"required"`
	Ingredients        []Ingredient   `json:"ingredients" gorm:"foreignKey:MealID"`
	NutritionalContent datatypes.JSON `json:"nutritional_content"`
	Category           string         `json:"category"`
	MealType           pq.StringArray `gorm:"type:text[]" json:"meal_type"`
	Cuisine            string         `json:"cuisine"`
	Tags               pq.StringArray `gorm:"type:text[]" json:"tags"`
	HealthScores       datatypes.JSON `json:"health_scores"`
	PreparationTime    int            `json:"preparation_time"`
	Difficulty         string         `json:"difficulty"`
	ServingSize        int            `json:"serving_size"`
	Instructions       string         `json:"instructions"`
	PortionSize        float64        `json:"portion_size"` // Total portion size calculated
}

type MealAdjustment struct {
	MealID                  uint       `json:"meal_id"`
	AddedIngredient         Ingredient `json:"add_ingredient"`
	RemovedIngredient       Ingredient `json:"remove_ingredient"`
	RecalculatedPortionSize float64    `json:"recalculated_portion_size"`
}

type MealReceipt struct {
	MealID                 uint    `json:"meal_id"`
	PreExistingCalculation bool    `json:"pre_existing_calculation"`
	TotalPortionSize       float64 `json:"total_portion_size"`
}

type Ingredient struct {
	ID                 uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	MealID             uint           `json:"meal_id"`
	Name               string         `json:"name"`
	Amount             float64        `json:"amount"`                // Total amount used
	Unit               string         `json:"unit"`                  // Measurement unit
	PortionSizePer100g float64        `json:"portion_size_per_100g"` // Portion size per 100 grams
	Nutritional        datatypes.JSON `json:"nutritional"`           // Nutritional info as JSON
}

type ScoredMeal struct {
	Meal  Meal    `json:"meal" gorm:"foreignKey:MealID"`
	Score float64 `json:"score"`
}

type Recommended struct {
	Meal               string  `json:"meal"`
	Score              float64 `json:"score"`
	Tags               []string
	NutritionalContent datatypes.JSON `json:"nutritional_content"`
	HealthScores       datatypes.JSON `json:"health_scores"`
}

type UserHealthInfoResponse struct {
	User                  User          `json:"user_info"`
	ScoreAndRecmensdtionS []Recommended `json:"recommendations"`
}
