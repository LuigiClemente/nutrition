# API Documentation
#### Base URL:
> /api

User Health Information
All routes under /api/user-health-info handle data related to users' health information.

> POST /api/user-health-info/
Description: Add a new user's health information.
Request Body:

```json
{
        "name": "John Doe",
        "age": 30,
        "gender": "Male",
        "activity_level": "Moderate",
        "blood_glucose": 90.5,
        "health_score": 75.3,
        "nutritional_deficiencies": [
            "Vitamin D",
            "Iron"
        ],
        "allergies": [
            "Peanuts",
            "Shellfish"
        ],
        "body_metrics": {
            "weight": 75.5,
            "height": 175,
            "bmi": 24.6
        },
        "dietary_preferences": {
            "vegetarian": false,
            "vegan": false,
            "gluten_free": true,
            "dairy_free": false,
            "specific_avoidances": [
                "Processed Foods",
                "Sugar"
            ]
        },
        "health_conditions": [
            {
                "name": "Hypertension",
                "severity": "Mild"
            },
            {
                "name": "Diabetes",
                "severity": "Moderate"
            }
        ],
        "goals": [
            {
                "type": "Weight Loss",
                "target": 5.0,
                "duration": 30
            },
            {
                "type": "Improve Cardio",
                "target": 10.0,
                "duration": 60
            }
        ],
        "microbiome_data": {
            "diversity_score": 8.5,
            "gut_health_recommendations": [
                "Increase fiber intake",
                "Reduce red meat consumption"
            ]
        },
        "lipid_profiles": {
            "cholesterol": 180.0,
            "hdl": 60.0,
            "ldl": 100.0,
            "triglycerides": 150.0
        },
        "environmental_factors": {
            "location": "New York",
            "climate": "Temperate",
            "season": "Fall"
        },
        "meal_histories": [
            {
                "meal_id": 1,
                "timestamp": "2024-09-07T12:30:00Z"
            },
            {
                "meal_id": 2,
                "timestamp": "2024-09-06T19:00:00Z"
            }
        ],
        "recent_meals": [
            {
                "meal_id": 3,
                "timestamp": "2024-09-07T18:00:00Z"
            }
        ],
        "preferences": {
            "preferred_cuisines": [
                "Italian",
                "Japanese"
            ],
            "meal_timings": [
                "Breakfast",
                "Lunch",
                "Dinner"
            ],
            "favorite_ingredients": [
                "Chicken",
                "Broccoli",
                "Garlic"
            ]
        }
    }
```

#####  Response:




> GET /api/user-health-info/

Description: Retrieve all users' health information.

#####  Response:




> GET /api/user-health-info/{userId}

Description: Get a specific user’s health information by their userId.

#####  Response:

```json
{
    "user_info": {
        "id": 13,
        "name": "John Doe",
        "age": 30,
        "gender": "Male",
        "body_metrics": {
            "user_id": 13,
            "weight": 75.5,
            "height": 1.75,
            "bmi": 24.6
        },
        "activity_level": "Moderate",
        "dietary_preferences": {
            "user_id": 13,
            "vegetarian": false,
            "vegan": false,
            "gluten_free": true,
            "dairy_free": false,
            "specific_avoidances": [
                "Processed Foods",
                "Sugar"
            ]
        },
        "health_conditions": [
            {
                "user_id": 13,
                "name": "Hypertension",
                "severity": "Mild"
            },
            {
                "user_id": 13,
                "name": "Diabetes",
                "severity": "Moderate"
            }
        ],
        "nutritional_deficiencies": [
            "Vitamin D",
            "Iron"
        ],
        "allergies": [
            "Peanuts",
            "Shellfish"
        ],
        "recent_meals": [
            {
                "id": 8,
                "user_id": 13,
                "meal_id": 3,
                "timestamp": "2024-09-07T18:00:00Z"
            }
        ],
        "goals": [
            {
                "id": 19,
                "user_id": 13,
                "type": "Weight Loss",
                "target": 5,
                "duration": 30
            },
            {
                "id": 20,
                "user_id": 13,
                "type": "Improve Cardio",
                "target": 10,
                "duration": 60
            }
        ],
        "microbiome_data": {
            "id": 10,
            "user_id": 13,
            "diversity_score": 8.5,
            "gut_health_recommendations": [
                "Increase fiber intake",
                "Reduce red meat consumption"
            ]
        },
        "blood_glucose": 90.5,
        "lipid_profiles": {
            "id": 10,
            "user_id": 13,
            "cholesterol": 180,
            "hdl": 60,
            "ldl": 100,
            "triglycerides": 150
        },
        "environmental_factors": {
            "id": 10,
            "user_id": 13,
            "location": "New York",
            "climate": "Temperate",
            "season": "Fall"
        },
        "meal_histories": [
            {
                "id": 7,
                "user_id": 13,
                "meal_id": 7,
                "timestamp": "2024-09-07T12:30:00Z"
            },
            {
                "id": 8,
                "user_id": 13,
                "meal_id": 12,
                "timestamp": "2024-09-06T19:00:00Z"
            }
        ],
        "health_score": 75.3,
        "preferences": {
            "id": 0,
            "user_id": 0,
            "preferred_cuisines": null,
            "meal_timings": null,
            "favorite_ingredients": null
        }
    },
    "recommendations": [
        {
            "meal": "Chicken Caesar Salad",
            "score": 9.68,
            "Tags": [
                "High Protein",
                "Low Carb"
            ],
            "nutritional_content": {
                "fat": 22.8,
                "carbs": 18.1,
                "protein": 43.7,
                "calories": 487
            },
            "health_scores": {
                "weight_loss": 70,
                "heart_health": 80,
                "diabetes_friendly": 75
            }
        },
        {
            "meal": "Grilled Chicken Salad",
            "score": 9.68,
            "Tags": [
                "High Protein",
                "Low Carb",
                "Gluten-Free"
            ],
            "nutritional_content": {
                "fat": 4.4,
                "carbs": 8.4,
                "protein": 34.1,
                "calories": 212
            },
            "health_scores": {
                "weight_loss": 90,
                "heart_health": 85,
                "diabetes_friendly": 90
            }
        },
        {
            "meal": "Mushroom Risotto",
            "score": 9.68,
            "Tags": [
                "Vegetarian",
                "Comfort Food",
                "Rich"
            ],
            "nutritional_content": {
                "fat": 25.3,
                "carbs": 165.5,
                "protein": 25.0,
                "calories": 993
            },
            "health_scores": {
                "weight_loss": 40,
                "heart_health": 60,
                "diabetes_friendly": 55
            }
        }
    ]
}
```



> PUT /api/user-health-info/{userId}

Description: Update a specific user's health information by userId.

##### Request Body:

```json
{
        "name": "John Doe",
        "age": 30,
        "gender": "Male",
        "activity_level": "Moderate",
        "blood_glucose": 90.5,
        "health_score": 75.3,
        "nutritional_deficiencies": [
            "Vitamin D",
            "Iron"
        ],
        "allergies": [
            "Peanuts",
            "Shellfish"
        ],
        "body_metrics": {
            "weight": 75.5,
            "height": 175,
            "bmi": 24.6
        },
        "dietary_preferences": {
            "vegetarian": false,
            "vegan": false,
            "gluten_free": true,
            "dairy_free": false,
            "specific_avoidances": [
                "Processed Foods",
                "Sugar"
            ]
        },
        "health_conditions": [
            {
                "name": "Hypertension",
                "severity": "Mild"
            },
            {
                "name": "Diabetes",
                "severity": "Moderate"
            }
        ],
        "goals": [
            {
                "type": "Weight Loss",
                "target": 5.0,
                "duration": 30
            },
            {
                "type": "Improve Cardio",
                "target": 10.0,
                "duration": 60
            }
        ],
        "microbiome_data": {
            "diversity_score": 8.5,
            "gut_health_recommendations": [
                "Increase fiber intake",
                "Reduce red meat consumption"
            ]
        },
        "lipid_profiles": {
            "cholesterol": 180.0,
            "hdl": 60.0,
            "ldl": 100.0,
            "triglycerides": 150.0
        },
        "environmental_factors": {
            "location": "New York",
            "climate": "Temperate",
            "season": "Fall"
        },
        "meal_histories": [
            {
                "meal_id": 1,
                "timestamp": "2024-09-07T12:30:00Z"
            },
            {
                "meal_id": 2,
                "timestamp": "2024-09-06T19:00:00Z"
            }
        ],
        "recent_meals": [
            {
                "meal_id": 3,
                "timestamp": "2024-09-07T18:00:00Z"
            }
        ],
        "preferences": {
            "preferred_cuisines": [
                "Italian",
                "Japanese"
            ],
            "meal_timings": [
                "Breakfast",
                "Lunch",
                "Dinner"
            ],
            "favorite_ingredients": [
                "Chicken",
                "Broccoli",
                "Garlic"
            ]
        }
    }
```


> DELETE /api/user-health-info/{userId}

Description: Delete a specific user's health information by userId.

##### Response:

Status 200: User health info deleted successfully


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