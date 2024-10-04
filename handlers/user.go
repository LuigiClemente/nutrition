package handlers

import (
	"net/http"
	"nutrition/models"
	"nutrition/utils"

	"github.com/gin-gonic/gin"
)

// GetUser returns all user health information
func (h *Handler) GetUser(c *gin.Context) {
	healthInfo, err := h.service.GetUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, healthInfo)
}

// PostUser creates a new user health information record
func (h *Handler) PostUser(c *gin.Context) {
	var info models.User
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate: max NumberOfStarter, NumberOfMain, and NumberOfDessert is 3 and min is 0
	if info.RequestedMeal.NumberOfStarter < 0 || info.RequestedMeal.NumberOfStarter > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NumberOfStarter must be between 0 and 3"})
		return
	}

	if info.RequestedMeal.NumberOfMain < 0 || info.RequestedMeal.NumberOfMain > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NumberOfMain must be between 0 and 3"})
		return
	}

	if info.RequestedMeal.NumberOfDessert < 0 || info.RequestedMeal.NumberOfDessert > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NumberOfDessert must be between 0 and 3"})
		return
	}

	postedUserHealthDataInfo, err := h.service.PostUser(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, postedUserHealthDataInfo)
}

// DeleteUserUserId deletes user health information by user ID
func (h *Handler) DeleteUserUserId(c *gin.Context) {
	id, err := utils.ParseAndValidateID(c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.service.DeleteUserUserId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User health info deleted successfully"})
}

// GetUserUserId gets user health information by user ID
func (h *Handler) GetUserUserId(c *gin.Context) {

	id, err := utils.ParseAndValidateID(c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	info, err := h.service.GetUserUserId(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, info)
}

// PutUserUserId updates user health information by user ID
func (h *Handler) PutUserUserId(c *gin.Context) {
	var info models.User
	
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := utils.ParseAndValidateID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	// Validate: max NumberOfStarter, NumberOfMain, and NumberOfDessert is 3 and min is 0
	if info.RequestedMeal.NumberOfStarter < 0 || info.RequestedMeal.NumberOfStarter > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NumberOfStarter must be between 0 and 3"})
		return
	}

	if info.RequestedMeal.NumberOfMain < 0 || info.RequestedMeal.NumberOfMain > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NumberOfMain must be between 0 and 3"})
		return
	}

	if info.RequestedMeal.NumberOfDessert < 0 || info.RequestedMeal.NumberOfDessert > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NumberOfDessert must be between 0 and 3"})
		return
	}

	updatedInfo, err := h.service.PutUserUserId(id, info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedInfo)
}
