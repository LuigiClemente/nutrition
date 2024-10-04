-- Sample Users
INSERT INTO users (name, age, gender, activity_level, blood_glucose, health_score, nutritional_deficiencies, allergies) VALUES
('Alice Johnson', 34, 'Female', 'Moderately active', 5.4, 75.3, '{"Vitamin D"}', '{"Peanuts"}'),
('Bob Smith', 45, 'Male', 'Sedentary', 6.1, 68.9, '{"Iron", "Vitamin B12"}', '{}'),
('Catherine Green', 29, 'Female', 'Very active', 4.8, 82.5, '{}', '{"Shellfish"}'),
('David Thompson', 52, 'Male', 'Lightly active', 7.2, 64.2, '{"Calcium"}', '{"Gluten"}'),
('Evelyn Parker', 38, 'Female', 'Moderately active', 5.9, 79.1, '{"Magnesium"}', '{"Soy"}'),
('Frank Baker', 27, 'Male', 'Active', 4.5, 85.7, '{}', '{}'),
('Grace White', 43, 'Female', 'Sedentary', 6.3, 70.8, '{"Zinc", "Vitamin A"}', '{}'),
('Henry Turner', 31, 'Male', 'Lightly active', 5.0, 76.2, '{}', '{"Lactose"}'),
('Ivy Wilson', 24, 'Female', 'Active', 4.7, 88.4, '{"Iron"}', '{"Eggs"}'),
('Jack Lewis', 50, 'Male', 'Very active', 6.0, 81.3, '{}', '{"Nuts"}');

-- Body Metrics for Users
INSERT INTO body_metrics (user_id, weight, height) VALUES
(1, 68.5, 165),
(2, 82.1, 178),
(3, 59.7, 160),
(4, 90.3, 172),
(5, 65.2, 168),
(6, 74.4, 180),
(7, 72.3, 164),
(8, 85.0, 175),
(9, 58.9, 162),
(10, 88.5, 182);

-- Dietary Preferences for Users
INSERT INTO dietary_preferences (user_id, vegetarian, vegan, gluten_free, dairy_free, specific_avoidances) VALUES
(1, TRUE, FALSE, FALSE, TRUE, '{"Soy", "Corn"}'),
(2, FALSE, FALSE, FALSE, FALSE, '{}'),
(3, FALSE, FALSE, TRUE, FALSE, '{"Wheat"}'),
(4, FALSE, FALSE, TRUE, TRUE, '{"Dairy", "Eggs"}'),
(5, TRUE, TRUE, FALSE, TRUE, '{"Meat"}'),
(6, FALSE, FALSE, FALSE, FALSE, '{}'),
(7, FALSE, FALSE, FALSE, FALSE, '{}'),
(8, FALSE, FALSE, TRUE, FALSE, '{"Milk"}'),
(9, TRUE, FALSE, FALSE, TRUE, '{"Chicken"}'),
(10, FALSE, FALSE, FALSE, FALSE, '{"Nuts"}');

-- Health Conditions for Users
INSERT INTO health_conditions (user_id, name, severity) VALUES
(1, 'Hypertension', 'Moderate'),
(2, 'Diabetes', 'Severe'),
(3, 'Asthma', 'Mild'),
(4, 'High Cholesterol', 'Moderate'),
(5, 'Celiac Disease', 'Severe'),
(6, 'No major conditions', 'N/A'),
(7, 'Osteoporosis', 'Moderate'),
(8, 'Hypertension', 'Mild'),
(9, 'Lactose Intolerance', 'Mild'),
(10, 'Peanut Allergy', 'Severe');

-- Requested Meals for Users
INSERT INTO requested_meals (user_id, meal_type, number_of_starter, number_of_main, number_of_dessert) VALUES
(1, 2, 1, 1, 2),
(2, 1, 1, 1, 2),
(3, 3, 1, 1, 2),
(4, 2, 1, 1, 2),
(5, 1, 1, 1, 2),
(6, 2, 1, 1, 2),
(7, 3, 1, 1, 2),
(8, 1, 1, 1, 2),
(9, 2, 1, 1, 2),
(10, 3, 1, 2, 1);

-- Goals for Users
INSERT INTO goals (user_id, type, target, duration) VALUES
(1, 'Weight Loss', 65.0, 90),
(2, 'Blood Sugar Management', 5.5, 180),
(3, 'Improve Stamina', 60.0, 120),
(4, 'Reduce Cholesterol', 80.0, 100),
(5, 'Weight Maintenance', 65.0, 365),
(6, 'Muscle Gain', 80.0, 180),
(7, 'Bone Health', 70.0, 120),
(8, 'Reduce Blood Pressure', 70.0, 90),
(9, 'Gut Health', 55.0, 60),
(10, 'Allergy Management', 85.0, 180);

-- Microbiome Data for Users
INSERT INTO microbiome_data (user_id, diversity_score, gut_health_recommendations) VALUES
(1, 75.5, '{"Increase Fiber", "Reduce Fat Intake"}'),
(2, 60.3, '{"Add Probiotics", "Eat More Vegetables"}'),
(3, 85.0, '{"Increase Water Intake", "Reduce Sugar"}'),
(4, 70.2, '{"Add Prebiotics", "Reduce Meat Consumption"}'),
(5, 68.4, '{"Avoid Gluten", "Eat More Fiber"}'),
(6, 90.1, '{"Increase Protein", "Eat Fermented Foods"}'),
(7, 65.7, '{"Increase Calcium", "Eat Leafy Greens"}'),
(8, 78.3, '{"Reduce Salt", "Increase Fruits"}'),
(9, 82.6, '{"Add Whole Grains", "Eat More Yogurt"}'),
(10, 69.8, '{"Avoid Allergens", "Increase Fiber"}');

-- Environmental Factors for Users
INSERT INTO environmental_factors (user_id, location, climate, season) VALUES
(1, 'New York, USA', 'Temperate', 'Winter'),
(2, 'Los Angeles, USA', 'Mediterranean', 'Summer'),
(3, 'Chicago, USA', 'Continental', 'Spring'),
(4, 'Dallas, USA', 'Subtropical', 'Autumn'),
(5, 'Miami, USA', 'Tropical', 'Summer'),
(6, 'Denver, USA', 'Arid', 'Winter'),
(7, 'Boston, USA', 'Temperate', 'Spring'),
(8, 'Seattle, USA', 'Marine', 'Autumn'),
(9, 'San Francisco, USA', 'Mediterranean', 'Summer'),
(10, 'Phoenix, USA', 'Desert', 'Winter');


-- Insert Meal Tags
INSERT INTO meal_tags (tag) VALUES
('Healthy'),
('Vegan'),
('Gluten-Free'),
('Spicy'),
('Dessert'),
('High Protein'),
('Low Carb'),
('Dairy-Free'),
('Low Calorie'),
('Keto');


-- Insert Meal Types
INSERT INTO meal_types (type) VALUES
('Salad'),
('Sandwich'),
('Burger'),
('Chicken'),
('Meat'),
('Pasta'),
('Soup'),
('Pizza'),
('Seafood'),
('Vegetarian');


-- Insert Meals
INSERT INTO meals (name, nutritional_content, meal_timings, meal_type_id, course, cuisine) VALUES
('Grilled Chicken Salad', '{"Calories": 350, "Protein": 30, "Carbs": 10, "Fats": 15}', '{"Lunch"}', 1, 'Main', 'American'),
('Vegetable Stir Fry', '{"Calories": 200, "Protein": 8, "Carbs": 35, "Fats": 5}', '{"Dinner"}', 2, 'Main', 'Asian'),
('Quinoa Salad', '{"Calories": 250, "Protein": 12, "Carbs": 40, "Fats": 10}', '{"Lunch", "Dinner"}', 1, 'Main', 'Mediterranean'),
('Gluten-Free Pancakes', '{"Calories": 300, "Protein": 5, "Carbs": 50, "Fats": 8}', '{"Breakfast"}', 3, 'Main', 'American'),
('Baked Salmon', '{"Calories": 400, "Protein": 35, "Carbs": 5, "Fats": 20}', '{"Dinner"}', 2, 'Main', 'European'),
('Vegan Buddha Bowl', '{"Calories": 350, "Protein": 15, "Carbs": 50, "Fats": 10}', '{"Lunch"}', 1, 'Main', 'Vegan'),
('Chicken Curry', '{"Calories": 500, "Protein": 35, "Carbs": 60, "Fats": 15}', '{"Dinner"}', 2, 'Main', 'Indian'),
('Fruit Parfait', '{"Calories": 150, "Protein": 3, "Carbs": 30, "Fats": 4}', '{"Breakfast", "Dessert"}', 3, 'Dessert', 'French'),
('Lentil Soup', '{"Calories": 180, "Protein": 12, "Carbs": 30, "Fats": 3}', '{"Lunch", "Dinner"}', 1, 'Main', 'Middle Eastern'),
('Tofu Scramble', '{"Calories": 220, "Protein": 18, "Carbs": 8, "Fats": 12}', '{"Breakfast"}', 3, 'Main', 'Vegan');

-- Insert Ingredients for Meals
INSERT INTO ingredients (meal_id, name, amount, unit, portion) VALUES
(1, 'Chicken Breast', 150, 'Grams', '1 pcs'),
(1, 'Mixed Greens', 100, 'Grams', '2 cups'),
(1, 'Olive Oil', 10, 'ML', '1 tbsp'),
(2, 'Broccoli', 150, 'Grams', '1.5 cups'),
(2, 'Bell Peppers', 100, 'Grams', '1 pcs'),
(2, 'Soy Sauce', 20, 'ML', '1.5 tbsp'),
(3, 'Quinoa', 100, 'Grams', '0.5 cup'),
(3, 'Cucumber', 50, 'Grams', '1 pcs'),
(3, 'Olive Oil', 10, 'ML', '1 tbsp'),
(4, 'Gluten-Free Flour', 100, 'Grams', '1 cup'),
(4, 'Almond Milk', 200, 'ML', '1 cup'),
(4, 'Maple Syrup', 20, 'ML', '1.5 tbsp'),
(5, 'Salmon Fillet', 200, 'Grams', '1 pcs'),
(5, 'Lemon', 1, 'Piece', '1 pcs'),
(5, 'Olive Oil', 10, 'ML', '1 tbsp'),
(6, 'Brown Rice', 100, 'Grams', '0.5 cup'),
(6, 'Chickpeas', 50, 'Grams', '0.25 cup'),
(6, 'Tahini', 20, 'Grams', '1 tbsp'),
(7, 'Chicken Thigh', 150, 'Grams', '1 pcs'),
(7, 'Curry Paste', 50, 'Grams', '2 tbsp'),
(7, 'Coconut Milk', 200, 'ML', '0.75 cup'),
(8, 'Greek Yogurt', 100, 'Grams', '0.5 cup'),
(8, 'Mixed Berries', 50, 'Grams', '0.25 cup'),
(8, 'Granola', 30, 'Grams', '0.25 cup'),
(9, 'Lentils', 100, 'Grams', '0.5 cup'),
(9, 'Carrots', 50, 'Grams', '1 pcs'),
(9, 'Olive Oil', 10, 'ML', '1 tbsp'),
(10, 'Tofu', 150, 'Grams', '1 pcs'),
(10, 'Turmeric', 5, 'Grams', '1 tsp'),
(10, 'Spinach', 100, 'Grams', '1 cup');


-- Insert Meal Tags for Meals
INSERT INTO meal_tag_relationships (meal_id, tag_id) VALUES
(1, 1),  -- Grilled Chicken Salad tagged as 'Healthy'
(2, 2),  -- Vegetable Stir Fry tagged as 'Vegan'
(3, 1),  -- Quinoa Salad tagged as 'Gluten-Free'
(4, 3),  -- Gluten-Free Pancakes tagged as 'Gluten-Free'
(5, 1),  -- Baked Salmon tagged as 'High Protein'
(6, 2),  -- Vegan Buddha Bowl tagged as 'Vegan'
(7, 4),  -- Chicken Curry tagged as 'Spicy'
(8, 5),  -- Fruit Parfait tagged as 'Dessert'
(9, 1),  -- Lentil Soup tagged as 'Healthy'
(10, 2); -- Tofu Scramble tagged as 'Vegan'

-- Insert Meal History for Users
INSERT INTO meal_histories (user_id, meal_id, timestamp) VALUES
(1, 1, NOW()),
(2, 2, NOW()),
(3, 3, NOW()),
(4, 4, NOW()),
(5, 5, NOW()),
(6, 6, NOW()),
(7, 7, NOW()),
(8, 8, NOW()),
(9, 9, NOW()),
(10, 10, NOW());

-- Insert Recent Meals for Users
INSERT INTO recent_meals (user_id, meal_id, timestamp) VALUES
(1, 2, NOW()),
(2, 3, NOW()),
(3, 4, NOW()),
(4, 5, NOW()),
(5, 6, NOW()),
(6, 7, NOW()),
(7, 8, NOW()),
(8, 9, NOW()),
(9, 10, NOW()),
(10, 1, NOW());




-- Insert Meals with Course Type as 'Starter'
INSERT INTO meals (name, nutritional_content, meal_timings, meal_type_id, course, cuisine) VALUES
('Bruschetta', '{"Calories": 120, "Protein": 5, "Fat": 3, "Carbohydrates": 20}', '{}', 1, 'Starter', 'Italian'),
('Stuffed Mushrooms', '{"Calories": 150, "Protein": 6, "Fat": 5, "Carbohydrates": 18}', '{}', 1, 'Starter', 'Italian');

-- Insert Ingredients for Meals
INSERT INTO ingredients (meal_id, name, amount, unit, portion) VALUES
(11, 'Tomato', 100, 'grams', '2 pcs'),  -- Bruschetta
(11, 'Basil', 10, 'grams', '1'),     -- Bruschetta
(12, 'Mushrooms', 150, 'grams', '2 pcs'), -- Stuffed Mushrooms
(12, 'Cream Cheese', 50, 'grams', '2 tbsp'); -- Stuffed Mushrooms


-- Assuming you have meal tags already inserted, let's tag meals
-- First, insert the relationships
INSERT INTO meal_tag_relationships (meal_id, tag_id) VALUES
(11, 1),  -- Bruschetta, Healthy
(12, 1);  -- Stuffed Mushrooms, Healthy



-- Insert into meal_types for Dessert
INSERT INTO meal_types (type) VALUES ('Dessert');

-- Insert some dessert meals into the meals table
INSERT INTO meals (name, nutritional_content, meal_timings, meal_type_id, course, cuisine)
VALUES 
    ('Chocolate Mousse', '{"Calories": 250, "Protein": 4, "Fat": 20, "Carbohydrates": 22}', 
        ARRAY['Evening'], 1, 'Dessert', 'French'),
    
    ('Fruit Salad', '{"Calories": 120, "Protein": 2, "Fat": 0.5, "Carbohydrates": 30}', 
        ARRAY['Afternoon'], 1, 'Dessert', 'Global'),
    
    ('Cheesecake', '{"Calories": 300, "Protein": 5, "Fat": 22, "Carbohydrates": 25}', 
        ARRAY['Evening'], 1, 'Dessert', 'American'),

    ('Apple Pie', '{"Calories": 240, "Protein": 3, "Fat": 10, "Carbohydrates": 36}', 
        ARRAY['Evening'], 1, 'Dessert', 'American'),

    ('Tiramisu', '{"Calories": 320, "Protein": 6, "Fat": 15, "Carbohydrates": 40}', 
        ARRAY['Evening'], 1, 'Dessert', 'Italian');

-- Optionally, you can also insert ingredient details for the desserts
INSERT INTO ingredients (meal_id, name, amount, unit, portion)
VALUES 
    (13, 'Dark Chocolate', 200, 'grams', '4 pcs'),
    (13, 'Eggs', 3, 'pieces', '3 pcs'),
    (14, 'Mixed Fruits', 300, 'grams', '1/2'),
    (15, 'Cream Cheese', 250, 'grams', '2 pcs'),
    (16, 'Apples', 4, 'pieces', 'Sliced'),
    (17, 'Mascarpone Cheese', 250, 'grams', '1 pcs'),
    (17, 'Ladyfingers', 200, 'grams', '2');

-- Insert meal tag relationships for the dessert meals
INSERT INTO meal_tag_relationships (meal_id, tag_id)
VALUES 
    (13, 5),  -- Chocolate Mousse - Dessert
    (13, 1),  -- Chocolate Mousse - Healthy (optional)
    (14, 5),  -- Fruit Salad - Dessert
    (14, 1),  -- Fruit Salad - Healthy
    (14, 2),  -- Fruit Salad - Vegan
    (15, 5),  -- Cheesecake - Dessert
    (15, 1),  -- Cheesecake - Healthy (optional)
    (15, 8),  -- Cheesecake - Low Calorie (optional)
    (16, 5),  -- Apple Pie - Dessert
    (16, 1),  -- Apple Pie - Healthy (optional)
    (17, 5),  -- Tiramisu - Dessert
    (17, 1),  -- Tiramisu - Healthy (optional)
    (17, 7);  -- Tiramisu - Low Carb
