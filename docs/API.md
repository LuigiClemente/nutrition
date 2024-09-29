# API Documentation
#### Base URL:
> /api

## User Health Information
All routes under /api/user-health-info handle data related to users' health information.

> POST /api/user-health-info/

Description: Add a new user's health information.

##### Request Body:

```json
{
    "name": "Ahmad",
    "age": 30,
    "gender": "Male",
    "activity_level": "Moderate",
    "blood_glucose": 90,
    "health_score": 85,
    "nutritional_deficiencies": [
        "Vitamin D",
        "Iron"
    ],
    "allergies": [
        "Peanuts"
    ],
    "body_metrics": {
        "weight": 70,
        "height": 175
    },
    "dietary_preferences": {
        "vegetarian": false,
        "vegan": false,
        "gluten_free": false,
        "dairy_free": false,
        "specific_avoidances": [
            "Shellfish"
        ]
    },
    "health_conditions": [
        {
            "name": "Hypertension",
            "severity": "Moderate"
        }
    ],
    "goals": [
        {
            "type": "Weight loss",
            "target": 5,
            "duration": 90
        }
    ],
    "microbiome_data": {
        "diversity_score": 75,
        "gut_health_recommendations": [
            "Increase fiber intake"
        ]
    },
    "lipid_profile": {
        "cholesterol": 180,
        "hdl": 60,
        "ldl": 100,
        "triglycerides": 150
    },
    "environmental_factors": {
        "location": "Urban",
        "climate": "Temperate",
        "season": "Fall"
    },
    "meal_history": [
        {
            "meal_id": 1,
            "timestamp": "2024-09-28T12:00:00Z"
        }
    ],
    "recent_meals": [
        {
            "meal_id": 2,
            "timestamp": "2024-09-27T18:00:00Z"
        }
    ],
    "preferences": {
        "preferred_cuisines": [
            "Italian",
            "Asian"
        ],
        "meal_timings": [
            "Lunch",
            "Dinner"
        ],
        "favorite_ingredients": [
            "Chicken",
            "Broccoli"
        ]
    },
    "requested_meal": {
        "meal_category": 1
    }
}
```

#####  Response:

```json

```


> GET /api/user-health-info/

Description: Retrieve all users' health information.

#####  Response:




> GET /api/user-health-info/{userId}

Description: Get a specific user’s health information by their userId.

#####  Response:

```json
{
    "user_info": {
        "id": 2,
        "name": "Jane Smith",
        "age": 25,
        "gender": "Female",
        "body_metrics": {
            "user_id": 2,
            "weight": 60,
            "height": 165
        },
        "activity_level": "Active",
        "dietary_preferences": {
            "user_id": 2,
            "vegetarian": true,
            "vegan": false,
            "gluten_free": true,
            "dairy_free": false,
            "specific_avoidances": [
                "Processed sugar"
            ]
        },
        "health_conditions": [
            {
                "user_id": 2,
                "name": "Asthma",
                "severity": "Mild"
            }
        ],
        "nutritional_deficiencies": [],
        "allergies": [],
        "recent_meals": [],
        "goals": [
            {
                "id": 2,
                "user_id": 2,
                "type": "Muscle gain",
                "target": 3,
                "duration": 60
            }
        ],
        "microbiome_data": {
            "user_id": 2,
            "diversity_score": 80,
            "gut_health_recommendations": [
                "Increase plant-based foods"
            ]
        },
        "blood_glucose": 85.3,
        "lipid_profiles": {
            "user_id": 2,
            "cholesterol": 200,
            "hdl": 70,
            "ldl": 120,
            "triglycerides": 130
        },
        "environmental_factors": {
            "user_id": 2,
            "location": "Los Angeles",
            "climate": "Mediterranean",
            "season": "Summer"
        },
        "requested_meal": {
            "user_id": 2,
            "meal_category": 10
        },
        "meal_histories": [],
        "health_score": 90,
        "preferences": {
            "user_id": 2,
            "preferred_cuisines": [
                "Mexican",
                "Japanese"
            ],
            "meal_timings": [
                "Breakfast",
                "Dinner"
            ],
            "favorite_ingredients": [
                "Avocado",
                "Eggs"
            ]
        }
    },
    "recommendation": {
        "meal": {
            "id": 22,
            "name": "Roasted Vegetable Quinoa",
            "ingredients": [
                {
                    "name": "Quinoa",
                    "amount": 100,
                    "portion": "1/2 cup",
                    "ounces": "3.53"
                },
                {
                    "name": "Zucchini",
                    "amount": 50,
                    "portion": "1/2 cup",
                    "ounces": "1.76"
                },
                {
                    "name": "Bell Peppers",
                    "amount": 50,
                    "portion": "1/2 cup",
                    "ounces": "1.76"
                },
                {
                    "name": "Olive Oil",
                    "amount": 15,
                    "portion": "1 tbsp",
                    "ounces": "0.53"
                },
                {
                    "name": "Lemon Juice",
                    "amount": 10,
                    "portion": "1 tbsp",
                    "ounces": "0.35"
                }
            ]
        },
        "score": 18.18
    }
}
```



> PUT /api/user-health-info/{userId}

Description: Update a specific user's health information by userId.

##### Request Body:

```json
{
    "name": "Ahmad",
    "age": 30,
    "gender": "Male",
    "activity_level": "Moderate",
    "blood_glucose": 90,
    "health_score": 85,
    "nutritional_deficiencies": [
        "Vitamin D",
        "Iron"
    ],
    "allergies": [
        "Peanuts"
    ],
    "body_metrics": {
        "weight": 70,
        "height": 175
    },
    "dietary_preferences": {
        "vegetarian": false,
        "vegan": false,
        "gluten_free": false,
        "dairy_free": false,
        "specific_avoidances": [
            "Shellfish"
        ]
    },
    "health_conditions": [
        {
            "name": "Hypertension",
            "severity": "Moderate"
        }
    ],
    "goals": [
        {
            "type": "Weight loss",
            "target": 5,
            "duration": 90
        }
    ],
    "microbiome_data": {
        "diversity_score": 75,
        "gut_health_recommendations": [
            "Increase fiber intake"
        ]
    },
    "lipid_profile": {
        "cholesterol": 180,
        "hdl": 60,
        "ldl": 100,
        "triglycerides": 150
    },
    "environmental_factors": {
        "location": "Urban",
        "climate": "Temperate",
        "season": "Fall"
    },
    "meal_history": [
        {
            "meal_id": 1,
            "timestamp": "2024-09-28T12:00:00Z"
        }
    ],
    "recent_meals": [
        {
            "meal_id": 2,
            "timestamp": "2024-09-27T18:00:00Z"
        }
    ],
    "preferences": {
        "preferred_cuisines": [
            "Italian",
            "Asian"
        ],
        "meal_timings": [
            "Lunch",
            "Dinner"
        ],
        "favorite_ingredients": [
            "Chicken",
            "Broccoli"
        ]
    },
    "requested_meal": {
        "meal_category": 1
    }
}
```

#### Response:

```json
{
    "user_info": {
        "id": 2,
        "name": "Jane Smith",
        "age": 25,
        "gender": "Female",
        "body_metrics": {
            "user_id": 2,
            "weight": 60,
            "height": 165
        },
        "activity_level": "Active",
        "dietary_preferences": {
            "user_id": 2,
            "vegetarian": true,
            "vegan": false,
            "gluten_free": true,
            "dairy_free": false,
            "specific_avoidances": [
                "Processed sugar"
            ]
        },
        "health_conditions": [
            {
                "user_id": 2,
                "name": "Asthma",
                "severity": "Mild"
            }
        ],
        "nutritional_deficiencies": [],
        "allergies": [],
        "recent_meals": [],
        "goals": [
            {
                "id": 2,
                "user_id": 2,
                "type": "Muscle gain",
                "target": 3,
                "duration": 60
            }
        ],
        "microbiome_data": {
            "user_id": 2,
            "diversity_score": 80,
            "gut_health_recommendations": [
                "Increase plant-based foods"
            ]
        },
        "blood_glucose": 85.3,
        "lipid_profiles": {
            "user_id": 2,
            "cholesterol": 200,
            "hdl": 70,
            "ldl": 120,
            "triglycerides": 130
        },
        "environmental_factors": {
            "user_id": 2,
            "location": "Los Angeles",
            "climate": "Mediterranean",
            "season": "Summer"
        },
        "requested_meal": {
            "user_id": 2,
            "meal_category": 10
        },
        "meal_histories": [],
        "health_score": 90,
        "preferences": {
            "user_id": 2,
            "preferred_cuisines": [
                "Mexican",
                "Japanese"
            ],
            "meal_timings": [
                "Breakfast",
                "Dinner"
            ],
            "favorite_ingredients": [
                "Avocado",
                "Eggs"
            ]
        }
    },
    "recommendation": {
        "meal": {
            "id": 22,
            "name": "Roasted Vegetable Quinoa",
            "ingredients": [
                {
                    "name": "Quinoa",
                    "amount": 100,
                    "portion": "1/2 cup",
                    "ounces": "3.53"
                },
                {
                    "name": "Zucchini",
                    "amount": 50,
                    "portion": "1/2 cup",
                    "ounces": "1.76"
                },
                {
                    "name": "Bell Peppers",
                    "amount": 50,
                    "portion": "1/2 cup",
                    "ounces": "1.76"
                },
                {
                    "name": "Olive Oil",
                    "amount": 15,
                    "portion": "1 tbsp",
                    "ounces": "0.53"
                },
                {
                    "name": "Lemon Juice",
                    "amount": 10,
                    "portion": "1 tbsp",
                    "ounces": "0.35"
                }
            ]
        },
        "score": 18.18
    }
}
```


> DELETE /api/user-health-info/{userId}

Description: Delete a specific user's health information by userId.

##### Response:

Status 200: User health info deleted successfully



## Meals
All routes under /api/meals handle meal data and meal options.

> POST /api/meals/

Description: Add a new meal.

#### Request Body:

```json
[
  {
    "name": "Grilled Chicken Salad",
    "ingredients": [
      {
        "name": "Chicken Breast",
        "amount": 150,
        "portion": "1 cup"
      },
      {
        "name": "Lettuce",
        "amount": 50,
        "portion": "2 cups"
      },
      {
        "name": "Tomatoes",
        "amount": 30,
        "portion": "1/2 cup"
      }
    ],
    "nutritional_content": {
      "calories": 400,
      "protein": 35,
      "fat": 10,
      "carbohydrates": 20
    },
    "meal_category_id": 10,
    "meal_type_id": 2,
    "cuisine": "American",
    "tags": [
      {
        "id": 6,
        "tag": "Low-Carb"
      },
      {
        "id": 49,
        "tag": "Salad"
      }
    ]
  }
]
```

#### Response:

Status 201

> GET /api/meals/

Description: Retrieve all meals.

#### Response:

```json
[
  {
    "name": "Grilled Chicken Salad",
    "ingredients": [
      {
        "name": "Chicken Breast",
        "amount": 150,
        "portion": "1 cup"
      },
      {
        "name": "Lettuce",
        "amount": 50,
        "portion": "2 cups"
      },
      {
        "name": "Tomatoes",
        "amount": 30,
        "portion": "1/2 cup"
      }
    ],
    "nutritional_content": {
      "calories": 400,
      "protein": 35,
      "fat": 10,
      "carbohydrates": 20
    },
    "meal_category_id": 10,
    "meal_type_id": 2,
    "cuisine": "American",
    "tags": [
      {
        "id": 6,
        "tag": "Low-Carb"
      },
      {
        "id": 49,
        "tag": "Salad"
      }
    ]
  }
]
```

> GET /api/meals/options

Description: Retrieve meal options based on user preferences.

#### Response:

```json
[
    {
        "id": 5,
        "name": "Lentil Soup"
        
    },
    {
        "id": 4,
        "name": "Quinoa Salad"
       
    },
    {
        "id": 3,
        "name": "Vegetable Stir Fry"
       
    }
]
```

> GET /api/meals/{mealId}

Description: Get a specific meal's information by mealId.

### Response:

```json
{
    "id": 47,
    "name": "Grilled Chicken Salad",
    "ingredients": [
        {
            "id": 215,
            "meal_id": 47,
            "name": "Chicken Breast",
            "amount": 150,
            "portion": "1 cup"
        },
        {
            "id": 216,
            "meal_id": 47,
            "name": "Lettuce",
            "amount": 50,
            "portion": "2 cups"
        },
        {
            "id": 217,
            "meal_id": 47,
            "name": "Tomatoes",
            "amount": 30,
            "portion": "1/2 cup"
        }
    ],
    "nutritional_content": {
        "fat": 10,
        "protein": 35,
        "calories": 400,
        "carbohydrates": 20
    },
    "meal_category_id": 10,
    "meal_category": {
        "id": 10,
        "category": "Salad"
    },
    "meal_type_id": 2,
    "meal_type": {
        "id": 2,
        "type": "Lunch"
    },
    "cuisine": "American",
    "tags": [
        {
            "id": 6,
            "tag": "Low-Carb"
        },
        {
            "id": 49,
            "tag": "Salad"
        }
    ]
}
```

> PUT /api/meals/{mealId}

Description: Update a specific meal's information by mealId.

#### Request Body:

```json
    {
        "name": "Grilled Chicken Salad",
        "ingredients": [
            {
                "name": "Chicken Breast",
                "amount": 150,
                "portion": "1 cup"
            },
            {
                "name": "Lettuce",
                "amount": 50,
                "portion": "2 cups"
            },
            {
                "name": "Tomatoes",
                "amount": 30,
                "portion": "1/2 cup"
            }
        ],
        "nutritional_content": {
            "calories": 400,
            "protein": 35,
            "fat": 10,
            "carbohydrates": 20
        },
        "meal_category_id": 10,
        "meal_type_id": 1,
        "cuisine": "American",
        "tags": [
            {
                
                "id": 22,
                "tag": "Comfort Food"
            }
        ]
    }
```


> DELETE /api/meals/{mealId}

Description: Delete a specific meal by mealId

### Goal Types:
   - "Weight loss"
   - "Muscle gain"
   - "Heart health"
   - "Blood sugar management"
   - "Gut health"
   - "Metabolic health"
   - "Cholesterol reduction"
   - "Weight maintenance"
   - "Lean muscle maintenance"
   - "Improved energy levels"
   - "Endurance training support"
   - "Cardiovascular fitness"
   - "Detoxification"
   - "Bone health"
   - "Skin health"
   - "Anti-inflammatory diet"
   - "Hormonal balance"
   - "Improved mental clarity"

### Nutritional and Health-related Keys:
 - "calories"
 - "sugar"
 - "fiber"
 - "protein"
 - "fat"
 - "cholesterol"
 - "sodium"
 - "carbs"
 - "antioxidants"
 - "calcium"
 - "vitamin A"
 - "vitamin C"
 - "healthy fats"
 - "omega-3"



# Possible Meal Categories
- Breakfast
- Lunch
- Dinner
- Snack
- Brunch
- Supper
- Appetizer
- Main Course
- Dessert
- Salad
- Sandwich
- Pizza
- Pasta
- Burger
- Wraps
- Tacos
- Soups
- Stir-fry
- Grill
- Seafood
- Vegetarian
- Vegan
- Gluten-Free
- Paleo
- Keto
- Mediterranean
- Asian
- Indian
- Mexican
- American
- Italian
- French
- Comfort Food
- Fusion
- Street Food





Explanation of the Tags
Dietary Tags: Tags like Gluten-Free, Vegan, Vegetarian, Dairy-Free, and others are used to indicate specific dietary preferences or restrictions.
Nutritional Tags: High-Protein, Low-Carb, Low-Fat, Keto, and Paleo help categorize meals based on their nutritional profile.
Health-Related Tags: Tags like Heart-Healthy, Diabetic-Friendly, and Low-Sodium provide a focus on health-conscious eating.
Cuisine Tags: These cover various cuisines, such as Mediterranean, Italian, Mexican, Asian, Indian, etc.
Cooking Methods: Tags like Grilled, Baked, Fried, Slow-Cooked, Roasted, and Steamed indicate the preparation method.
Meal Categories: Tags such as Breakfast, Lunch, Dinner, Appetizer, Main Course, and Dessert categorize the type of meal.
Miscellaneous Tags: Tags like Kid-Friendly, Quick & Easy, Budget-Friendly, Comfort Food, Party Food, and Finger Food help with specific user preferences.