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
	Course             string         `json:"course"`
	MeaTimings         pq.StringArray `gorm:"type:text[]" json:"meal_timings"` // diner, breackfast, snack, luanch
	MealTypeID         uint           `json:"meal_type_id"`                    // Foreign key for MealType
	MealType           MealType       `gorm:"foreignKey:MealTypeID" json:"meal_type"`
	Cuisine            string         `json:"cuisine"`
	MealTags           []MealTag      `json:"tags" gorm:"many2many:meal_tag_relationships;joinForeignKey:MealID;JoinReferences:TagID"`
}

type Ingredient struct {
	ID      uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	MealID  uint    `json:"meal_id"`
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`  // Total amount used
	Unit    string  `json:"unit"`    // Measurement unit
	Portion string  `json:"portion"` // Measurement unit (e.g. "2 tbsp", "1/2 cup")
}

type ScoredMeal struct {
	Meal  *Meal    `json:"meal" gorm:"foreignKey:MealID"`
	Score float64 `json:"score"`
}

type MealType struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Type string `json:"type"`
}

type MealTag struct {
	ID  uint   `json:"id"`
	Tag string `json:"tag"`
}

type MealTagRelationship struct {
	MealID uint `json:"meal_id" gorm:"primaryKey"`
	TagID  uint `json:"tag_id" gorm:"primaryKey"`
}

type MealResponse struct {
	ID          uint                 `json:"meal_id"`
	Course      string               `json:"course"`
	Name        string               `json:"meal_name" binding:"required"`
	Ingredients []IngredientResponse `json:"ingredients" gorm:"foreignKey:MealID"`
	Score       float64              `json:"score"` // Score for this meal
}

type IngredientResponse struct {
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`  // Total amount used
	Unit    string  `json:"unit"`    // Measurement unit
	Portion string  `json:"portion"` // Measurement unit
	Ounces  string  `json:"ounces"`  // Measurement unit
}

type UserHealthInfoResponse struct {
	User                   User             `json:"user_info"`
	ScoreAndRecommendation []Recommendation `json:"recommendations"`
}

type MealForOption struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Recommendation struct {
	Combination string         `json:"combination"` // One Course, Two Courses, etc.
	Courses     []MealResponse `json:"courses"`     // List of courses in this recommendation

}

