-- Insert meal categories into the meal_categories table
INSERT INTO meal_categories (category)
VALUES 
('Breakfast'),
('Lunch'),
('Dinner'),
('Snack'),
('Brunch'),
('Supper'),
('Appetizer'),
('Main Course'),
('Dessert'),
('Salad'),
('Sandwich'),
('Pizza'),
('Pasta'),
('Burger'),
('Wraps'),
('Tacos'),
('Soups'),
('Stir-fry'),
('Grill'),
('Seafood'),
('Vegetarian'),
('Vegan'),
('Gluten-Free'),
('Paleo'),
('Keto'),
('Mediterranean'),
('Asian'),
('Indian'),
('Mexican'),
('American'),
('Italian'),
('French'),
('Comfort Food'),
('Fusion'),
('Street Food');



-- Insert meal types into the meal_types table
INSERT INTO meal_types (type)
VALUES 
('Breakfast'),
('Lunch'),
('Dinner'),
('Snack'),
('Brunch'),
('Supper');



-- Insert meals into the meals table
INSERT INTO meals (name, nutritional_content, meal_category_id, meal_type_id, cuisine)
VALUES 
('Pancakes', '{"Calories":350,"Carbs":60,"Protein":8}', 1, 1, 'American'),  -- Breakfast category, Meal Type: Breakfast
('Caesar Salad', '{"Calories":400,"Carbs":20,"Protein":10}', 10, 2, 'Italian'),  -- Salad category, Meal Type: Lunch
('Spaghetti Bolognese', '{"Calories":650,"Carbs":75,"Protein":30}', 3, 3, 'Italian'), -- Dinner category, Meal Type: Dinner
('Margherita Pizza', '{"Calories":700,"Carbs":80,"Protein":20}', 12, 3, 'Italian'),  -- Pizza category, Meal Type: Dinner
('Veggie Burger', '{"Calories":450,"Carbs":40,"Protein":20}', 14, 2, 'American'), -- Burger category, Meal Type: Lunch
('Chicken Tacos', '{"Calories":550,"Carbs":45,"Protein":25}', 16, 2, 'Mexican'), -- Taco category, Meal Type: Lunch
('Chocolate Cake', '{"Calories":400,"Carbs":50,"Protein":5}', 9, 2, 'American'), -- Dessert category, Meal Type: Dessert
('Grilled Salmon', '{"Calories":500,"Carbs":0,"Protein":35}', 20, 3, 'Seafood'), -- Seafood category, Meal Type: Dinner
('Fruit Salad', '{"Calories":200,"Carbs":50,"Protein":2}', 10, 1, 'American'), -- Snack category, Meal Type: Snack
('Quinoa Bowl', '{"Calories":300,"Carbs":40,"Protein":10}', 10, 2, 'Mediterranean'), -- Salad category, Meal Type: Lunch
('French Toast', '{"Calories":400,"Carbs":50,"Protein":10}', 1, 1, 'American'),  -- Breakfast category, Meal Type: Breakfast
('Chicken Stir-fry', '{"Calories":450,"Carbs":30,"Protein":30}', 18, 3, 'Asian'); -- Stir-fry category, Meal Type: Dinner


-- Insert more meals into the meals table
INSERT INTO meals (name, nutritional_content, meal_category_id, meal_type_id, cuisine)
VALUES 
('Greek Yogurt Parfait', '{"Calories":250,"Carbs":35,"Protein":15}', 1, 1, 'Greek'),  -- Breakfast
('Caprese Salad', '{"Calories":350,"Carbs":10,"Protein":12}', 10, 2, 'Italian'),  -- Salad
('Chicken Alfredo', '{"Calories":800,"Carbs":70,"Protein":40}', 3, 3, 'Italian'), -- Dinner
('BBQ Chicken Pizza', '{"Calories":720,"Carbs":85,"Protein":30}', 12, 3, 'American'),  -- Pizza
('Beef Tacos', '{"Calories":500,"Carbs":40,"Protein":30}', 16, 2, 'Mexican'), -- Tacos
('Mango Sticky Rice', '{"Calories":350,"Carbs":80,"Protein":5}', 9, 2, 'Thai'), -- Dessert
('Shrimp Fried Rice', '{"Calories":600,"Carbs":70,"Protein":25}', 20, 3, 'Chinese'), -- Seafood
('Veggie Stir-fry', '{"Calories":300,"Carbs":40,"Protein":10}', 18, 2, 'Asian'), -- Stir-fry
('Peanut Butter Banana Smoothie', '{"Calories":400,"Carbs":50,"Protein":15}', 1, 1, 'American'), -- Breakfast
('Roasted Vegetable Quinoa', '{"Calories":450,"Carbs":60,"Protein":12}', 10, 2, 'Mediterranean'), -- Salad
('Omelette', '{"Calories":300,"Carbs":5,"Protein":20}', 1, 1, 'French'),  -- Breakfast
('Chocolate Chip Cookies', '{"Calories":150,"Carbs":25,"Protein":2}', 9, 2, 'American'), -- Dessert
('Beef Stir-fry', '{"Calories":500,"Carbs":30,"Protein":35}', 18, 3, 'Asian'), -- Stir-fry
('Lentil Soup', '{"Calories":250,"Carbs":40,"Protein":18}', 3, 2, 'Mediterranean'); -- Soup category







-- Insert users into the users table
INSERT INTO users (name, age, gender, activity_level, blood_glucose, health_score, nutritional_deficiencies, allergies)
VALUES 
('John Doe', 30, 'Male', 'Moderate', 90.5, 85.0, ARRAY['Vitamin D'], ARRAY['Peanuts']),
('Jane Smith', 25, 'Female', 'High', 85.3, 90.0, ARRAY['Iron'], ARRAY['Shellfish']),
('Emily Johnson', 40, 'Female', 'Low', 110.7, 78.5, ARRAY['Calcium'], ARRAY['None']),
('Michael Brown', 35, 'Male', 'Moderate', 100.1, 80.0, ARRAY['B12'], ARRAY['Lactose']),
('Sarah Davis', 28, 'Female', 'High', 92.4, 88.0, ARRAY['Folate'], ARRAY['Gluten']);



-- Insert body metrics into body_metrics table
INSERT INTO body_metrics (user_id, weight, height)
VALUES 
(1, 75.0, 1.80),  -- John Doe
(2, 65.0, 1.70),  -- Jane Smith
(3, 70.0, 1.65),  -- Emily Johnson
(4, 85.0, 1.85),  -- Michael Brown
(5, 60.0, 1.68);  -- Sarah Davis



-- Insert dietary preferences into dietary_preferences table
INSERT INTO dietary_preferences (user_id, vegetarian, vegan, gluten_free, dairy_free, specific_avoidances)
VALUES 
(1, FALSE, FALSE, FALSE, TRUE, ARRAY['None']),        -- John Doe
(2, TRUE, FALSE, FALSE, TRUE, ARRAY['Shellfish']),     -- Jane Smith
(3, FALSE, FALSE, TRUE, FALSE, ARRAY['Gluten']),       -- Emily Johnson
(4, FALSE, FALSE, FALSE, TRUE, ARRAY['Lactose']),      -- Michael Brown
(5, TRUE, TRUE, TRUE, TRUE, ARRAY['Gluten', 'Dairy']); -- Sarah Davis



-- Insert health conditions into health_conditions table
INSERT INTO health_conditions (user_id, name, severity)
VALUES 
(1, 'Hypertension', 'Moderate'),  -- John Doe
(2, 'Asthma', 'Mild'),            -- Jane Smith
(3, 'Diabetes', 'Severe'),        -- Emily Johnson
(4, 'None', 'None'),              -- Michael Brown
(5, 'Celiac Disease', 'Moderate');-- Sarah Davis


-- Insert requested meals into requested_meals table
INSERT INTO requested_meals (user_id, meal_category, timestamp)
VALUES 
(1, 'Pasta', NOW()),        -- John Doe
(2, 'Salad', NOW()),        -- Jane Smith
(3, 'Pizza', NOW()),        -- Emily Johnson
(4, 'Steak', NOW()),        -- Michael Brown
(5, 'Vegan Burger', NOW()); -- Sarah Davis


-- Insert goals into goals table
INSERT INTO goals (user_id, type, target, duration)
VALUES 
(1, 'Weight Loss', 5.0, 30),  -- John Doe
(2, 'Muscle Gain', 3.0, 45),  -- Jane Smith
(3, 'Blood Sugar Management', 15.0, 60),  -- Emily Johnson
(4, 'Heart Health', 10.0, 90),  -- Michael Brown
(5, 'Weight Maintenance', 0.0, 0);  -- Sarah Davis



-- Insert health conditions into health_conditions table
INSERT INTO health_conditions (user_id, name, severity)
VALUES 
(1, 'Hypertension', 'Moderate'),  -- John Doe
(2, 'Asthma', 'Mild'),            -- Jane Smith
(3, 'Diabetes', 'Severe'),        -- Emily Johnson
(4, 'None', 'None'),              -- Michael Brown
(5, 'Celiac Disease', 'Moderate');-- Sarah Davis




-- Insert lipid profiles into lipid_profiles table
INSERT INTO lipid_profiles (user_id, cholesterol, hdl, ldl, triglycerides)
VALUES 
(1, 200.0, 60.0, 130.0, 150.0),  -- John Doe
(2, 180.0, 55.0, 120.0, 140.0),  -- Jane Smith
(3, 220.0, 50.0, 140.0, 160.0),  -- Emily Johnson
(4, 190.0, 65.0, 110.0, 130.0),  -- Michael Brown
(5, 210.0, 70.0, 115.0, 145.0);  -- Sarah Davis



-- Insert environmental factors into environmental_factors table
INSERT INTO environmental_factors (user_id, location, climate, season)
VALUES 
(1, 'New York', 'Temperate', 'Spring'),    -- John Doe
(2, 'Los Angeles', 'Mediterranean', 'Fall'), -- Jane Smith
(3, 'Chicago', 'Continental', 'Winter'),   -- Emily Johnson
(4, 'Houston', 'Humid Subtropical', 'Summer'), -- Michael Brown
(5, 'Denver', 'Mountain', 'Winter');       -- Sarah Davis



-- Insert possible meal tags into meal_tags table
INSERT INTO meal_tags (tag)
VALUES 
('Gluten-Free'),
('Vegan'),
('Vegetarian'),
('Dairy-Free'),
('High-Protein'),
('Low-Carb'),
('Low-Fat'),
('Low-Calorie'),
('Keto'),
('Paleo'),
('Whole30'),
('Nut-Free'),
('Soy-Free'),
('Sugar-Free'),
('Low-Sodium'),
('Heart-Healthy'),
('Diabetic-Friendly'),
('High-Fiber'),
('Kid-Friendly'),
('Quick & Easy'),
('Budget-Friendly'),
('Meal Prep'),
('Comfort Food'),
('Spicy'),
('Sweet'),
('Savory'),
('Mediterranean'),
('Italian'),
('Mexican'),
('Asian'),
('Indian'),
('American'),
('Middle Eastern'),
('French'),
('Chinese'),
('Japanese'),
('Korean'),
('Thai'),
('Seafood'),
('Grilled'),
('Baked'),
('Fried'),
('Raw'),
('Slow-Cooked'),
('Stir-Fried'),
('Roasted'),
('Steamed'),
('One-Pot'),
('Salad'),
('Soup'),
('Appetizer'),
('Main Course'),
('Side Dish'),
('Snack'),
('Dessert'),
('Breakfast'),
('Lunch'),
('Dinner'),
('Brunch'),
('Beverage'),
('Smoothie'),
('Cocktail'),
('Barbecue'),
('Holiday'),
('Party Food'),
('Finger Food');

