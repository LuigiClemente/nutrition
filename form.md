#### Section 1: Basic Health Information

1. **Age**  
   - **Field Type**: Integer  
   - **Description**: User's age in years, relevant for caloric and nutritional recommendations.

2. **Gender Assigned at Birth**  
   - **Field Type**: Enum  
   - **Options**:  
     - Male  
     - Female  
     - Prefer Not to Say  
   - **Description**: Biological gender to assess specific nutritional needs and caloric requirements.

3. **Current Gender Identity**  
   - **Field Type**: Enum  
   - **Options**:  
     - Male  
     - Female  
     - Non-Binary  
     - Prefer Not to Say  
   - **Description**: User's current gender identity to provide more personalized recommendations.

4. **Hormone Use**  
   - **Field Type**: Enum  
   - **Options**:  
     - No  
     - Yes, Estrogen-based (e.g., HRT)  
     - Yes, Testosterone-based (e.g., TRT)  
     - Other (with a text field to specify)  
   - **Description**: Hormone use can affect metabolism and nutrient needs.

5. **Height**  
   - **Field Type**: Float  
   - **Description**: User's height in centimeters or inches, used to calculate BMI and caloric needs.

6. **Weight**  
   - **Field Type**: Float  
   - **Description**: User's weight in kilograms or pounds, used to determine BMI and caloric requirements.

7. **Body Type**  
   - **Field Type**: Enum  
   - **Options**:  
     - Ectomorph (lean and long)  
     - Mesomorph (muscular and well-built)  
     - Endomorph (higher body fat percentage)  
   - **Description**: Helps refine caloric and macronutrient distribution.

8. **Body Fat Percentage**  
   - **Field Type**: Float  
   - **Description**: User's estimated or measured body fat percentage for a more accurate caloric and macronutrient assessment.

9. **Skin Tone**  
   - **Field Type**: Enum  
   - **Options**:  
     - Very Light  
     - Light  
     - Medium  
     - Olive  
     - Dark  
     - Very Dark  
   - **Description**: May affect vitamin D synthesis and related nutritional needs.

10. **Ethnic Background**  
    - **Field Type**: Enum  
    - **Options**:  
      - African  
      - Asian  
      - Caucasian  
      - Hispanic  
      - Middle Eastern  
      - Native American  
      - Pacific Islander  
      - Mixed  
      - Other (with a text field to specify)  
    - **Description**: Used to better understand potential dietary preferences and nutritional needs.

11. **Activity Level**  
    - **Field Type**: Enum  
    - **Options**:  
      - Sedentary (little or no exercise)  
      - Lightly active (light exercise/sports 1-3 days/week)  
      - Moderately active (moderate exercise/sports 3-5 days/week)  
      - Very active (hard exercise/sports 6-7 days a week)  
      - Super active (very hard exercise, physical job, or training twice a day)  
    - **Description**: Provides information on caloric expenditure to tailor meal recommendations appropriately.

12. **Exercise Types**  
    - **Field Type**: Repeated Enum  
    - **Options**:  
      - Cardio (e.g., running, cycling)  
      - Strength Training (e.g., weightlifting)  
      - Flexibility (e.g., yoga)  
      - Mixed (e.g., HIIT, CrossFit)  
      - Other (with a text field to specify)  
    - **Description**: Types of physical activities performed regularly, used to estimate caloric needs.

13. **Meal Timing Preferences**  
    - **Field Type**: Enum  
    - **Options**:  
      - Regular Meals (3 meals a day)  
      - Small Frequent Meals (5-6 meals a day)  
      - Intermittent Fasting  
      - Other (with a text field to specify)  
    - **Description**: Helps customize meal plans to match user eating patterns.

14. **Fluid Intake**  
    - **Field Type**: Message  
    - **Sub-fields**:  
      - Water (Float: Liters per day)  
      - Other Beverages (Enum: Coffee, Tea, Soda, Alcohol, Other)  
    - **Description**: Provides insight into hydration levels and beverage choices affecting nutrition.

15. **Seasonal Dietary Changes**  
    - **Field Type**: Enum  
    - **Options**:  
      - None  
      - Lighter in Summer  
      - Heavier in Winter  
      - Other (with a text field to specify)  
    - **Description**: Indicates any seasonal variations in dietary habits.

16. **Travel Frequency**  
    - **Field Type**: Enum  
    - **Options**:  
      - Rarely  
      - Occasionally  
      - Frequently  
      - Constantly  
    - **Description**: Indicates how often the user travels, which may affect meal consistency and planning.

#### Section 2: Dietary Preferences and Restrictions

1. **Dietary Preferences**  
   - **Field Type**: Repeated Enum  
   - **Options**:  
     - Vegetarian  
     - Vegan  
     - Pescatarian  
     - Keto  
     - Paleo  
     - Low-Carb  
     - Low-Fat  
     - Mediterranean  
     - DASH (Dietary Approaches to Stop Hypertension)  
     - Whole30  
     - Carnivore  
     - Plant-Based  
     - Flexitarian  
     - No Preference  
   - **Description**: User’s dietary preferences for personalized meal planning.

2. **Food Allergies and Intolerances**  
   - **Field Type**: Repeated Enum  
   - **Options**:  
     - Gluten  
     - Dairy  
     - Nuts (specify type if possible)  
     - Soy  
     - Eggs  
     - Shellfish  
     - Nightshades (tomatoes, peppers, etc.)  
     - FODMAPs (Fermentable Oligo-, Di-, Mono-saccharides And Polyols)  
     - Histamine Intolerance  
     - Sulfite Sensitivity  
     - Other (with a text field to specify)  
   - **Description**: Allergies and intolerances to avoid in meal recommendations.

3. **Specific Food Avoidances**  
   - **Field Type**: Repeated Enum  
   - **Options**:  
     - Red Meat  
     - Processed Foods  
     - Added Sugars  
     - Refined Grains  
     - Artificial Sweeteners  
     - Caffeine  
     - Alcohol  
     - High-FODMAP Foods  
     - Spicy Foods  
     - GMOs  
     - Raw Foods  
     - Unpasteurized Products  
     - Other (with a text field to specify)  
   - **Description**: Foods the user prefers to avoid in their diet.

4. **Nutritional Deficiencies**  
   - **Field Type**: Repeated Enum  
   - **Options**:  
     - Iron Deficiency  
     - Vitamin D Deficiency  
     - Calcium Deficiency  
     - Vitamin B12 Deficiency  
     - Magnesium Deficiency  
     - Other (with a text field to specify)  
   - **Description**: Known deficiencies to address in meal planning.

5. **Frequency of Consuming Specific Foods**  
   - **Field Type**: Repeated Message  
   - **Sub-fields**:  
     - Food Name (String)  
     - Frequency (Enum: Daily, Weekly, Fortnightly, Monthly, Quarterly)  
   - **Description**: Tracks how often the user consumes specific foods or food groups.

6. **Cooking Habits and Skills**  
   - **Field Type**: Enum  
   - **Options**:  
     - Rarely Cooks  
     - Occasionally Cooks  
     - Regularly Cooks  
     - Expert in Cooking  
   - **Description**: User's cooking habits and skill level for adjusting meal recommendations.

7. **Meal Preparation Preferences**  
   - **Field Type**: Enum  
   - **Options**:  
     - Batch Cooking  
     - Daily Preparation  
     - Meal Kits  
     - Meal Delivery  
     - Other (with a text field to specify)  
   - **Description**: User's preferred method for meal preparation.

8. **Preferred Ingredients**  
   - **Field Type**: Repeated String  
   - **Description**: Ingredients or foods that the user enjoys and prefers to include in their diet.

9. **Disliked Ingredients**  
   - **Field Type**: Repeated String  
   - **Description**: Ingredients or foods that the user dislikes or prefers to exclude from their diet.

10. **Favorite Cuisines**  
    - **Field Type**: Repeated Enum  
    - **Options**:  
      - Italian  
      - Mexican  
      - Indian  
      - Chinese  
      - Japanese  
      - Mediterranean  
      - American  
      - Other (with a text field to specify)  
    - **Description**: Types of cuisines the user enjoys, to tailor meal suggestions to taste preferences.

#### Section 3

: Health Conditions and Goals

1. **Existing Health Conditions**  
   - **Field Type**: Repeated Enum  
   - **Options**:  
     - Diabetes  
     - High Cholesterol  
     - Hypertension  
     - Thyroid Disorders  
     - Celiac Disease  
     - IBS (Irritable Bowel Syndrome)  
     - Lactose Intolerance  
     - Cardiovascular Diseases  
     - Obesity  
     - PCOS (Polycystic Ovary Syndrome)  
     - Osteoporosis  
     - Anemia  
     - Hereditary Diseases (with a text field to specify)  
     - Autoimmune Disorders (e.g., Rheumatoid Arthritis, Lupus)  
     - Cancer (specify type if possible)  
     - Other (with a text field to specify)  
   - **Description**: Health conditions that may influence meal recommendations and caloric needs.

2. **Health and Nutrition Goals**  
   - **Field Type**: Repeated Enum  
   - **Options**:  
     - Weight Loss  
     - Muscle Gain  
     - Improve Energy Levels  
     - Enhance Athletic Performance  
     - Manage Blood Sugar Levels  
     - Lower Cholesterol  
     - Improve Gut Health  
     - Bone Health Improvement  
     - Immune System Support  
     - General Wellness  
     - Reducing Inflammation  
     - Improving Digestion  
     - Other (with a text field to specify)  
   - **Description**: User's primary health goals for tailored meal planning.

3. **Medical Supervision**  
   - **Field Type**: Enum  
   - **Options**:  
     - Yes  
     - No  
   - **Description**: Indicates if the user is under medical supervision for their diet.

4. **Family Medical History**  
   - **Field Type**: Repeated Enum  
   - **Options**:  
     - Diabetes  
     - Heart Disease  
     - Hypertension  
     - Cancer  
     - Other (with a text field to specify)  
   - **Description**: Relevant family medical history that could impact nutritional needs.

5. **Eating Behavior**  
   - **Field Type**: Enum  
   - **Options**:  
     - Slow Eater  
     - Fast Eater  
     - Emotional Eater  
     - Night Eater  
     - Mindful Eater  
     - Intuitive Eater  
     - Other (with a text field to specify)  
   - **Description**: User's eating behaviors that can influence dietary recommendations.

6. **Food Addiction Concerns**  
   - **Field Type**: Enum  
   - **Options**:  
     - Yes  
     - No  
   - **Description**: Indicates if the user has concerns about food addiction, which might influence recommendations.

7. **Adverse Reactions to Foods**  
   - **Field Type**: Repeated String  
   - **Description**: Non-allergic adverse reactions to specific foods (e.g., acid reflux, bloating).

8. **Health Monitoring and Technology**  
   - **Field Type**: Enum  
   - **Options**:  
     - No Devices Used  
     - Fitness Tracker (e.g., Fitbit, Garmin)  
     - Smartwatch (e.g., Apple Watch, Samsung Galaxy Watch)  
     - Health Apps (e.g., MyFitnessPal, Lose It!)  
     - Other (with a text field to specify)  
   - **Description**: Information on health tracking devices or apps the user utilizes.

#### Section 4: Historical Eating Patterns and Meal Logs

1. **Typical Daily Meals**  
   - **Field Type**: Repeated Message  
   - **Sub-fields**:  
     - Meal Time (Enum: Breakfast, Lunch, Dinner, Snacks)  
     - Food Items (String: Detailed description of typical meals)  
   - **Description**: Details of the user's typical daily meals for pattern recognition.

2. **Average Daily Caloric Intake**  
   - **Field Type**: Integer  
   - **Description**: User's estimated daily caloric intake, used to better tailor meal recommendations.

3. **Frequency of Eating Out**  
   - **Field Type**: Enum  
   - **Options**:  
     - Never  
     - Rarely (1-2 times a month)  
     - Occasionally (1-2 times a week)  
     - Frequently (3-4 times a week)  
     - Very Frequently (5 or more times a week)  
   - **Description**: How often the user eats meals prepared outside the home.

4. **Snacking Habits**  
   - **Field Type**: Enum  
   - **Options**:  
     - No Snacks  
     - Occasional Snacking (1-2 times a week)  
     - Regular Snacking (daily)  
   - **Description**: User's snacking frequency for additional caloric intake consideration.

5. **Alcohol Consumption**  
   - **Field Type**: Enum  
   - **Options**:  
     - None  
     - Occasional (1-2 times a month)  
     - Moderate (1-2 times a week)  
     - Frequent (3-4 times a week)  
     - Heavy (daily or almost daily)  
   - **Description**: Information about alcohol consumption habits that may impact nutrition.

6. **Smoking Status**  
   - **Field Type**: Enum  
   - **Options**:  
     - Never Smoked  
     - Former Smoker  
     - Occasional Smoker  
     - Regular Smoker  
   - **Description**: Information about smoking habits that may influence nutritional needs and overall health.

7. **Recent Diet Changes**  
   - **Field Type**: Enum  
   - **Options**:  
     - Yes  
     - No  
   - **Description**: Indicates if the user has recently made any significant diet changes.

8. **Feedback on Past Diets**  
   - **Field Type**: String  
   - **Description**: Option to provide feedback on previous diets and their effects to guide more effective future recommendations.

9. **Eating Schedule Flexibility**  
   - **Field Type**: Enum  
   - **Options**:  
     - Very Rigid  
     - Somewhat Rigid  
     - Flexible  
     - Very Flexible  
   - **Description**: Understanding the user’s eating schedule flexibility to accommodate spontaneous or rigid meal timing.

#### Section 5: Health Test Results (Collected via Protobuf)

1. **Microbiome Data**  
   - **Field Type**: Message  
   - **Sub-fields**:  
     - Microbiome Diversity Score (Float)  
     - Specific Bacteria Levels (Repeated Float)  
   - **Description**: Data related to the user's gut health for personalized nutrition suggestions.

2. **Continuous Glucose Monitoring (CGM) Data**  
   - **Field Type**: Message  
   - **Sub-fields**:  
     - Average Blood Glucose Levels (Float)  
     - Glucose Variability (Float)  
   - **Description**: Glucose data to understand sugar management needs.

3. **Lipid Profile Data**  
   - **Field Type**: Message  
   - **Sub-fields**:  
     - Total Cholesterol (Float)  
     - HDL (Float)  
     - LDL (Float)  
     - Triglycerides (Float)  
   - **Description**: Lipid profile details to tailor fat intake in recommendations.

4. **Other Relevant Health Tests**  
   - **Field Type**: Repeated Message  
   - **Sub-fields**:  
     - Test Name (String)  
     - Test Value (Float)  
     - Test Units (String)  
   - **Description**: Additional health metrics relevant to diet and nutrition.

#### Section 6: Environmental Factors

1. **Current Location**  
   - **Field Type**: String  
   - **Description**: User's current geographical location to assess local dietary options and weather impact.

2. **Current Weather Data**  
   - **Field Type**: Message  
   - **Sub-fields**:  
     - Temperature (Float)  
     - Humidity (Float)  
     - Weather Condition (String: Clear, Cloudy, Rainy, etc.)  
   - **Description**: Weather conditions to adjust meal recommendations based on temperature and humidity.

