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
    "name": "Optimus",
    "age": 50,
    "gender": "Male",
    "body_metrics": {
        "user_id": 11,
        "weight": 70,
        "height": 175
    },
    "activity_level": "Moderate",
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
    "nutritional_deficiencies": [
        "Vitamin D",
        "Iron"
    ],
    "allergies": [
        "Peanuts"
    ],
    "recent_meals": [
        {
             "meal_id": 2,
            "timestamp": "2024-09-30T18:00:00Z"
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
    "blood_glucose": 90,
    "lipid_profiles": {
        "cholesterol": 0,
        "hdl": 0,
        "ldl": 0,
        "triglycerides": 0
    },
    "environmental_factors": {
        "location": "Urban",
        "climate": "Temperate",
        "season": "Fall"
    },
    "requested_meal": {
        "starter_meal_type": 1,
        "main_meal_type": 1,
        "dessert_meal_type": 1,
        "number_of_starter": 1,
        "number_of_main": 1,
        "number_of_dessert ": 1
    },
    "meal_histories": [],
    "health_score": 85,
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
    }
}
```

#####  Response:

```json
{
    "user_info": {
        "id": 13,
        "name": "Optimus",
        "age": 50,
        "gender": "Male",
        "body_metrics": {
            "user_id": 11,
            "weight": 70,
            "height": 175
        },
        "activity_level": "Moderate",
        "dietary_preferences": {
            "user_id": 0,
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
                "user_id": 0,
                "name": "Hypertension",
                "severity": "Moderate"
            }
        ],
        "nutritional_deficiencies": [
            "Vitamin D",
            "Iron"
        ],
        "allergies": [
            "Peanuts"
        ],
        "recent_meals": [
            {
                "id": 0,
                "user_id": 0,
                "meal_id": 2,
                "timestamp": "2024-09-30T18:00:00Z"
            }
        ],
        "goals": [
            {
                "id": 0,
                "user_id": 0,
                "type": "Weight loss",
                "target": 5,
                "duration": 90
            }
        ],
        "microbiome_data": {
            "user_id": 0,
            "diversity_score": 75,
            "gut_health_recommendations": [
                "Increase fiber intake"
            ]
        },
        "blood_glucose": 90,
        "lipid_profiles": {
            "user_id": 0,
            "cholesterol": 0,
            "hdl": 0,
            "ldl": 0,
            "triglycerides": 0
        },
        "environmental_factors": {
            "user_id": 0,
            "location": "Urban",
            "climate": "Temperate",
            "season": "Fall"
        },
        "requested_meal": {
            "user_id": 0,
            "starter_meal_type": 1,
            "main_meal_type": 1,
            "dessert_meal_type": 1,
            "number_of_starter": 1,
            "number_of_main": 1,
            "number_of_dessert ": 1
        },
        "meal_histories": [],
        "health_score": 85,
        "preferences": {
            "user_id": 0,
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
        }
    },
    "recommendations": [
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
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
        "id": 11,
        "name": "Ahmad",
        "age": 30,
        "gender": "Male",
        "body_metrics": {
            "user_id": 11,
            "weight": 70,
            "height": 175
        },
        "activity_level": "Moderate",
        "dietary_preferences": {
            "user_id": 11,
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
                "user_id": 11,
                "name": "Hypertension",
                "severity": "Moderate"
            }
        ],
        "nutritional_deficiencies": [
            "Vitamin D",
            "Iron"
        ],
        "allergies": [
            "Peanuts"
        ],
        "recent_meals": [
            {
                "id": 130,
                "user_id": 11,
                "meal_id": 2,
                "timestamp": "2024-09-30T18:00:00Z"
            }
        ],
        "goals": [
            {
                "id": 132,
                "user_id": 11,
                "type": "Weight loss",
                "target": 5,
                "duration": 90
            }
        ],
        "microbiome_data": {
            "user_id": 11,
            "diversity_score": 75,
            "gut_health_recommendations": [
                "Increase fiber intake"
            ]
        },
        "blood_glucose": 90,
        "lipid_profiles": {
            "user_id": 11,
            "cholesterol": 0,
            "hdl": 0,
            "ldl": 0,
            "triglycerides": 0
        },
        "environmental_factors": {
            "user_id": 11,
            "location": "Urban",
            "climate": "Temperate",
            "season": "Fall"
        },
        "requested_meal": {
            "user_id": 11,
            "starter_meal_type": 1,
            "main_meal_type": 1,
            "dessert_meal_type": 1,
            "number_of_starter": 1,
            "number_of_main": 1,
            "number_of_dessert ": 1
        },
        "meal_histories": [],
        "health_score": 85,
        "preferences": {
            "user_id": 11,
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
        }
    },
    "recommendations": [
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        }
    ]
}
```



> PUT /api/user-health-info/{userId}

Description: Update a specific user's health information by userId.

##### Request Body:

```json
  {
        "id": 11,
        "name": "Ahmad",
        "age": 30,
        "gender": "Male",
        "body_metrics": {
            "user_id": 11,
            "weight": 70,
            "height": 175
        },
        "activity_level": "Moderate",
        "dietary_preferences": {
            "user_id": 11,
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
                "user_id": 11,
                "name": "Hypertension",
                "severity": "Moderate"
            }
        ],
        "nutritional_deficiencies": [
            "Vitamin D",
            "Iron"
        ],
        "allergies": [
            "Peanuts"
        ],
        "recent_meals": [
            {
                "id": 130,
                "user_id": 11,
                "meal_id": 2,
                "timestamp": "2024-09-30T18:00:00Z"
            }
        ],
        "goals": [
            {
                "id": 132,
                "user_id": 11,
                "type": "Weight loss",
                "target": 5,
                "duration": 90
            }
        ],
        "microbiome_data": {
            "user_id": 11,
            "diversity_score": 75,
            "gut_health_recommendations": [
                "Increase fiber intake"
            ]
        },
        "blood_glucose": 90,
        "lipid_profiles": {
            "user_id": 11,
            "cholesterol": 0,
            "hdl": 0,
            "ldl": 0,
            "triglycerides": 0
        },
        "environmental_factors": {
            "user_id": 11,
            "location": "Urban",
            "climate": "Temperate",
            "season": "Fall"
        },
        "requested_meal": {
            "user_id": 11,
            "starter_meal_type": 1,
            "main_meal_type": 1,
            "dessert_meal_type": 1,
            "number_of_starter": 1,
            "number_of_main": 1,
            "number_of_dessert ": 1
        },
        "meal_histories": [],
        "health_score": 85,
        "preferences": {
            "user_id": 11,
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
        }
    }
```

#### Response:

```json
{
    "user_info": {
        "id": 11,
        "name": "Ahmad",
        "age": 30,
        "gender": "Male",
        "body_metrics": {
            "user_id": 11,
            "weight": 70,
            "height": 175
        },
        "activity_level": "Moderate",
        "dietary_preferences": {
            "user_id": 0,
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
                "user_id": 0,
                "name": "Hypertension",
                "severity": "Moderate"
            }
        ],
        "nutritional_deficiencies": [
            "Vitamin D",
            "Iron"
        ],
        "allergies": [
            "Peanuts"
        ],
        "recent_meals": [
            {
                "id": 0,
                "user_id": 0,
                "meal_id": 2,
                "timestamp": "2024-09-30T18:00:00Z"
            }
        ],
        "goals": [
            {
                "id": 0,
                "user_id": 0,
                "type": "Weight loss",
                "target": 5,
                "duration": 90
            }
        ],
        "microbiome_data": {
            "user_id": 0,
            "diversity_score": 75,
            "gut_health_recommendations": [
                "Increase fiber intake"
            ]
        },
        "blood_glucose": 90,
        "lipid_profiles": {
            "user_id": 0,
            "cholesterol": 0,
            "hdl": 0,
            "ldl": 0,
            "triglycerides": 0
        },
        "environmental_factors": {
            "user_id": 0,
            "location": "Urban",
            "climate": "Temperate",
            "season": "Fall"
        },
        "requested_meal": {
            "user_id": 0,
            "starter_meal_type": 1,
            "main_meal_type": 1,
            "dessert_meal_type": 1,
            "number_of_starter": 1,
            "number_of_main": 1,
            "number_of_dessert ": 1
        },
        "meal_histories": [],
        "health_score": 85,
        "preferences": {
            "user_id": 0,
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
        }
    },
    "recommendations": [
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 12,
                    "course": "Starter",
                    "meal_name": "Stuffed Mushrooms",
                    "ingredients": [
                        {
                            "name": "Mushrooms",
                            "amount": 150,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Cream Cheese",
                            "amount": 50,
                            "unit": "grams",
                            "portion": "2 tbsp",
                            "ounces": "1.76"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 1,
                    "course": "Main",
                    "meal_name": "Grilled Chicken Salad",
                    "ingredients": [
                        {
                            "name": "Chicken Breast",
                            "amount": 150,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "5.29"
                        },
                        {
                            "name": "Mixed Greens",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "2 cups",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 68.18
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 9,
                    "course": "Main",
                    "meal_name": "Lentil Soup",
                    "ingredients": [
                        {
                            "name": "Lentils",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Carrots",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 6,
                    "course": "Main",
                    "meal_name": "Vegan Buddha Bowl",
                    "ingredients": [
                        {
                            "name": "Brown Rice",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Chickpeas",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "0.25 cup",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Tahini",
                            "amount": 20,
                            "unit": "Grams",
                            "portion": "1 tbsp",
                            "ounces": "0.71"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 17,
                    "course": "Dessert",
                    "meal_name": "Tiramisu",
                    "ingredients": [
                        {
                            "name": "Mascarpone Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "1 pcs",
                            "ounces": "8.82"
                        },
                        {
                            "name": "Ladyfingers",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "2",
                            "ounces": "7.05"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 16,
                    "course": "Dessert",
                    "meal_name": "Apple Pie",
                    "ingredients": [
                        {
                            "name": "Apples",
                            "amount": 4,
                            "unit": "pieces",
                            "portion": "Sliced",
                            "ounces": "0.14"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 15,
                    "course": "Dessert",
                    "meal_name": "Cheesecake",
                    "ingredients": [
                        {
                            "name": "Cream Cheese",
                            "amount": 250,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "8.82"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 14,
                    "course": "Dessert",
                    "meal_name": "Fruit Salad",
                    "ingredients": [
                        {
                            "name": "Mixed Fruits",
                            "amount": 300,
                            "unit": "grams",
                            "portion": "1/2",
                            "ounces": "10.58"
                        }
                    ],
                    "score": 45.45
                }
            ]
        },
        {
            "courses": [
                {
                    "meal_id": 11,
                    "course": "Starter",
                    "meal_name": "Bruschetta",
                    "ingredients": [
                        {
                            "name": "Tomato",
                            "amount": 100,
                            "unit": "grams",
                            "portion": "2 pcs",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Basil",
                            "amount": 10,
                            "unit": "grams",
                            "portion": "1",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 3,
                    "course": "Main",
                    "meal_name": "Quinoa Salad",
                    "ingredients": [
                        {
                            "name": "Quinoa",
                            "amount": 100,
                            "unit": "Grams",
                            "portion": "0.5 cup",
                            "ounces": "3.53"
                        },
                        {
                            "name": "Cucumber",
                            "amount": 50,
                            "unit": "Grams",
                            "portion": "1 pcs",
                            "ounces": "1.76"
                        },
                        {
                            "name": "Olive Oil",
                            "amount": 10,
                            "unit": "ML",
                            "portion": "1 tbsp",
                            "ounces": "0.35"
                        }
                    ],
                    "score": 45.45
                },
                {
                    "meal_id": 13,
                    "course": "Dessert",
                    "meal_name": "Chocolate Mousse",
                    "ingredients": [
                        {
                            "name": "Dark Chocolate",
                            "amount": 200,
                            "unit": "grams",
                            "portion": "4 pcs",
                            "ounces": "7.05"
                        },
                        {
                            "name": "Eggs",
                            "amount": 3,
                            "unit": "pieces",
                            "portion": "3 pcs",
                            "ounces": "0.11"
                        }
                    ],
                    "score": 45.45
                }
            ]
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





