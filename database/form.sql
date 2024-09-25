-- User Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(32),
    activity_level VARCHAR(50),
    blood_glucose FLOAT,
    health_score FLOAT,
    nutritional_deficiencies TEXT[],
    allergies TEXT[]
);

-- Body Metrics Table
CREATE TABLE body_metrics (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    weight FLOAT,
    height FLOAT,
    PRIMARY KEY(user_id)
);

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

-- Health Conditions Table
CREATE TABLE health_conditions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(128),
    severity VARCHAR(50)
);



-- Last Requested  meal Category Table
CREATE TABLE requested_meals (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    meal_category  VARCHAR(64),
    number_of_courses INT,
    timestamp TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY(user_id)
);

-- Goals Table
CREATE TABLE goals (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(128),
    target FLOAT,
    duration INT
);

-- Microbiome Data Table
CREATE TABLE microbiome_data (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    diversity_score FLOAT,
    gut_health_recommendations TEXT[],
    PRIMARY KEY(user_id)
);

-- Lipid Profile Table
CREATE TABLE lipid_profiles (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    cholesterol FLOAT,
    hdl FLOAT,
    ldl FLOAT,
    triglycerides FLOAT,
    PRIMARY KEY(user_id)
);

-- Environmental Factors Table
CREATE TABLE environmental_factors (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    location VARCHAR(128),
    climate VARCHAR(64),
    season VARCHAR(50),
    PRIMARY KEY(user_id)
    
);

-- Meals Table
CREATE TABLE meals (
    id SERIAL PRIMARY KEY,  -- Auto-incrementing ID for the meal
    name VARCHAR(128) NOT NULL,  -- Name of the meal
    nutritional_content JSONB,  -- Nutritional content stored as JSONB
    category VARCHAR(64),  -- Category (e.g., 'Salad')
    meal_type TEXT[],  -- Meal type as a PostgreSQL array (e.g., ['Lunch', 'Dinner'])
    cuisine VARCHAR(64),  -- Cuisine type (e.g., 'American')
    tags TEXT[]  -- Tags as a PostgreSQL array (e.g., ['Healthy', 'Low-Carb'])
);


CREATE INDEX idx_category ON meals (category);


-- Meal History Table
CREATE TABLE meal_histories (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,
    timestamp TIMESTAMP
);

-- Meal History Table
CREATE TABLE recent_meals (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,
    timestamp TIMESTAMP
);

-- Ingredients Table
CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,  -- Auto-incrementing ID for the ingredient
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,  -- Foreign key linking to the meal
    name VARCHAR(128) NOT NULL,  -- Name of the ingredient
    amount FLOAT,  -- Amount of the ingredient (e.g., 150) in grams
    portion VARCHAR(64),  -- Total portion size calculated
    portion_unit VARCHAR(64)  -- tbsp
);


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
    score FLOAT
);
