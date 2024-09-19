package models

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

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
}

type Ingredient struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	MealID      uint           `json:"meal_id"`
	Name        string         `json:"name"`
	Amount      float64        `json:"amount"`      // Total amount used
	Unit        string         `json:"unit"`        // Measurement unit
	Nutritional datatypes.JSON `json:"nutritional"` // Nutritional info as JSON
	Percentage  float64        `json:"percentage"`  //The proportion of that ingredient in the dish
}

type ScoredMeal struct {
	Meal  Meal    `json:"meal" gorm:"foreignKey:MealID"`
	Score float64 `json:"score"`
}

type Course struct {
	MealName    string       `json:"meal_name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Recommended struct {
	Type    string   `json:"type"`
	Courses []Course `json:"courses"`
	Score   float64  `json:"score"`
}

type UserHealthInfoResponse struct {
	User                  User          `json:"user_info"`
	ScoreAndRecmensdtions []Recommended `json:"recommendations"`
}

type MealForOption struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
