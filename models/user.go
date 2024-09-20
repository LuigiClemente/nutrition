package models

import (
	"time"

	"github.com/lib/pq"
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
	UserID                   uint           `json:"user_id"`
	DiversityScore           float64        `json:"diversity_score"`
	GutHealthRecommendations pq.StringArray `gorm:"type:text[]" json:"gut_health_recommendations"`
}

type LipidProfile struct {
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
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id"`
	MealID    uint      `json:"meal_id"` // Foreign key for Meal
	Timestamp time.Time `json:"timestamp"`
}

type MealHistory struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id"`
	MealID    uint      `json:"meal_id"` // Foreign key for Meal
	Timestamp time.Time `json:"timestamp"`
}

type RequestedMeal struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	MealType  string    `json:"meal_type"`
	Timestamp time.Time `json:"timestamp"`
}

type UserPreferences struct {
	ID                  uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID              uint           `json:"user_id"`
	PreferredCuisines   pq.StringArray `gorm:"type:text[]" json:"preferred_cuisines"`
	MealTimings         pq.StringArray `gorm:"type:text[]" json:"meal_timings"`
	FavoriteIngredients pq.StringArray `gorm:"type:text[]" json:"favorite_ingredients"`
}
