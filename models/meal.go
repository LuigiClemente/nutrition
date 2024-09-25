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
	Category           string         `json:"category"`                     // pasta, pizza, stew
	MealType           pq.StringArray `gorm:"type:text[]" json:"meal_type"` // diner, breackfast, snack, luanch
	Cuisine            string         `json:"cuisine"`
	Tags               pq.StringArray `gorm:"type:text[]" json:"tags"`
}

type Ingredient struct {
	ID          uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	MealID      uint    `json:"meal_id"`
	Name        string  `json:"name"`
	Amount      float64 `json:"amount"` // Total amount used in grams 
	Portion     string  `json:"portion"`
	PortionUnit string  `json:"portion_unit"` // Measurement unit
}

type ScoredMeal struct {
	Meal  Meal    `json:"meal" gorm:"foreignKey:MealID"`
	Score float64 `json:"score"`
}

type Course struct {
	Meal  Meal    `json:"meal"`
	Score float64 `json:"score"`
}

type MealResponse struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name" binding:"required"`
	Ingredients []IngredientReponse `json:"ingredients" gorm:"foreignKey:MealID"`
}

type IngredientReponse struct {
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`  // Total amount used
	Portion string  `json:"portion"` // Measurement unit
}

type Recommended struct {
	Type    string   `json:"type"`
	Courses []Course `json:"courses"`
}

type UserHealthInfoResponse struct {
	User                   User        `json:"user_info"`
	ScoreAndRecommendation Recommended `json:"recommendation"`
}

type MealForOption struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
