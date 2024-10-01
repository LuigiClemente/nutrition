package utils

import (
	"log"
	"nutrition/database"
	"time"
)

// ScheduleDeletion starts the deletion process once and schedules it to run every day at 12 PM.
func ScheduleDeletion() {
	// Retrieve database instance once
	db := database.GetDatabaseInstance().DB

	// Run deletion process immediately when the application starts
	if err := DeleteOldRecentMeals(db); err != nil {
		log.Printf("Error during initial deletion: %v", err)
	}

	// Run the deletion process every 24 hours at 12 PM
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			now := time.Now()
			nextRun := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.Local)

			// If it's after 12 PM, schedule for tomorrow
			if now.After(nextRun) {
				nextRun = nextRun.Add(24 * time.Hour)
			}

			// Calculate the duration until the next scheduled time
			duration := nextRun.Sub(now)

			// Wait until the next run time
			time.Sleep(duration)

			// Execute the deletion process
			log.Println("Running deletion process at 12 PM")
			if err := DeleteOldRecentMeals(db); err != nil {
				log.Printf("Error during scheduled deletion: %v", err)
			}
		}
	}()
}
