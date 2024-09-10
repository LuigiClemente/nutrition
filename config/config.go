package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ScoringConfig struct {
	DietaryPreferencesScores map[string]float64 `json:"dietary_preferences_scores"`
	HealthGoalsScores        map[string]float64 `json:"health_goals_scores"`
	DefaultScores            map[string]float64 `json:"default_scores"`
}

// LoadConfig reads the config file and unmarshals it into ScoringConfig
func LoadConfig(filepath string) (ScoringConfig, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return ScoringConfig{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	var config ScoringConfig
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return ScoringConfig{}, fmt.Errorf("error decoding config file: %w", err)
	}

	return config, nil
}
