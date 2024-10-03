package handlers

import (
	"net/http"
	"nutrition/models"
	"nutrition/utils"

	"github.com/gin-gonic/gin"
)

// GetMeal returns all Meal
func (h *Handler) GetMeal(c *gin.Context) {
	items, err := h.service.GetMeal()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetMealForOption returns all Meal for listing in options value
func (h *Handler) GetMealForOption(c *gin.Context) {
	items, err := h.service.GetMealForOption()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetMealTypes returns all Meal for listing in options value
func (h *Handler) GetMealTypes(c *gin.Context) {
	categories, err := h.service.GetMealTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// PostMeal creates a new Meal record
func (h *Handler) PostMeal(c *gin.Context) {
	var item []models.Meal
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postedItem, err := h.service.PostMeals(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, postedItem)
}

// DeleteMealMealId deletes Meal by meal ID
func (h *Handler) DeleteMealMealId(c *gin.Context) {
	id, err := utils.ParseAndValidateID(c.Param("mealId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meal ID"})
		return
	}

	err = h.service.DeleteMealMealId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Meal deleted successfully"})
}

// GetMealMealId gets Meal by user ID
func (h *Handler) GetMealMealId(c *gin.Context) {
	id, err := utils.ParseAndValidateID(c.Param("mealId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	item, err := h.service.GetMealMealId(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// PutMealMealId updates Meal by user ID
func (h *Handler) PutMealMealId(c *gin.Context) {
	var item models.Meal
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := utils.ParseAndValidateID(c.Param("mealId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	updatedItem, err := h.service.PutMealMealId(id, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedItem)
}
