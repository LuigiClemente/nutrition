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

```json
{
    "user_info": {
        "id": 4,
        "name": "Ahmad",
        "age": 25,
        "gender": "Male",
        "body_metrics": {
            "user_id": 4,
            "weight": 78,
            "height": 175,
            "bmi": 24.6
        },
        "activity_level": "Moderate",
        "dietary_preferences": {
            "user_id": 4,
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
                "user_id": 4,
                "name": "Hypertension",
                "severity": "Mild"
            },
            {
                "user_id": 4,
                "name": "Diabetes",
                "severity": "Moderate"
            }
        ],
        "nutritional_deficiencies": [
            "Vitamin A"
        ],
        "allergies": [
            "Shellfish"
        ],
        "recent_meals": [
            {
                "id": 6,
                "user_id": 4,
                "meal_id": 3,
                "timestamp": "2024-09-07T18:00:00Z"
            }
        ],
        "goals": [
            {
                "id": 11,
                "user_id": 4,
                "type": "Muscle gain",
                "target": 5,
                "duration": 90
            },
            {
                "id": 12,
                "user_id": 4,
                "type": "Improve Cardio",
                "target": 10,
                "duration": 60
            }
        ],
        "microbiome_data": {
            "user_id": 4,
            "diversity_score": 8.5,
            "gut_health_recommendations": [
                "Increase fiber intake",
                "Reduce red meat consumption"
            ]
        },
        "blood_glucose": 80.5,
        "lipid_profiles": {
            "user_id": 4,
            "cholesterol": 180,
            "hdl": 60,
            "ldl": 100,
            "triglycerides": 150
        },
        "environmental_factors": {
            "id": 0,
            "user_id": 4,
            "location": "New York",
            "climate": "Temperate",
            "season": "Fall"
        },
        "meal_histories": [
            {
                "id": 3,
                "user_id": 4,
                "meal_id": 1,
                "timestamp": "2024-09-07T12:30:00Z"
            },
            {
                "id": 4,
                "user_id": 4,
                "meal_id": 2,
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
            "meal_name": "Oatmeal",
            "ingredients": [
                {
                    "id": 37,
                    "meal_id": 10,
                    "name": "Oats",
                    "amount": 100,
                    "unit": "g",
                    "nutritional": {
                        "fat": 6.9,
                        "carbs": 66.3,
                        "protein": 16.9,
                        "calories": 389
                    }
                },
                {
                    "id": 38,
                    "meal_id": 10,
                    "name": "Milk",
                    "amount": 200,
                    "unit": "ml",
                    "nutritional": {
                        "fat": 4.5,
                        "carbs": 12.0,
                        "protein": 8.0,
                        "calories": 122
                    }
                },
                {
                    "id": 39,
                    "meal_id": 10,
                    "name": "Banana",
                    "amount": 80,
                    "unit": "g",
                    "nutritional": {
                        "fat": 0.2,
                        "carbs": 18.5,
                        "protein": 0.9,
                        "calories": 72
                    }
                }
            ],
            "score": 67.74,
            "Tags": [
                "Vegetarian",
                "High Fiber",
                "Easy"
            ],
            "nutritional_content": {
                "fat": 11.6,
                "carbs": 96.8,
                "protein": 25.8,
                "calories": 583
            }
        },
        {
            "meal_name": "Grilled Salmon",
            "ingredients": [
                {
                    "id": 75,
                    "meal_id": 22,
                    "name": "Salmon Fillet",
                    "amount": 200,
                    "unit": "g",
                    "nutritional": {
                        "fat": 20.0,
                        "carbs": 0,
                        "protein": 25.0,
                        "calories": 280
                    }
                },
                {
                    "id": 76,
                    "meal_id": 22,
                    "name": "Lemon Juice",
                    "amount": 1,
                    "unit": "tbsp",
                    "nutritional": {
                        "fat": 0,
                        "carbs": 1.3,
                        "protein": 0.1,
                        "calories": 4
                    }
                },
                {
                    "id": 77,
                    "meal_id": 22,
                    "name": "Olive Oil",
                    "amount": 1,
                    "unit": "tbsp",
                    "nutritional": {
                        "fat": 14.0,
                        "carbs": 0,
                        "protein": 0,
                        "calories": 120
                    }
                }
            ],
            "score": 54.84,
            "Tags": [
                "High Protein",
                "Gluten-Free",
                "Low Carb"
            ],
            "nutritional_content": {
                "fat": 34.0,
                "carbs": 1.3,
                "protein": 25.1,
                "calories": 404
            }
        },
        {
            "meal_name": "Grilled Salmon",
            "ingredients": [
                {
                    "id": 43,
                    "meal_id": 12,
                    "name": "Salmon Fillet",
                    "amount": 200,
                    "unit": "g",
                    "nutritional": {
                        "fat": 20.0,
                        "carbs": 0,
                        "protein": 25.0,
                        "calories": 280
                    }
                },
                {
                    "id": 44,
                    "meal_id": 12,
                    "name": "Lemon Juice",
                    "amount": 1,
                    "unit": "tbsp",
                    "nutritional": {
                        "fat": 0,
                        "carbs": 1.3,
                        "protein": 0.1,
                        "calories": 4
                    }
                },
                {
                    "id": 45,
                    "meal_id": 12,
                    "name": "Olive Oil",
                    "amount": 1,
                    "unit": "tbsp",
                    "nutritional": {
                        "fat": 14.0,
                        "carbs": 0,
                        "protein": 0,
                        "calories": 120
                    }
                }
            ],
            "score": 54.84,
            "Tags": [
                "High Protein",
                "Gluten-Free",
                "Low Carb"
            ],
            "nutritional_content": {
                "fat": 34.0,
                "carbs": 1.3,
                "protein": 25.1,
                "calories": 404
            }
        }
    ]
}
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
        "id": 4,
        "name": "Ahmad",
        "age": 25,
        "gender": "Male",
        "body_metrics": {
            "user_id": 4,
            "weight": 78,
            "height": 175,
            "bmi": 24.6
        },
        "activity_level": "Moderate",
        "dietary_preferences": {
            "user_id": 4,
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
                "user_id": 4,
                "name": "Hypertension",
                "severity": "Mild"
            },
            {
                "user_id": 4,
                "name": "Diabetes",
                "severity": "Moderate"
            }
        ],
        "nutritional_deficiencies": [
            "Vitamin A"
        ],
        "allergies": [
            "Shellfish"
        ],
        "recent_meals": [
            {
                "id": 6,
                "user_id": 4,
                "meal_id": 3,
                "timestamp": "2024-09-07T18:00:00Z"
            }
        ],
        "goals": [
            {
                "id": 11,
                "user_id": 4,
                "type": "Muscle gain",
                "target": 5,
                "duration": 90
            },
            {
                "id": 12,
                "user_id": 4,
                "type": "Improve Cardio",
                "target": 10,
                "duration": 60
            }
        ],
        "microbiome_data": {
            "user_id": 4,
            "diversity_score": 8.5,
            "gut_health_recommendations": [
                "Increase fiber intake",
                "Reduce red meat consumption"
            ]
        },
        "blood_glucose": 80.5,
        "lipid_profiles": {
            "user_id": 4,
            "cholesterol": 180,
            "hdl": 60,
            "ldl": 100,
            "triglycerides": 150
        },
        "environmental_factors": {
            "id": 0,
            "user_id": 4,
            "location": "New York",
            "climate": "Temperate",
            "season": "Fall"
        },
        "meal_histories": [
            {
                "id": 3,
                "user_id": 4,
                "meal_id": 1,
                "timestamp": "2024-09-07T12:30:00Z"
            },
            {
                "id": 4,
                "user_id": 4,
                "meal_id": 2,
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
            "meal_name": "Oatmeal",
            "ingredients": [
                {
                    "id": 37,
                    "meal_id": 10,
                    "name": "Oats",
                    "amount": 100,
                    "unit": "g",
                    "nutritional": {
                        "fat": 6.9,
                        "carbs": 66.3,
                        "protein": 16.9,
                        "calories": 389
                    }
                },
                {
                    "id": 38,
                    "meal_id": 10,
                    "name": "Milk",
                    "amount": 200,
                    "unit": "ml",
                    "nutritional": {
                        "fat": 4.5,
                        "carbs": 12.0,
                        "protein": 8.0,
                        "calories": 122
                    }
                },
                {
                    "id": 39,
                    "meal_id": 10,
                    "name": "Banana",
                    "amount": 80,
                    "unit": "g",
                    "nutritional": {
                        "fat": 0.2,
                        "carbs": 18.5,
                        "protein": 0.9,
                        "calories": 72
                    }
                }
            ],
            "score": 67.74,
            "Tags": [
                "Vegetarian",
                "High Fiber",
                "Easy"
            ],
            "nutritional_content": {
                "fat": 11.6,
                "carbs": 96.8,
                "protein": 25.8,
                "calories": 583
            }
        },
        {
            "meal_name": "Grilled Salmon",
            "ingredients": [
                {
                    "id": 75,
                    "meal_id": 22,
                    "name": "Salmon Fillet",
                    "amount": 200,
                    "unit": "g",
                    "nutritional": {
                        "fat": 20.0,
                        "carbs": 0,
                        "protein": 25.0,
                        "calories": 280
                    }
                },
                {
                    "id": 76,
                    "meal_id": 22,
                    "name": "Lemon Juice",
                    "amount": 1,
                    "unit": "tbsp",
                    "nutritional": {
                        "fat": 0,
                        "carbs": 1.3,
                        "protein": 0.1,
                        "calories": 4
                    }
                },
                {
                    "id": 77,
                    "meal_id": 22,
                    "name": "Olive Oil",
                    "amount": 1,
                    "unit": "tbsp",
                    "nutritional": {
                        "fat": 14.0,
                        "carbs": 0,
                        "protein": 0,
                        "calories": 120
                    }
                }
            ],
            "score": 54.84,
            "Tags": [
                "High Protein",
                "Gluten-Free",
                "Low Carb"
            ],
            "nutritional_content": {
                "fat": 34.0,
                "carbs": 1.3,
                "protein": 25.1,
                "calories": 404
            }
        },
        {
            "meal_name": "Grilled Salmon",
            "ingredients": [
                {
                    "id": 43,
                    "meal_id": 12,
                    "name": "Salmon Fillet",
                    "amount": 200,
                    "unit": "g",
                    "nutritional": {
                        "fat": 20.0,
                        "carbs": 0,
                        "protein": 25.0,
                        "calories": 280
                    }
                },
                {
                    "id": 44,
                    "meal_id": 12,
                    "name": "Lemon Juice",
                    "amount": 1,
                    "unit": "tbsp",
                    "nutritional": {
                        "fat": 0,
                        "carbs": 1.3,
                        "protein": 0.1,
                        "calories": 4
                    }
                },
                {
                    "id": 45,
                    "meal_id": 12,
                    "name": "Olive Oil",
                    "amount": 1,
                    "unit": "tbsp",
                    "nutritional": {
                        "fat": 14.0,
                        "carbs": 0,
                        "protein": 0,
                        "calories": 120
                    }
                }
            ],
            "score": 54.84,
            "Tags": [
                "High Protein",
                "Gluten-Free",
                "Low Carb"
            ],
            "nutritional_content": {
                "fat": 34.0,
                "carbs": 1.3,
                "protein": 25.1,
                "calories": 404
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

#### Response:

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
        "name": "Chickpea Curry",
        "ingredients": [
            {
                "name": "Chickpeas",
                "amount": 200,
                "unit": "g",
                "nutritional": {
                    "calories": 364,
                    "protein": 19.0,
                    "fat": 6.0,
                    "carbs": 61.0
                }
            },
            {
                "name": "Coconut Milk",
                "amount": 100,
                "unit": "ml",
                "nutritional": {
                    "calories": 200,
                    "protein": 2.0,
                    "fat": 20.0,
                    "carbs": 5.0
                }
            },
            {
                "name": "Onion",
                "amount": 80,
                "unit": "g",
                "nutritional": {
                    "calories": 32,
                    "protein": 0.9,
                    "fat": 0.1,
                    "carbs": 7.6
                }
            },
            {
                "name": "Curry Powder",
                "amount": 1,
                "unit": "tbsp",
                "nutritional": {
                    "calories": 20,
                    "protein": 1.0,
                    "fat": 0.5,
                    "carbs": 4.0
                }
            }
        ],
        "nutritional_content": {
            "calories": 616,
            "protein": 23.0,
            "fat": 26.6,
            "carbs": 77.6
        },
        "category": "Main Dish",
        "meal_type": [
            "Lunch",
            "Dinner"
        ],
        "cuisine": "Indian",
        "tags": [
            "Vegan",
            "High Protein",
            "Gluten-Free"
        ],
        "health_scores": {
            "heart_health": 80,
            "diabetes_friendly": 75,
            "weight_loss": 70
        },
        "preparation_time": 30,
        "difficulty": "Medium",
        "serving_size": 2,
        "instructions": "Simmer chickpeas with coconut milk, onion, and curry powder. Serve with rice."
    },
    {
        "name": "Mushroom Risotto",
        "ingredients": [
            {
                "name": "Arborio Rice",
                "amount": 200,
                "unit": "g",
                "nutritional": {
                    "calories": 720,
                    "protein": 13.0,
                    "fat": 3.0,
                    "carbs": 158.0
                }
            },
            {
                "name": "Mushrooms",
                "amount": 150,
                "unit": "g",
                "nutritional": {
                    "calories": 33,
                    "protein": 4.0,
                    "fat": 0.3,
                    "carbs": 6.5
                }
            },
            {
                "name": "Parmesan",
                "amount": 30,
                "unit": "g",
                "nutritional": {
                    "calories": 120,
                    "protein": 8.0,
                    "fat": 8.0,
                    "carbs": 1.0
                }
            },
            {
                "name": "Olive Oil",
                "amount": 1,
                "unit": "tbsp",
                "nutritional": {
                    "calories": 120,
                    "protein": 0,
                    "fat": 14.0,
                    "carbs": 0
                }
            }
        ],
        "nutritional_content": {
            "calories": 993,
            "protein": 25.0,
            "fat": 25.3,
            "carbs": 165.5
        },
        "category": "Main Dish",
        "meal_type": [
            "Lunch",
            "Dinner"
        ],
        "cuisine": "Italian",
        "tags": [
            "Vegetarian",
            "Comfort Food",
            "Rich"
        ],
        "health_scores": {
            "heart_health": 60,
            "diabetes_friendly": 55,
            "weight_loss": 40
        },
        "preparation_time": 40,
        "difficulty": "Hard",
        "serving_size": 3,
        "instructions": "Cook rice with mushrooms and Parmesan, stirring constantly until creamy."
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
        "name": "Chickpea Curry",
        "ingredients": [
            {
                "name": "Chickpeas",
                "amount": 200,
                "unit": "g",
                "nutritional": {
                    "calories": 364,
                    "protein": 19.0,
                    "fat": 6.0,
                    "carbs": 61.0
                }
            },
            {
                "name": "Coconut Milk",
                "amount": 100,
                "unit": "ml",
                "nutritional": {
                    "calories": 200,
                    "protein": 2.0,
                    "fat": 20.0,
                    "carbs": 5.0
                }
            },
            {
                "name": "Onion",
                "amount": 80,
                "unit": "g",
                "nutritional": {
                    "calories": 32,
                    "protein": 0.9,
                    "fat": 0.1,
                    "carbs": 7.6
                }
            },
            {
                "name": "Curry Powder",
                "amount": 1,
                "unit": "tbsp",
                "nutritional": {
                    "calories": 20,
                    "protein": 1.0,
                    "fat": 0.5,
                    "carbs": 4.0
                }
            }
        ],
        "nutritional_content": {
            "calories": 616,
            "protein": 23.0,
            "fat": 26.6,
            "carbs": 77.6
        },
        "category": "Main Dish",
        "meal_type": [
            "Lunch",
            "Dinner"
        ],
        "cuisine": "Indian",
        "tags": [
            "Vegan",
            "High Protein",
            "Gluten-Free"
        ],
        "health_scores": {
            "heart_health": 80,
            "diabetes_friendly": 75,
            "weight_loss": 70
        },
        "preparation_time": 30,
        "difficulty": "Medium",
        "serving_size": 2,
        "instructions": "Simmer chickpeas with coconut milk, onion, and curry powder. Serve with rice."
    },
    {
        "name": "Mushroom Risotto",
        "ingredients": [
            {
                "name": "Arborio Rice",
                "amount": 200,
                "unit": "g",
                "nutritional": {
                    "calories": 720,
                    "protein": 13.0,
                    "fat": 3.0,
                    "carbs": 158.0
                }
            },
            {
                "name": "Mushrooms",
                "amount": 150,
                "unit": "g",
                "nutritional": {
                    "calories": 33,
                    "protein": 4.0,
                    "fat": 0.3,
                    "carbs": 6.5
                }
            },
            {
                "name": "Parmesan",
                "amount": 30,
                "unit": "g",
                "nutritional": {
                    "calories": 120,
                    "protein": 8.0,
                    "fat": 8.0,
                    "carbs": 1.0
                }
            },
            {
                "name": "Olive Oil",
                "amount": 1,
                "unit": "tbsp",
                "nutritional": {
                    "calories": 120,
                    "protein": 0,
                    "fat": 14.0,
                    "carbs": 0
                }
            }
        ],
        "nutritional_content": {
            "calories": 993,
            "protein": 25.0,
            "fat": 25.3,
            "carbs": 165.5
        },
        "category": "Main Dish",
        "meal_type": [
            "Lunch",
            "Dinner"
        ],
        "cuisine": "Italian",
        "tags": [
            "Vegetarian",
            "Comfort Food",
            "Rich"
        ],
        "health_scores": {
            "heart_health": 60,
            "diabetes_friendly": 55,
            "weight_loss": 40
        },
        "preparation_time": 40,
        "difficulty": "Hard",
        "serving_size": 3,
        "instructions": "Cook rice with mushrooms and Parmesan, stirring constantly until creamy."
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
        "name": "Lentil Soup",
        "category": "Soup"
    },
    {
        "id": 4,
        "name": "Quinoa Salad",
        "category": "Salad"
    },
    {
        "id": 3,
        "name": "Vegetable Stir Fry",
        "category": "Main Dish"
    }
]
```

> GET /api/meals/{mealId}

Description: Get a specific meal's information by mealId.

### Response:

```json
{
    "id": 12,
    "name": "Oatmeal",
    "ingredients": [
        {
            "id": 45,
            "meal_id": 12,
            "name": "Oats",
            "amount": 100,
            "unit": "g",
            "nutritional": {
                "fat": 6.9,
                "carbs": 66.3,
                "protein": 16.9,
                "calories": 389
            }
        },
        {
            "id": 46,
            "meal_id": 12,
            "name": "Milk",
            "amount": 200,
            "unit": "ml",
            "nutritional": {
                "fat": 4.5,
                "carbs": 12.0,
                "protein": 8.0,
                "calories": 122
            }
        },
        {
            "id": 47,
            "meal_id": 12,
            "name": "Banana",
            "amount": 80,
            "unit": "g",
            "nutritional": {
                "fat": 0.2,
                "carbs": 18.5,
                "protein": 0.9,
                "calories": 72
            }
        }
    ],
    "nutritional_content": {
        "fat": 11.6,
        "carbs": 96.8,
        "protein": 25.8,
        "calories": 583
    },
    "category": "Breakfast",
    "meal_type": [
        "Breakfast"
    ],
    "cuisine": "American",
    "tags": [
        "Vegetarian",
        "High Fiber",
        "Easy"
    ],
    "health_scores": {
        "weight_loss": 75,
        "heart_health": 90,
        "diabetes_friendly": 80
    },
    "preparation_time": 10,
    "difficulty": "Easy",
    "serving_size": 1,
    "instructions": "Cook oats in milk, top with banana slices. Serve warm."
}
```

> PUT /api/meals/{mealId}

Description: Update a specific meal's information by mealId.

#### Request Body:

```json
{
    "name": "Oatmeal",
    "ingredients": [
        {
            "id": 45,
            "meal_id": 12,
            "name": "Oats",
            "amount": 100,
            "unit": "g",
            "nutritional": {
                "fat": 6.9,
                "carbs": 66.3,
                "protein": 16.9,
                "calories": 389
            }
        },
        {
            "id": 46,
            "meal_id": 12,
            "name": "Milk",
            "amount": 200,
            "unit": "ml",
            "nutritional": {
                "fat": 4.5,
                "carbs": 12.0,
                "protein": 8.0,
                "calories": 122
            }
        },
        {
            "id": 47,
            "meal_id": 12,
            "name": "Banana",
            "amount": 80,
            "unit": "g",
            "nutritional": {
                "fat": 0.2,
                "carbs": 18.5,
                "protein": 0.9,
                "calories": 72
            }
        }
    ],
    "nutritional_content": {
        "fat": 11.6,
        "carbs": 96.8,
        "protein": 25.8,
        "calories": 583
    },
    "category": "Breakfast",
    "meal_type": [
        "Breakfast"
    ],
    "cuisine": "American",
    "tags": [
        "Vegetarian",
        "High Fiber",
        "Easy"
    ],
    "health_scores": {
        "weight_loss": 75,
        "heart_health": 90,
        "diabetes_friendly": 80
    },
    "preparation_time": 10,
    "difficulty": "Easy",
    "serving_size": 1,
    "instructions": "Cook oats in milk, top with banana slices. Serve warm."
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