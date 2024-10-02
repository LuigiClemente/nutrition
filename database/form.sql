CREATE DATABASE nutrition_app;


-- connect to database and create the tables


-- User Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    age INT NOT NULL CHECK (age > 0),  -- Ensuring positive age
    gender VARCHAR(6) CHECK (gender IN ('Male', 'Female', 'Other')),  -- Enum-like constraint
    activity_level VARCHAR(50),  -- Could consider ENUM for predefined levels
    blood_glucose FLOAT CHECK (blood_glucose >= 0),  -- Non-negative check
    health_score FLOAT CHECK (health_score >= 0),  -- Non-negative health score
    nutritional_deficiencies TEXT[] DEFAULT '{}',  -- Empty array by default
    allergies TEXT[] DEFAULT '{}'  -- Default empty array
);

-- Adding Index on Foreign Keys for Optimization
CREATE INDEX idx_users_name ON users(name);

-- Optimized Body Metrics Table
CREATE TABLE body_metrics (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    weight FLOAT,
    height FLOAT,
    PRIMARY KEY(user_id)
);

-- Index for user_id for faster lookups
CREATE INDEX idx_body_metrics_user_id ON body_metrics(user_id);

-- Dietary Preferences Table
CREATE TABLE dietary_preferences (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    vegetarian BOOLEAN,
    vegan BOOLEAN,
    gluten_free BOOLEAN,
    dairy_free BOOLEAN,
    specific_avoidances TEXT[],  -- Array of avoidances
    PRIMARY KEY(user_id)
);

-- Index for user_id for faster lookups
CREATE INDEX idx_dietary_preferences_user_id ON dietary_preferences(user_id);

-- Health Conditions Table
CREATE TABLE health_conditions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(128),
    severity VARCHAR(50)  -- Consider ENUM for predefined severity levels
);

-- Index on user_id for optimization
CREATE INDEX idx_health_conditions_user_id ON health_conditions(user_id);

-- Last Requested Meal Category Table
CREATE TABLE requested_meals (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    meal_type INT,  
    number_of_course INT,  
    timestamp TIMESTAMP DEFAULT NOW(),  -- Automatically logs the request time
    PRIMARY KEY(user_id)
);

-- Index for user_id and timestamp to optimize filtering
CREATE INDEX idx_requested_meals_user_id_timestamp ON requested_meals(user_id, timestamp);

-- Goals Table
CREATE TABLE goals (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(128),
    target DECIMAL(10, 2),  -- Using DECIMAL for precise float values
    duration INT  -- Duration in days or months (you can specify)
);

-- Index on user_id to speed up querying user goals
CREATE INDEX idx_goals_user_id ON goals(user_id);

-- Microbiome Data Table
CREATE TABLE microbiome_data (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    diversity_score DECIMAL(10, 2),  -- Using DECIMAL for precise scores
    gut_health_recommendations TEXT[],  -- Array of recommendations
    PRIMARY KEY(user_id)
);

-- Index on user_id for faster lookups
CREATE INDEX idx_microbiome_data_user_id ON microbiome_data(user_id);

-- Environmental Factors Table
CREATE TABLE environmental_factors (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    location VARCHAR(128),
    climate VARCHAR(64),  -- Consider ENUM for predefined climates
    season VARCHAR(50),  -- Consider ENUM for predefined seasons
    PRIMARY KEY(user_id)
);

-- Index on user_id to speed up querying by environmental factors
CREATE INDEX idx_environmental_factors_user_id ON environmental_factors(user_id);

-- Suggested Meal Types Table
CREATE TABLE meal_types (
    id SERIAL PRIMARY KEY,
    type VARCHAR(64) UNIQUE NOT NULL
);

-- Suggested Tags Table
CREATE TABLE meal_tags (
    id SERIAL PRIMARY KEY,
    tag VARCHAR(64) UNIQUE NOT NULL
);

-- Meals Table with Foreign Key Relationships
CREATE TABLE meals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    nutritional_content JSONB,  -- JSONB for flexible nutritional data
    meal_timings TEXT[] DEFAULT '{}',  -- Default empty array
    meal_type_id INT REFERENCES meal_types(id) ON DELETE SET NULL,  -- Foreign key to meal types
    course VARCHAR(7) CHECK (course IN ('Starter', 'Main', 'Dessert')),  -- Enum-like constraint,
    cuisine VARCHAR(64)
);

-- Index on meal_type_id for faster lookups
CREATE INDEX idx_meals_meal_type_id ON meals(meal_type_id);

CREATE INDEX idx_meals_course ON meals(course);

-- Separate Table for Meal Tag Relationships
CREATE TABLE meal_tag_relationships (
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,
    tag_id INT REFERENCES meal_tags(id) ON DELETE CASCADE,
    PRIMARY KEY(meal_id, tag_id)
);

-- Meal History Table
CREATE TABLE meal_histories (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,
    timestamp TIMESTAMP DEFAULT NOW()
);

-- Meal History Table
CREATE TABLE recent_meals (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,
    timestamp TIMESTAMP DEFAULT NOW()
);

-- Index for user_id and meal_id for faster lookups
CREATE INDEX idx_meal_logs_user_id_meal_id ON meal_logs(user_id, meal_id);

-- Ingredients Table
CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,  -- Auto-incrementing ID for the ingredient
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,  -- Foreign key linking to the meal
    name VARCHAR(128) NOT NULL,  -- Name of the ingredient
    amount DECIMAL(10, 2),  -- Amount of the ingredient (e.g., 150) in grams, using DECIMAL for precision
    portion VARCHAR(64),  -- Portion description
);

-- Index on meal_id for faster lookups
CREATE INDEX idx_ingredients_meal_id ON ingredients(meal_id);

-- User Preferences Table
CREATE TABLE user_preferences (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    preferred_cuisines TEXT[],    -- Array of preferred cuisines
    meal_timings TEXT[],          -- Array of meal timings
    favorite_ingredients TEXT[],   -- Array of favorite ingredients
    PRIMARY KEY(user_id)
);

-- Scored Meals Table
CREATE TABLE scored_meals (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,
    score FLOAT  -- Meal score
);

-- Index for user_id for optimization
CREATE INDEX idx_scored_meals_user_id ON scored_meals(user_id);


-- Lipid Profile Table
CREATE TABLE lipid_profiles (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    cholesterol FLOAT,
    hdl FLOAT,
    ldl FLOAT,
    triglycerides FLOAT,
    PRIMARY KEY(user_id)
);