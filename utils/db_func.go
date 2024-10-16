package utils

import (
	"fmt"
	"log"
	"nutrition/models"
	"time"

	"gorm.io/gorm"
)

// deleteOldRecentMeals deletes meal records older than 7 days from the recent meals table.
func DeleteOldRecentMeals(db *gorm.DB) error {
	// Calculate the cutoff date (7 days ago from now)
	cutoffDate := time.Now().AddDate(0, 0, -7)

	// Perform the delete operation
	if err := db.Where("timestamp < ?", cutoffDate).Delete(&models.RecentMeals{}).Error; err != nil {
		return fmt.Errorf("failed to delete old recent meals: %w", err)
	}

	log.Println("Deleted old recent meals successfully")
	return nil
}
