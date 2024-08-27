-- Step 1: Create the Database if it does not exist
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_catalog.pg_database WHERE datname = 'nutrition_app'
    ) THEN
        CREATE DATABASE nutrition_app;
    END IF;
END $$;

-- Connect to the nutrition_app database
\c nutrition_app;

-- Step 2: Create Tables

-- Table to store basic user health information
CREATE TABLE IF NOT EXISTS user_health_info (
    user_id SERIAL PRIMARY KEY,
    age INTEGER NOT NULL,
    gender_assigned_at_birth VARCHAR(50),
    current_gender_identity VARCHAR(50),
    hormone_use VARCHAR(100),
    height FLOAT,
    weight FLOAT,
    body_type VARCHAR(50),
    body_fat_percentage FLOAT,
    skin_tone VARCHAR(50),
    ethnic_background VARCHAR(100),
    activity_level VARCHAR(50),
    meal_timing_preferences VARCHAR(100),
    fluid_intake_water FLOAT,
    fluid_intake_other VARCHAR(255),
    seasonal_dietary_changes VARCHAR(100),
    travel_frequency VARCHAR(50)
);

-- Table to store dietary preferences and restrictions
CREATE TABLE IF NOT EXISTS dietary_preferences (
    user_id INTEGER REFERENCES user_health_info(user_id),
    dietary_preference VARCHAR(100),
    PRIMARY KEY (user_id, dietary_preference)
);

CREATE TABLE IF NOT EXISTS food_allergies_intolerances (
    user_id INTEGER REFERENCES user_health_info(user_id),
    food_allergy_intolerance VARCHAR(100),
    PRIMARY KEY (user_id, food_allergy_intolerance)
);

CREATE TABLE IF NOT EXISTS specific_food_avoidances (
    user_id INTEGER REFERENCES user_health_info(user_id),
    food_avoidance VARCHAR(100),
    PRIMARY KEY (user_id, food_avoidance)
);

CREATE TABLE IF NOT EXISTS nutritional_deficiencies (
    user_id INTEGER REFERENCES user_health_info(user_id),
    deficiency VARCHAR(100),
    PRIMARY KEY (user_id, deficiency)
);

-- Table to store frequency of consuming specific foods
CREATE TABLE IF NOT EXISTS food_consumption_frequency (
    user_id INTEGER REFERENCES user_health_info(user_id),
    food_name VARCHAR(100),
    frequency VARCHAR(50),
    PRIMARY KEY (user_id, food_name)
);

-- Table to store cooking habits and preferences
CREATE TABLE IF NOT EXISTS cooking_habits (
    user_id INTEGER REFERENCES user_health_info(user_id),
    cooking_habit VARCHAR(100),
    meal_preparation_preference VARCHAR(100),
    preferred_ingredients TEXT,
    disliked_ingredients TEXT,
    favorite_cuisines TEXT
);

-- Table to store health conditions and goals
CREATE TABLE IF NOT EXISTS health_conditions (
    user_id INTEGER REFERENCES user_health_info(user_id),
    health_condition VARCHAR(100),
    PRIMARY KEY (user_id, health_condition)
);

CREATE TABLE IF NOT EXISTS health_goals (
    user_id INTEGER REFERENCES user_health_info(user_id),
    health_goal VARCHAR(100),
    PRIMARY KEY (user_id, health_goal)
);

CREATE TABLE IF NOT EXISTS family_medical_history (
    user_id INTEGER REFERENCES user_health_info(user_id),
    family_medical_history VARCHAR(100),
    PRIMARY KEY (user_id, family_medical_history)
);

-- Table to store eating behavior and food addiction concerns
CREATE TABLE IF NOT EXISTS eating_behavior (
    user_id INTEGER REFERENCES user_health_info(user_id),
    eating_behavior VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS food_addiction_concerns (
    user_id INTEGER REFERENCES user_health_info(user_id),
    food_addiction_concern BOOLEAN DEFAULT FALSE
);

-- Table to store adverse reactions to foods
CREATE TABLE IF NOT EXISTS adverse_reactions_to_foods (
    user_id INTEGER REFERENCES user_health_info(user_id),
    adverse_reaction VARCHAR(255)
);

-- Table to store health monitoring and technology usage
CREATE TABLE IF NOT EXISTS health_monitoring_technology (
    user_id INTEGER REFERENCES user_health_info(user_id),
    health_monitoring_device VARCHAR(100),
    PRIMARY KEY (user_id, health_monitoring_device)
);

-- Table to store historical eating patterns and meal logs
CREATE TABLE IF NOT EXISTS historical_eating_patterns (
    user_id INTEGER REFERENCES user_health_info(user_id),
    meal_time VARCHAR(50),
    food_items TEXT,
    PRIMARY KEY (user_id, meal_time)
);

CREATE TABLE IF NOT EXISTS eating_out_frequency (
    user_id INTEGER REFERENCES user_health_info(user_id),
    frequency VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS snacking_habits (
    user_id INTEGER REFERENCES user_health_info(user_id),
    snacking_habit VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS alcohol_consumption (
    user_id INTEGER REFERENCES user_health_info(user_id),
    alcohol_consumption VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS smoking_status (
    user_id INTEGER REFERENCES user_health_info(user_id),
    smoking_status VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS recent_diet_changes (
    user_id INTEGER REFERENCES user_health_info(user_id),
    recent_diet_change BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS feedback_on_past_diets (
    user_id INTEGER REFERENCES user_health_info(user_id),
    feedback TEXT
);

CREATE TABLE IF NOT EXISTS eating_schedule_flexibility (
    user_id INTEGER REFERENCES user_health_info(user_id),
    eating_schedule_flexibility VARCHAR(100)
);

-- Table to store health test results
CREATE TABLE IF NOT EXISTS microbiome_data (
    user_id INTEGER REFERENCES user_health_info(user_id),
    microbiome_diversity_score FLOAT,
    specific_bacteria_levels FLOAT[]
);

CREATE TABLE IF NOT EXISTS cgm_data (
    user_id INTEGER REFERENCES user_health_info(user_id),
    average_blood_glucose_levels FLOAT,
    glucose_variability FLOAT
);

CREATE TABLE IF NOT EXISTS lipid_profile_data (
    user_id INTEGER REFERENCES user_health_info(user_id),
    total_cholesterol FLOAT,
    hdl FLOAT,
    ldl FLOAT,
    triglycerides FLOAT
);

CREATE TABLE IF NOT EXISTS other_health_tests (
    user_id INTEGER REFERENCES user_health_info(user_id),
    test_name VARCHAR(100),
    test_value FLOAT,
    test_units VARCHAR(50),
    PRIMARY KEY (user_id, test_name)
);

-- Table to store environmental factors
CREATE TABLE IF NOT EXISTS environmental_factors (
    user_id INTEGER REFERENCES user_health_info(user_id),
    current_location VARCHAR(255),
    temperature FLOAT,
    humidity FLOAT,
    weather_condition VARCHAR(50)
);

-- Step 3: Establish Relationships
-- Relationships are defined through foreign keys in table definitions where appropriate.

-- All tables reference the primary `user_health_info` table through `user_id`.
