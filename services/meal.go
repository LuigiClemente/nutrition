package services

import (
	"fmt"
	"nutrition/models"
)

func (s *Service) PostMeals(meals []models.Meal) (*[]models.Meal, error) {
	if err := s.db.CreateInBatches(&meals, 100).Error; err != nil {
		return nil, err
	}

	return &meals, nil
}

func (s *Service) GetMeal() (*[]models.Meal, error) {
	var meals []models.Meal
	if err := s.db.Preload("MealType").
		Preload("MealTags").
		Preload("Ingredients").
		Order("id DESC").Find(&meals).Error; err != nil {
		return nil, err
	}
	return &meals, nil
}

func (s *Service) GetMealsByType(typeId int) (*[]models.Meal, error) {
	var meals []models.Meal
	if err := s.db.Where("meal_type_id = ?", typeId).
		Preload("MealType").
		Preload("MealTags").
		Preload("Ingredients").Order("id DESC").Find(&meals).Error; err != nil {
		return nil, err
	}

	return &meals, nil
}

func (s *Service) GetMealsByTypeAndCourse(starterTypeId, mainTypeId, dessertTypeId, numberOfStarter, numberOfMain, numberOfDessert int) (*[]models.Meal, error) {
	var meals []models.Meal
	var courses []string
	var mealTypeIDs []int

	// Collect meal type IDs based on the input
	if numberOfStarter > 0 {
		courses = append(courses, "Starter")
		mealTypeIDs = append(mealTypeIDs, starterTypeId)
	}
	if numberOfMain > 0 {
		courses = append(courses, "Main")
		mealTypeIDs = append(mealTypeIDs, mainTypeId)
	}
	if numberOfDessert > 0 {
		courses = append(courses, "Dessert")
		mealTypeIDs = append(mealTypeIDs, dessertTypeId)
	}

	fmt.Println("number of dessert: ", numberOfDessert)

	// Adjust the query to handle multiple courses
	query := s.db.Where("meal_type_id IN ? AND course IN ?", mealTypeIDs, courses)

	// Only preload related models if needed
	query = query.Preload("MealType").
		Preload("MealTags").
		Preload("Ingredients")

	// Add pagination and sorting
	if err := query.Order("id DESC").Find(&meals).Error; err != nil {
		return nil, err
	}

	return &meals, nil
}


func (s *Service) GetMealTypes() (*[]models.MealType, error) {
	var types []models.MealType
	if err := s.db.
		Order("type ASC").
		Find(&types).Error; err != nil {
		return nil, err
	}
	return &types, nil
}

func (s *Service) GetMealForOption() (*[]models.MealForOption, error) {
	var factors []models.MealForOption
	if err := s.db.Model(&models.Meal{}).Order("id DESC").Find(&factors).Error; err != nil {
		return nil, err
	}
	return &factors, nil
}

func (s *Service) GetMealMealId(id int) (*models.Meal, error) {
	var meal models.Meal
	if err := s.db.Where("id = ?", id).
		Preload("MealType").
		Preload("Ingredients").
		Preload("MealTags").First(&meal).Error; err != nil {
		return nil, err
	}

	return &meal, nil
}

func (s *Service) PutMealMealId(mealId int, meal models.Meal) (*models.Meal, error) {
	// Start a transaction
	tx := s.db.Begin()

	// Update the main meal fields that have changed
	if err := tx.Model(&models.Meal{}).Where("id = ?", mealId).Updates(&meal).Error; err != nil {
		tx.Rollback() // Rollback transaction on error
		return nil, err
	}

	// Get existing ingredients for the meal in a map for quick lookup
	var existingIngredients []models.Ingredient
	if err := tx.Where("meal_id = ?", mealId).Find(&existingIngredients).Error; err != nil {
		tx.Rollback() // Rollback transaction on error
		return nil, err
	}

	// Map existing ingredients by ID for efficient comparison
	existingIngredientsMap := make(map[uint]models.Ingredient)
	for _, ingredient := range existingIngredients {
		existingIngredientsMap[ingredient.ID] = ingredient
	}

	// Create slices for batch insert, update, and delete
	var ingredientsToInsert []models.Ingredient
	var ingredientsToUpdate []models.Ingredient
	ingredientIDsToKeep := make(map[uint]bool)

	for _, newIngredient := range meal.Ingredients {
		newIngredient.MealID = uint(mealId) // Ensure the MealID is set

		if existingIngredient, found := existingIngredientsMap[newIngredient.ID]; found {
			// Check if there's a change before updating
			if newIngredient.Name != existingIngredient.Name ||
				newIngredient.Amount != existingIngredient.Amount ||
				newIngredient.Unit != existingIngredient.Unit ||
				newIngredient.Portion != existingIngredient.Portion {
				ingredientsToUpdate = append(ingredientsToUpdate, newIngredient)
			}
			// Mark ingredient as one to keep
			ingredientIDsToKeep[newIngredient.ID] = true
		} else {
			// New ingredient to insert
			ingredientsToInsert = append(ingredientsToInsert, newIngredient)
		}
	}

	// Perform batch updates for changed ingredients
	if len(ingredientsToUpdate) > 0 {
		if err := tx.Save(&ingredientsToUpdate).Error; err != nil {
			tx.Rollback() // Rollback on error
			return nil, err
		}
	}

	// Perform batch insert for new ingredients
	if len(ingredientsToInsert) > 0 {
		if err := tx.Create(&ingredientsToInsert).Error; err != nil {
			tx.Rollback() // Rollback on error
			return nil, err
		}
	}

	// Delete ingredients that are no longer part of the meal
	var ingredientIDs []uint
	for _, ingredient := range existingIngredients {
		if !ingredientIDsToKeep[ingredient.ID] {
			ingredientIDs = append(ingredientIDs, ingredient.ID)
		}
	}
	if len(ingredientIDs) > 0 {
		if err := tx.Where("meal_id = ? AND id IN (?)", mealId, ingredientIDs).Delete(&models.Ingredient{}).Error; err != nil {
			tx.Rollback() // Rollback on error
			return nil, err
		}
	}

	// Manage tags
	// Delete all existing tags for the meal
	if err := tx.Where("meal_id = ?", mealId).Delete(&models.MealTagRelationship{}).Error; err != nil {
		tx.Rollback() // Rollback on error
		return nil, err
	}

	// Insert new tags
	var tagsToInsert []models.MealTagRelationship
	for _, tag := range meal.MealTags {
		// Create new tag relationship
		tagsToInsert = append(tagsToInsert, models.MealTagRelationship{
			MealID: uint(mealId),
			TagID:  tag.ID, // Assuming tag.ID is the identifier for existing tags
		})
	}

	if len(tagsToInsert) > 0 {
		if err := tx.Create(&tagsToInsert).Error; err != nil {
			tx.Rollback() // Rollback on error
			return nil, err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Fetch the updated meal with ingredients
	updatedMeal, err := s.GetMealMealId(mealId)
	if err != nil {
		return nil, err
	}

	return updatedMeal, nil
}

func (s *Service) DeleteMealMealId(mealId int) error {
	if err := s.db.Where("id = ?", mealId).Delete(&models.Meal{}).Error; err != nil {
		return err
	}
	return nil
}
