package services

import (
	"fmt"
	"log"
	"nutrition/models"
	"nutrition/utils"
	"strings"
)

// =============== User Health Info ===============
func (s *Service) PostUser(userHealthInfo models.User) (*models.UserHealthInfoResponse, error) {
	// Start a transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Panic recovered:", r)
		}
	}()

	if tx.Error != nil {
		return nil, &CustomError{"Failed to begin transaction", tx.Error}
	}

	// Step 1: Insert into users table and return userID
	userSQL := `INSERT INTO users (name, age, gender, activity_level, blood_glucose, health_score, nutritional_deficiencies, allergies) 
               VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING id`
	var userID uint
	if err := tx.Raw(userSQL, userHealthInfo.Name, userHealthInfo.Age, userHealthInfo.Gender, userHealthInfo.ActivityLevel, userHealthInfo.BloodGlucose, userHealthInfo.HealthScore, userHealthInfo.NutritionalDeficiencies, userHealthInfo.Allergies).Scan(&userID).Error; err != nil {
		tx.Rollback()
		return nil, &CustomError{"Failed to insert user", err}
	}
	userHealthInfo.ID = userID

	// Reusable function for batch inserts
	batchInsert := func(table string, columns string, values []interface{}, valuePlaceholders string) error {
		insertSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", table, columns, valuePlaceholders)
		return tx.Exec(insertSQL, values...).Error
	}

	// Step 2: Insert into body_metrics table
	if err := tx.Exec(`INSERT INTO body_metrics (user_id, weight, height) VALUES (?, ?, ?)`, userID, userHealthInfo.BodyMetrics.Weight, userHealthInfo.BodyMetrics.Height).Error; err != nil {
		tx.Rollback()
		return nil, &CustomError{"Failed to insert body metrics", err}
	}

	// Step 3: Insert into dietary_preferences table
	if err := tx.Exec(`INSERT INTO dietary_preferences (user_id, vegetarian, vegan, gluten_free, dairy_free, specific_avoidances) 
                      VALUES (?, ?, ?, ?, ?, ?)`, userID, userHealthInfo.DietaryPreferences.Vegetarian, userHealthInfo.DietaryPreferences.Vegan, userHealthInfo.DietaryPreferences.GlutenFree, userHealthInfo.DietaryPreferences.DairyFree, userHealthInfo.DietaryPreferences.SpecificAvoidances).Error; err != nil {
		tx.Rollback()
		return nil, &CustomError{"Failed to insert dietary preferences", err}
	}

	// Batch inserts for health_conditions
	if len(userHealthInfo.HealthConditions) > 0 {
		var healthConditionsValues []interface{}
		placeholders := make([]string, len(userHealthInfo.HealthConditions))
		for i, condition := range userHealthInfo.HealthConditions {
			placeholders[i] = "(?, ?, ?)"
			healthConditionsValues = append(healthConditionsValues, userID, condition.Name, condition.Severity)
		}
		if err := batchInsert("health_conditions", "user_id, name, severity", healthConditionsValues, strings.Join(placeholders, ",")); err != nil {
			tx.Rollback()
			return nil, &CustomError{"Failed to insert health conditions", err}
		}
	}

	// Batch inserts for goals
	if len(userHealthInfo.Goals) > 0 {
		var goalsValues []interface{}
		placeholders := make([]string, len(userHealthInfo.Goals))
		for i, goal := range userHealthInfo.Goals {
			placeholders[i] = "(?, ?, ?, ?)"
			goalsValues = append(goalsValues, userID, goal.Type, goal.Target, goal.Duration)
		}
		if err := batchInsert("goals", "user_id, type, target, duration", goalsValues, strings.Join(placeholders, ",")); err != nil {
			tx.Rollback()
			return nil, &CustomError{"Failed to insert goals", err}
		}
	}

	// Insert microbiome data
	if err := tx.Exec(`INSERT INTO microbiome_data (user_id, diversity_score, gut_health_recommendations) 
                      VALUES (?, ?, ?)`, userID, userHealthInfo.MicrobiomeData.DiversityScore, userHealthInfo.MicrobiomeData.GutHealthRecommendations).Error; err != nil {
		tx.Rollback()
		return nil, &CustomError{"Failed to insert microbiome data", err}
	}

	// Insert lipid profile data
	if err := tx.Exec(`INSERT INTO lipid_profiles (user_id, cholesterol, hdl, ldl, triglycerides) 
                      VALUES (?, ?, ?, ?, ?)`, userID, userHealthInfo.LipidProfile.Cholesterol, userHealthInfo.LipidProfile.HDL, userHealthInfo.LipidProfile.LDL, userHealthInfo.LipidProfile.Triglycerides).Error; err != nil {
		tx.Rollback()
		return nil, &CustomError{"Failed to insert lipid profile data", err}
	}

	// Insert environmental factors
	if err := tx.Exec(`INSERT INTO environmental_factors (user_id, location, climate, season) 
                      VALUES (?, ?, ?, ?)`, userID, userHealthInfo.EnvironmentalFactors.Location, userHealthInfo.EnvironmentalFactors.Climate, userHealthInfo.EnvironmentalFactors.Season).Error; err != nil {
		tx.Rollback()
		return nil, &CustomError{"Failed to insert environmental factors", err}
	}

	// Batch insert meal histories
	if len(userHealthInfo.MealHistory) > 0 {
		var mealHistoryValues []interface{}
		placeholders := make([]string, len(userHealthInfo.MealHistory))
		for i, mealHistory := range userHealthInfo.MealHistory {
			placeholders[i] = "(?, ?, ?)"
			mealHistoryValues = append(mealHistoryValues, userID, mealHistory.MealID, mealHistory.Timestamp)
		}
		if err := batchInsert("meal_histories", "user_id, meal_id, timestamp", mealHistoryValues, strings.Join(placeholders, ",")); err != nil {
			tx.Rollback()
			return nil, &CustomError{"Failed to insert meal histories", err}
		}
	}

	// Batch insert recent meals
	if len(userHealthInfo.RecentMeals) > 0 {
		var recentMealsValues []interface{}
		placeholders := make([]string, len(userHealthInfo.RecentMeals))
		for i, recentMeal := range userHealthInfo.RecentMeals {
			placeholders[i] = "(?, ?, ?)"
			recentMealsValues = append(recentMealsValues, userID, recentMeal.MealID, recentMeal.Timestamp)
		}
		if err := batchInsert("recent_meals", "user_id, meal_id, timestamp", recentMealsValues, strings.Join(placeholders, ",")); err != nil {
			tx.Rollback()
			return nil, &CustomError{"Failed to insert recent meals", err}
		}
	}

	// Insert user preferences
	if err := tx.Exec(`INSERT INTO user_preferences (user_id, preferred_cuisines, meal_timings, favorite_ingredients) 
                      VALUES (?, ?, ?, ?)`, userID, userHealthInfo.Preferences.PreferredCuisines, userHealthInfo.Preferences.MealTimings, userHealthInfo.Preferences.FavoriteIngredients).Error; err != nil {
		tx.Rollback()
		return nil, &CustomError{"Failed to insert user preferences", err}
	}

	RequestedMealSQL := `INSERT INTO requested_meals (
		user_id, 
		starter_meal_type, 
		main_meal_type, 
		dessert_meal_type, 
		number_of_starter, 
		number_of_main, 
		number_of_dessert
	) VALUES (?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT (user_id) DO UPDATE
	SET 
		starter_meal_type = EXCLUDED.starter_meal_type, 
		main_meal_type = EXCLUDED.main_meal_type, 
		dessert_meal_type = EXCLUDED.dessert_meal_type, 
		number_of_starter = EXCLUDED.number_of_starter, 
		number_of_main = EXCLUDED.number_of_main, 
		number_of_dessert = EXCLUDED.number_of_dessert`
	
	if err := tx.Exec(RequestedMealSQL,
		userID,
		userHealthInfo.RequestedMeal.StarterMealType,
		userHealthInfo.RequestedMeal.MainMealType,
		userHealthInfo.RequestedMeal.DessertMealType,
		userHealthInfo.RequestedMeal.NumberOfStarter,
		userHealthInfo.RequestedMeal.NumberOfMain,
		userHealthInfo.RequestedMeal.NumberOfDessert,
	).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, &CustomError{"Failed to commit transaction", err}
	}

	// Fetch meals and calculate recommendations concurrently
	resultChan := make(chan []models.Recommendation, 1)
	errorChan := make(chan error, 1)

	go func() {
		// Fetch meals based on the requested meal type and number of courses
		meals, err := s.GetMealsByTypeAndCourse(userHealthInfo.RequestedMeal.StarterMealType,
			userHealthInfo.RequestedMeal.MainMealType,
			userHealthInfo.RequestedMeal.DessertMealType, userHealthInfo.RequestedMeal.NumberOfStarter, userHealthInfo.RequestedMeal.NumberOfMain, userHealthInfo.RequestedMeal.NumberOfDessert)
		if err != nil {
			errorChan <- &CustomError{"Failed to fetch meals", err}
			return
		}

		// Calculate  meals, specifying the number of courses needed
		calculateMeals := utils.CalculateMealScore(userHealthInfo, *meals)

		// Prepare recommendations based on the number of courses
		var scoreRecommendations []models.Recommendation

		// Generate course combinations based on the requested number of courses
		courseCombinations := utils.GenerateCourseCombinations(calculateMeals, userHealthInfo.RequestedMeal.NumberOfStarter, userHealthInfo.RequestedMeal.NumberOfMain, userHealthInfo.RequestedMeal.NumberOfDessert)

		// Check for valid number of course combinations
		totalCourses := userHealthInfo.RequestedMeal.NumberOfStarter + userHealthInfo.RequestedMeal.NumberOfMain + userHealthInfo.RequestedMeal.NumberOfDessert

		// Create a recommendation with the specified number of courses
		var courses []models.MealResponse

		// Add the correct number of courses for this recommendation
		for j := 0; j < totalCourses && (j) < len(courseCombinations); j++ {
			topMeal := courseCombinations[j]
			courses = append(courses, models.MealResponse{
				ID:          topMeal.Meal.ID,
				Course:      topMeal.Meal.Course,
				Name:        topMeal.Meal.Name,
				Score:       topMeal.Score,
				Ingredients: mapIngredientsToResponse(topMeal.Meal.Ingredients),
			})
		}

		// Create the recommendation and add it to the list
		recommendation := models.Recommendation{
			Combination: fmt.Sprintf("%d Meal(s)", len(courses)),
			Courses:     courses,
		}
		scoreRecommendations = append(scoreRecommendations, recommendation)

		// Send the result to the result channel
		resultChan <- scoreRecommendations
	}()

	// Wait for the result from the goroutine or handle errors
	select {
	case scoreRecommendations := <-resultChan:
		response := &models.UserHealthInfoResponse{
			User:                   userHealthInfo,
			ScoreAndRecommendation: scoreRecommendations,
		}
		return response, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (s *Service) GetUser() (*[]models.User, error) {
	var userHealthInfo []models.User

	// Fetch users with related data using Preload for eager loading
	if err := s.db.Preload("BodyMetrics").
		Preload("DietaryPreferences").
		Preload("Preferences").
		Preload("HealthConditions").
		Preload("MicrobiomeData").
		Preload("Goals").
		Preload("MealHistory").
		Preload("RequestedMeal").
		Preload("RecentMeals").
		Preload("EnvironmentalFactors").
		Preload("LipidProfile").
		Order("id DESC").
		Find(&userHealthInfo).Error; err != nil {
		log.Println("Error fetching user data:", err)
		return nil, &CustomError{"Failed to retrieve user data", err}
	}

	return &userHealthInfo, nil
}

func (s *Service) GetUserUserId(userId int) (*models.UserHealthInfoResponse, error) {
	var userHealthInfo models.User

	// Fetch user health information with related data
	if err := s.db.Preload("BodyMetrics").
		Preload("DietaryPreferences").
		Preload("Preferences").
		Preload("HealthConditions").
		Preload("MicrobiomeData").
		Preload("MealHistory").
		Preload("RequestedMeal").
		Preload("RecentMeals").
		Preload("Goals").
		Preload("EnvironmentalFactors").
		Preload("LipidProfile").
		Where("id = ?", userId).
		First(&userHealthInfo).Error; err != nil {
		log.Println("Error fetching user health info:", err)
		return nil, &CustomError{"Failed to retrieve user health information", err}
	}

	// Fetch meals and calculate recommendations concurrently
	resultChan := make(chan []models.Recommendation, 1)
	errorChan := make(chan error, 1)

	go func() {
		// Fetch meals based on the requested meal type and number of courses
		meals, err := s.GetMealsByTypeAndCourse(userHealthInfo.RequestedMeal.StarterMealType,
			userHealthInfo.RequestedMeal.MainMealType,
			userHealthInfo.RequestedMeal.DessertMealType, userHealthInfo.RequestedMeal.NumberOfStarter, userHealthInfo.RequestedMeal.NumberOfMain, userHealthInfo.RequestedMeal.NumberOfDessert)
		if err != nil {
			errorChan <- &CustomError{"Failed to fetch meals", err}
			return
		}

		// Calculate  meals, specifying the number of courses needed
		calculateMeals := utils.CalculateMealScore(userHealthInfo, *meals)

		// Prepare recommendations based on the number of courses
		var scoreRecommendations []models.Recommendation

		// Generate course combinations based on the requested number of courses
		courseCombinations := utils.GenerateCourseCombinations(calculateMeals, userHealthInfo.RequestedMeal.NumberOfStarter, userHealthInfo.RequestedMeal.NumberOfMain, userHealthInfo.RequestedMeal.NumberOfDessert)

		// Check for valid number of course combinations
		totalCourses := userHealthInfo.RequestedMeal.NumberOfStarter + userHealthInfo.RequestedMeal.NumberOfMain + userHealthInfo.RequestedMeal.NumberOfDessert

		// Create a recommendation with the specified number of courses
		var courses []models.MealResponse

		// Add the correct number of courses for this recommendation
		for j := 0; j < totalCourses && (j) < len(courseCombinations); j++ {
			topMeal := courseCombinations[j]
			courses = append(courses, models.MealResponse{
				ID:          topMeal.Meal.ID,
				Course:      topMeal.Meal.Course,
				Name:        topMeal.Meal.Name,
				Score:       topMeal.Score,
				Ingredients: mapIngredientsToResponse(topMeal.Meal.Ingredients),
			})
		}

		// Create the recommendation and add it to the list
		recommendation := models.Recommendation{
			Combination: fmt.Sprintf("%d Meal(s)", len(courses)),
			Courses:     courses,
		}
		scoreRecommendations = append(scoreRecommendations, recommendation)

		// Send the result to the result channel
		resultChan <- scoreRecommendations
	}()

	// Wait for the result from the goroutine or handle errors
	select {
	case scoreRecommendations := <-resultChan:
		response := &models.UserHealthInfoResponse{
			User:                   userHealthInfo,
			ScoreAndRecommendation: scoreRecommendations,
		}
		return response, nil
	case err := <-errorChan:
		return nil, err
	}
}

// update
func (s *Service) PutUserUserId(userId int, userHealthInfo models.User) (*models.UserHealthInfoResponse, error) {

	userHealthInfo.ID = uint(userId)
	// Start a transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Step 1: Update users table
	userSQL := `UPDATE users SET name = ?, age = ?, gender = ?, activity_level = ?, blood_glucose = ?, health_score = ?, nutritional_deficiencies = ?, allergies = ? WHERE id = ?`
	if err := tx.Exec(userSQL, userHealthInfo.Name, userHealthInfo.Age, userHealthInfo.Gender, userHealthInfo.ActivityLevel, userHealthInfo.BloodGlucose, userHealthInfo.HealthScore, userHealthInfo.NutritionalDeficiencies, userHealthInfo.Allergies, userId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 2: Update or Insert into body_metrics table
	BodyMetricsSQL := `INSERT INTO body_metrics (user_id, weight, height)
	                   VALUES (?, ?, ?)
	                   ON CONFLICT (user_id) DO UPDATE
	                   SET weight = EXCLUDED.weight, height = EXCLUDED.height`
	if err := tx.Exec(BodyMetricsSQL, userId, userHealthInfo.BodyMetrics.Weight, userHealthInfo.BodyMetrics.Height).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 3: Update or Insert into dietary_preferences table
	DietaryPreferencesSQL := `INSERT INTO dietary_preferences (user_id, vegetarian, vegan, gluten_free, dairy_free, specific_avoidances)
	                          VALUES (?, ?, ?, ?, ?, ?)
	                          ON CONFLICT (user_id) DO UPDATE
	                          SET vegetarian = EXCLUDED.vegetarian, vegan = EXCLUDED.vegan, gluten_free = EXCLUDED.gluten_free, dairy_free = EXCLUDED.dairy_free, specific_avoidances = EXCLUDED.specific_avoidances`
	if err := tx.Exec(DietaryPreferencesSQL, userId, userHealthInfo.DietaryPreferences.Vegetarian, userHealthInfo.DietaryPreferences.Vegan, userHealthInfo.DietaryPreferences.GlutenFree, userHealthInfo.DietaryPreferences.DairyFree, userHealthInfo.DietaryPreferences.SpecificAvoidances).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 4: Update or Batch insert health_conditions table
	if len(userHealthInfo.HealthConditions) > 0 {
		// Delete old records
		if err := tx.Exec(`DELETE FROM health_conditions WHERE user_id = ?`, userId).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Insert new records
		var healthConditionsValues []interface{}
		healthConditionSQL := `INSERT INTO health_conditions (user_id, name, severity) VALUES `
		for _, condition := range userHealthInfo.HealthConditions {
			healthConditionSQL += `(?, ?, ?),`
			healthConditionsValues = append(healthConditionsValues, userId, condition.Name, condition.Severity)
		}
		healthConditionSQL = healthConditionSQL[:len(healthConditionSQL)-1]
		if err := tx.Exec(healthConditionSQL, healthConditionsValues...).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Step 5: Update or Batch insert goals table
	if len(userHealthInfo.Goals) > 0 {
		// Delete old records
		if err := tx.Exec(`DELETE FROM goals WHERE user_id = ?`, userId).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Insert new records
		var goalsValues []interface{}
		goalsSQL := `INSERT INTO goals (user_id, type, target, duration) VALUES `
		for _, goal := range userHealthInfo.Goals {
			goalsSQL += `(?, ?, ?, ?),`
			goalsValues = append(goalsValues, userId, goal.Type, goal.Target, goal.Duration)
		}
		goalsSQL = goalsSQL[:len(goalsSQL)-1]
		if err := tx.Exec(goalsSQL, goalsValues...).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Step 6: Update or Insert into microbiome_data table
	MicrobiomeDataSQL := `INSERT INTO microbiome_data (user_id, diversity_score, gut_health_recommendations)
	                      VALUES (?, ?, ?)
	                      ON CONFLICT (user_id) DO UPDATE
	                      SET diversity_score = EXCLUDED.diversity_score, gut_health_recommendations = EXCLUDED.gut_health_recommendations`
	if err := tx.Exec(MicrobiomeDataSQL, userId, userHealthInfo.MicrobiomeData.DiversityScore, userHealthInfo.MicrobiomeData.GutHealthRecommendations).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 7: Update or Insert into lipid_profiles table
	LipidProfileSQL := `INSERT INTO lipid_profiles (user_id, cholesterol, hdl, ldl, triglycerides)
	                    VALUES (?, ?, ?, ?, ?)
	                    ON CONFLICT (user_id) DO UPDATE
	                    SET cholesterol = EXCLUDED.cholesterol, hdl = EXCLUDED.hdl, ldl = EXCLUDED.ldl, triglycerides = EXCLUDED.triglycerides`
	if err := tx.Exec(LipidProfileSQL, userId, userHealthInfo.LipidProfile.Cholesterol, userHealthInfo.LipidProfile.HDL, userHealthInfo.LipidProfile.LDL, userHealthInfo.LipidProfile.Triglycerides).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 8: Update or Insert into environmental_factors table
	EnvironmentalFactorsSQL := `INSERT INTO environmental_factors (user_id, location, climate, season)
	                            VALUES (?, ?, ?, ?)
	                            ON CONFLICT (user_id) DO UPDATE
	                            SET location = EXCLUDED.location, climate = EXCLUDED.climate, season = EXCLUDED.season`
	if err := tx.Exec(EnvironmentalFactorsSQL, userId, userHealthInfo.EnvironmentalFactors.Location, userHealthInfo.EnvironmentalFactors.Climate, userHealthInfo.EnvironmentalFactors.Season).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 9: Batch update or insert meal_histories table
	if len(userHealthInfo.MealHistory) > 0 {
		if err := tx.Exec(`DELETE FROM meal_histories WHERE user_id = ?`, userId).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		var mealHistoryValues []interface{}
		MealHistorySQL := `INSERT INTO meal_histories (user_id, meal_id, timestamp) VALUES `
		for _, mealHistory := range userHealthInfo.MealHistory {
			MealHistorySQL += `(?, ?, ?),`
			mealHistoryValues = append(mealHistoryValues, userId, mealHistory.MealID, mealHistory.Timestamp)
		}
		MealHistorySQL = MealHistorySQL[:len(MealHistorySQL)-1]
		if err := tx.Exec(MealHistorySQL, mealHistoryValues...).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Step 10: Batch update or insert recent_meals table
	if len(userHealthInfo.RecentMeals) > 0 {
		if err := tx.Exec(`DELETE FROM recent_meals WHERE user_id = ?`, userId).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		var recentMealsValues []interface{}
		RecentMealsSQL := `INSERT INTO recent_meals (user_id, meal_id, timestamp) VALUES `
		for _, recentMeal := range userHealthInfo.RecentMeals {
			RecentMealsSQL += `(?, ?, ?),`
			recentMealsValues = append(recentMealsValues, userId, recentMeal.MealID, recentMeal.Timestamp)
		}
		RecentMealsSQL = RecentMealsSQL[:len(RecentMealsSQL)-1]
		if err := tx.Exec(RecentMealsSQL, recentMealsValues...).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Step 11: Update or Insert into user_preferences table
	UserPreferencesSQL := `INSERT INTO user_preferences (user_id, preferred_cuisines, meal_timings, favorite_ingredients)
	                       VALUES (?, ?, ?, ?)
	                       ON CONFLICT (user_id) DO UPDATE
	                       SET preferred_cuisines = EXCLUDED.preferred_cuisines, meal_timings = EXCLUDED.meal_timings, favorite_ingredients = EXCLUDED.favorite_ingredients`
	if err := tx.Exec(UserPreferencesSQL, userId, userHealthInfo.Preferences.PreferredCuisines, userHealthInfo.Preferences.MealTimings, userHealthInfo.Preferences.FavoriteIngredients).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	RequestedMealSQL := `INSERT INTO requested_meals (
		user_id, 
		starter_meal_type, 
		main_meal_type, 
		dessert_meal_type, 
		number_of_starter, 
		number_of_main, 
		number_of_dessert
	) VALUES (?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT (user_id) DO UPDATE
	SET 
		starter_meal_type = EXCLUDED.starter_meal_type, 
		main_meal_type = EXCLUDED.main_meal_type, 
		dessert_meal_type = EXCLUDED.dessert_meal_type, 
		number_of_starter = EXCLUDED.number_of_starter, 
		number_of_main = EXCLUDED.number_of_main, 
		number_of_dessert = EXCLUDED.number_of_dessert`
	
	if err := tx.Exec(RequestedMealSQL,
		userId,
		userHealthInfo.RequestedMeal.StarterMealType,
		userHealthInfo.RequestedMeal.MainMealType,
		userHealthInfo.RequestedMeal.DessertMealType,
		userHealthInfo.RequestedMeal.NumberOfStarter,
		userHealthInfo.RequestedMeal.NumberOfMain,
		userHealthInfo.RequestedMeal.NumberOfDessert,
	).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Fetch meals and calculate recommendations concurrently
	resultChan := make(chan []models.Recommendation, 1)
	errorChan := make(chan error, 1)

	go func() {
		// Fetch meals based on the requested meal type and number of courses
		meals, err := s.GetMealsByTypeAndCourse(userHealthInfo.RequestedMeal.StarterMealType,
			userHealthInfo.RequestedMeal.MainMealType,
			userHealthInfo.RequestedMeal.DessertMealType, userHealthInfo.RequestedMeal.NumberOfStarter, userHealthInfo.RequestedMeal.NumberOfMain, userHealthInfo.RequestedMeal.NumberOfDessert)
		if err != nil {
			errorChan <- &CustomError{"Failed to fetch meals", err}
			return
		}

		// Calculate  meals, specifying the number of courses needed
		calculateMeals := utils.CalculateMealScore(userHealthInfo, *meals)

		// Prepare recommendations based on the number of courses
		var scoreRecommendations []models.Recommendation

		// Generate course combinations based on the requested number of courses
		courseCombinations := utils.GenerateCourseCombinations(calculateMeals, userHealthInfo.RequestedMeal.NumberOfStarter, userHealthInfo.RequestedMeal.NumberOfMain, userHealthInfo.RequestedMeal.NumberOfDessert)

		// Check for valid number of course combinations
		totalCourses := userHealthInfo.RequestedMeal.NumberOfStarter + userHealthInfo.RequestedMeal.NumberOfMain + userHealthInfo.RequestedMeal.NumberOfDessert

		// Create a recommendation with the specified number of courses
		var courses []models.MealResponse

		// Add the correct number of courses for this recommendation
		for j := 0; j < totalCourses && (j) < len(courseCombinations); j++ {
			topMeal := courseCombinations[j]
			courses = append(courses, models.MealResponse{
				ID:          topMeal.Meal.ID,
				Course:      topMeal.Meal.Course,
				Name:        topMeal.Meal.Name,
				Score:       topMeal.Score,
				Ingredients: mapIngredientsToResponse(topMeal.Meal.Ingredients),
			})
		}

		// Create the recommendation and add it to the list
		recommendation := models.Recommendation{
			Combination: fmt.Sprintf("%d Meal(s)", len(courses)),
			Courses:     courses,
		}
		scoreRecommendations = append(scoreRecommendations, recommendation)

		// Send the result to the result channel
		resultChan <- scoreRecommendations
	}()

	// Wait for the result from the goroutine or handle errors
	select {
	case scoreRecommendations := <-resultChan:
		response := &models.UserHealthInfoResponse{
			User:                   userHealthInfo,
			ScoreAndRecommendation: scoreRecommendations,
		}
		return response, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (s *Service) DeleteUserUserId(userId int) error {
	if err := s.db.Where("id = ?", userId).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

// Helper function to map ingredients to response format
func mapIngredientsToResponse(ingredients []models.Ingredient) []models.IngredientResponse {
	var ingredientResponses []models.IngredientResponse
	for _, ingredient := range ingredients {
		ingredientResponses = append(ingredientResponses, models.IngredientResponse{
			Name:    ingredient.Name,
			Amount:  ingredient.Amount,
			Unit:    ingredient.Unit,
			Portion: ingredient.Portion,
			Ounces:  utils.FloatToString(utils.GramsToOunces(ingredient.Amount)),
		})
	}
	return ingredientResponses
}
