-- Sample Users
INSERT INTO users (name, age, gender, activity_level, blood_glucose, health_score, nutritional_deficiencies, allergies) VALUES
('Alice Johnson', 34, 'Female', 'Moderately Active', 5.4, 75.3, '{"Vitamin D"}', '{"Peanuts"}'),
('Bob Smith', 45, 'Male', 'Sedentary', 6.1, 68.9, '{"Iron", "Vitamin B12"}', '{"Shellfish"}'),
('Catherine Green', 29, 'Female', 'Very Active', 4.8, 82.5, '{}', '{"Shellfish"}'),
('David Thompson', 52, 'Male', 'Lightly Active', 7.2, 64.2, '{"Calcium"}', '{"Gluten"}'),
('Evelyn Parker', 38, 'Female', 'Moderately Active', 5.9, 79.1, '{"Magnesium"}', '{"Soy"}'),
('Frank Baker', 27, 'Male', 'Active', 4.5, 85.7, '{}', '{}'),
('Grace White', 43, 'Female', 'Sedentary', 6.3, 70.8, '{"Zinc", "Vitamin A"}', '{}'),
('Henry Turner', 31, 'Male', 'Lightly Active', 5.0, 76.2, '{"Vitamin B12", "Vitamin B6"}', '{"Lactose"}'),
('Ivy Wilson', 24, 'Female', 'Active', 4.7, 88.4, '{"Iron"}', '{"Eggs"}'),
('Jack Lewis', 50, 'Male', 'Very Active', 6.0, 81.3, '{"Vitamin D"}', '{"Nuts"}');

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
(2, FALSE, FALSE, FALSE, FALSE, '{Beans}'),
(3, FALSE, FALSE, TRUE, FALSE, '{"Wheat"}'),
(4, FALSE, FALSE, TRUE, TRUE, '{"Dairy", "Eggs"}'),
(5, TRUE, TRUE, FALSE, TRUE, '{"Meat"}'),
(6, FALSE, FALSE, FALSE, FALSE, '{Beans}'),
(7, FALSE, FALSE, FALSE, FALSE, '{Beans}'),
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
(9, 'Gut health', 55.0, 60),
(10, 'Improved Mental Clarity', 85.0, 180);

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
(1, 'USA', 'Temperate', 'Winter'),
(2, 'USA', 'Mediterranean', 'Summer'),
(3, 'USA', 'Continental', 'Spring'),
(4, 'USA', 'Subtropical', 'Autumn'),
(5, 'USA', 'Tropical', 'Summer'),
(6, 'USA', 'Arid', 'Winter'),
(7, 'USA', 'Temperate', 'Spring'),
(8, 'USA', 'Marine', 'Autumn'),
(9, 'USA', 'Mediterranean', 'Summer'),
(10, 'UK', 'Desert', 'Winter');

INSERT INTO user_preferences (user_id, preferred_cuisines, meal_timings, favorite_ingredients) VALUES
(1, '{"Italian", "Mexican"}', '{"Breakfast", "Lunch"}', '{"Tomato", "Basil", "Cheese"}'),
(2, '{"Chinese", "Indian"}', '{"Dinner", "Snack"}', '{"Ginger", "Garlic", "Rice"}'),
(3, '{"Mediterranean", "Vegetarian"}', '{"Lunch", "Dinner"}', '{"Olive Oil", "Lentils", "Feta"}'),
(4, '{"American", "Barbecue"}', '{"Breakfast", "Dinner"}', '{"Beef", "Corn", "BBQ Sauce"}'),
(5, '{"Japanese", "Sushi"}', '{"Lunch", "Dinner"}', '{"Fish", "Avocado", "Seaweed"}'),
(6, '{"Korean", "Fusion"}', '{"Dinner", "Snack"}', '{"Kimchi", "Tofu", "Noodles"}'),
(7, '{"Thai", "Spicy"}', '{"Lunch", "Dinner"}', '{"Chili", "Coconut Milk", "Lime"}'),
(8, '{"French", "Pastry"}', '{"Breakfast", "Dessert"}', '{"Butter", "Flour", "Chocolate"}'),
(9, '{"Indian", "Vegan"}', '{"Lunch", "Dinner"}', '{"Chickpeas", "Spinach", "Curry Powder"}'),
(10, '{"Mediterranean", "Healthy"}', '{"Breakfast", "Lunch", "Dinner"}', '{"Quinoa", "Kale", "Avocado"}');

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
('Keto'),
('Vegetarian'),
('Paleo'),
('Whole30'),
('Sugar-Free'),
('Low Fat'),
('Organic'),
('High Fiber'),
('Nut-Free'),
('Soy-Free'),
('Comfort Food'),
('Quick & Easy'),
('Family-Friendly'),
('Mediterranean'),
('Diabetic-Friendly'),
('Heart-Healthy'),
('Anti-Inflammatory'),
('Meal Prep'),
('Low Sodium'),
('Whole Grain'),
('Plant-Based');


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
('Grilled Chicken Salad', '{"Calories": 350, "Protein": 30, "Carbs": 10, "Fat": 15}', '{"Lunch"}', 1, 'Main', 'American'),
('Vegetable Stir Fry', '{"Calories": 200, "Protein": 8, "Carbs": 35, "Fat": 5}', '{"Dinner"}', 2, 'Main', 'Asian'),
('Quinoa Salad', '{"Calories": 250, "Protein": 12, "Carbs": 40, "Fat": 10}', '{"Lunch", "Dinner"}', 1, 'Main', 'Mediterranean'),
('Gluten-Free Pancakes', '{"Calories": 300, "Protein": 5, "Carbs": 50, "Fat": 8}', '{"Breakfast"}', 3, 'Main', 'American'),
('Baked Salmon', '{"Calories": 400, "Protein": 35, "Carbs": 5, "Fat": 20}', '{"Dinner"}', 2, 'Main', 'European'),
('Vegan Buddha Bowl', '{"Calories": 350, "Protein": 15, "Carbs": 50, "Fat": 10}', '{"Lunch"}', 1, 'Main', 'Vegan'),
('Chicken Curry', '{"Calories": 500, "Protein": 35, "Carbs": 60, "Fat": 15}', '{"Dinner"}', 2, 'Main', 'Indian'),
('Fruit Parfait', '{"Calories": 150, "Protein": 3, "Carbs": 30, "Fat": 4}', '{"Breakfast", "Dessert"}', 3, 'Dessert', 'French'),
('Lentil Soup', '{"Calories": 180, "Protein": 12, "Carbs": 30, "Fat": 3}', '{"Lunch", "Dinner"}', 1, 'Main', 'Middle Eastern'),
('Tofu Scramble', '{"Calories": 220, "Protein": 18, "Carbs": 8, "Fat": 12}', '{"Breakfast"}', 3, 'Main', 'Vegan'),
('Bruschetta', '{"Calories": 120, "Protein": 5, "Fat": 3, "Carbohydrates": 20}', '{}', 1, 'Starter', 'Italian'),
('Stuffed Mushrooms', '{"Calories": 150, "Protein": 6, "Fat": 5, "Carbohydrates": 18}', '{}', 1, 'Starter', 'Italian');

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
(10, 'Spinach', 100, 'Grams', '1 cup'),
(11, 'Tomato', 100, 'grams', '2 pcs'),  -- Bruschetta
(11, 'Basil', 10, 'grams', '1'),     -- Bruschetta
(12, 'Mushrooms', 150, 'grams', '2 pcs'), -- Stuffed Mushrooms
(12, 'Cream Cheese', 50, 'grams', '2 tbsp'); -- Stuffed Mushrooms;


-- Insert Meal Tags for Meals
INSERT INTO meal_tag_relationships (meal_id, tag_id) VALUES
(1, 1),   -- Grilled Chicken Salad tagged as 'Healthy'
(1, 6),   -- Grilled Chicken Salad tagged as 'High Protein'
(1, 7),   -- Grilled Chicken Salad tagged as 'Low Carb'
(2, 2),   -- Vegetable Stir Fry tagged as 'Vegan'
(2, 9),   -- Vegetable Stir Fry tagged as 'Low Calorie'
(3, 1),   -- Quinoa Salad tagged as 'Healthy'
(3, 3),   -- Quinoa Salad tagged as 'Gluten-Free'
(3, 16),  -- Quinoa Salad tagged as 'Organic'
(4, 3),   -- Gluten-Free Pancakes tagged as 'Gluten-Free'
(4, 5),   -- Gluten-Free Pancakes tagged as 'Dessert'
(4, 9),   -- Gluten-Free Pancakes tagged as 'Low Calorie'
(5, 1),   -- Baked Salmon tagged as 'Healthy'
(5, 6),   -- Baked Salmon tagged as 'High Protein'
(5, 10),  -- Baked Salmon tagged as 'Keto'
(6, 2),   -- Vegan Buddha Bowl tagged as 'Vegan'
(6, 7),   -- Vegan Buddha Bowl tagged as 'Low Carb'
(6, 16),  -- Vegan Buddha Bowl tagged as 'Organic'
(7, 4),   -- Chicken Curry tagged as 'Spicy'
(7, 19),  -- Chicken Curry tagged as 'Comfort Food'
(7, 1),   -- Chicken Curry tagged as 'Healthy'
(8, 5),   -- Fruit Parfait tagged as 'Dessert'
(8, 9),   -- Fruit Parfait tagged as 'Low Calorie'
(9, 1),   -- Lentil Soup tagged as 'Healthy'
(9, 18),  -- Lentil Soup tagged as 'Nut-Free'
(9, 16),  -- Lentil Soup tagged as 'Organic'
(10, 2),  -- Tofu Scramble tagged as 'Vegan'
(10, 9),  -- Tofu Scramble tagged as 'Low Calorie'
(10, 18), -- Tofu Scramble tagged as 'Nut-Free'
(11, 1),  -- Bruschetta tagged as 'Healthy'
(11, 16), -- Bruschetta tagged as 'Organic'
(12, 1),  -- Stuffed Mushrooms tagged as 'Healthy'
(12, 16); -- Stuffed Mushrooms tagged as 'Organic'


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





-- Requested Meals for Users
INSERT INTO requested_meals (user_id, starter_meal_type, main_meal_type, dessert_meal_type, number_of_starter, number_of_main, number_of_dessert) 
VALUES (1, 1, 1, 1, 1, 1, 1),
(2, 1, 1, 1, 1, 1, 2),
(3, 3, 1, 1, 1, 1, 2),
(4, 2, 1, 1, 1, 1, 2),
(5, 1, 1, 1, 1, 1, 2),
(6, 2, 1, 1, 1, 1, 2),
(7, 3, 1, 1, 1, 1, 2),
(8, 1, 1, 1, 1, 1, 2),
(9, 2, 1, 1, 1, 1, 2),
(10, 3,1, 1,  1, 2, 1);