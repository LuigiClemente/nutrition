package models

import (
	"gorm.io/datatypes"
)

type Meal struct {
	ID                 uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name               string         `json:"name" binding:"required"`
	Ingredients        []Ingredient   `json:"ingredients" gorm:"foreignKey:MealID"`
	NutritionalContent datatypes.JSON `json:"nutritional_content"`
	MealCategoryID     uint           `json:"meal_category_id"`                               // Foreign key for MealCategory
	MealCategory       MealCategory   `gorm:"foreignKey:MealCategoryID" json:"meal_category"` // Relationship to MealCategory
	MealTypeID         uint           `json:"meal_type_id"`                                   // Foreign key for MealType
	MealType           MealType       `gorm:"foreignKey:MealTypeID" json:"meal_type"`         // Relationship to MealType                     // diner, breackfast, snack, luanch
	Cuisine            string         `json:"cuisine"`
	MealTags           []MealTag      `json:"tags" gorm:"many2many:meal_tag_relationship;joinForeignKey:MealID;JoinReferences:TagID"`
}

type Ingredient struct {
	ID      uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	MealID  uint    `json:"meal_id"`
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"` // Total amount used in grams
	Portion string  `json:"portion"`
}

type ScoredMeal struct {
	Meal  Meal    `json:"meal" gorm:"foreignKey:MealID"`
	Score float64 `json:"score"`
}

type MealType struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Type string `json:"type"`
}

type MealCategory struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Category string `json:"category"`
}

type MealTag struct {
	ID  uint   `json:"id"`
	Tag string `json:"tag"`
}

type MealTagRelationship struct {
	MealID uint `gorm:"primaryKey"`
	TagID  uint `gorm:"primaryKey"`
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
	Ounces  string  `json:"ounces"`  // Measurement unit
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

type Recommended struct {
	Meal  MealResponse `json:"meal"`
	Score float64      `json:"score"`
}
