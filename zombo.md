#### 1. **After installation, load the extension into your database**

```sql
CREATE EXTENSION zombodb;
```

#### 2. **Create Your Database and Tables**

Based on your OpenAPI schema, you'll need to create tables in PostgreSQL. For simplicity, we'll create the main tables for `UserHealthInfo` and `DietaryPreferences` as per your schema.

```sql
CREATE DATABASE nutrition_app;

\c nutrition_app  -- Connect to the database

-- Create UserHealthInfo table
CREATE TABLE user_health_info (
    user_id SERIAL PRIMARY KEY,
    age INTEGER,
    gender_assigned_at_birth TEXT,
    current_gender_identity TEXT,
    hormone_use TEXT,
    height NUMERIC,
    weight NUMERIC,
    body_type TEXT,
    body_fat_percentage NUMERIC,
    skin_tone TEXT,
    ethnic_background TEXT,
    activity_level TEXT,
    meal_timing_preferences TEXT,
    fluid_intake_water NUMERIC,
    fluid_intake_other TEXT,
    seasonal_dietary_changes TEXT,
    travel_frequency TEXT
);

-- Create DietaryPreferences table
CREATE TABLE dietary_preferences (
    user_id INTEGER PRIMARY KEY,
    dietary_preference TEXT
);
```

#### 3. **Configure Elasticsearch**

Ensure Elasticsearch is running and accessible. You can check if Elasticsearch is running by accessing `http://localhost:9200` in your browser or using curl:

```bash
curl -X GET "localhost:9200/"
```

You should receive a response that confirms Elasticsearch is up and running.

#### 4. **Create ZomboDB Indexes**

ZomboDB allows you to create indexes that link PostgreSQL tables with Elasticsearch indices. Let's create a ZomboDB index for the `user_health_info` table to enable full-text search.

```sql
-- Create ZomboDB index for user_health_info
CREATE INDEX idx_user_health_info
ON user_health_info
USING zombodb ((user_health_info.*))
WITH (url='http://localhost:9200/');
```

In this example, the `user_health_info.*` syntax tells ZomboDB to index all columns of the `user_health_info` table.

#### 5. **Insert Sample Data**

Insert some sample data into your tables to test the setup:

```sql
INSERT INTO user_health_info (age, gender_assigned_at_birth, current_gender_identity, hormone_use, height, weight, body_type, body_fat_percentage, skin_tone, ethnic_background, activity_level, meal_timing_preferences, fluid_intake_water, fluid_intake_other, seasonal_dietary_changes, travel_frequency)
VALUES (25, 'female', 'female', 'none', 170, 65, 'average', 20, 'medium', 'Asian', 'active', 'regular', 2.0, 'juice', 'none', 'frequent');

INSERT INTO dietary_preferences (user_id, dietary_preference)
VALUES (1, 'vegetarian');
```

#### 6. **Querying with ZomboDB**

Now that we have data and indexes set up, we can perform some searches.

- **Full-Text Search:**

```sql
-- Search for users with the word 'active' in their activity_level
SELECT * FROM user_health_info WHERE user_health_info ==> 'active';
```

This query will search the `user_health_info` table for any rows where the `activity_level` field contains the word 'active'.

- **Advanced Search:**

You can also perform more complex searches. For example, let's find users with a specific dietary preference and a certain activity level:

```sql
-- Join tables and perform a search
SELECT uhi.*
FROM user_health_info uhi
JOIN dietary_preferences dp ON uhi.user_id = dp.user_id
WHERE uhi ==> 'active' AND dp.dietary_preference = 'vegetarian';
```

This query joins the `user_health_info` and `dietary_preferences` tables, looking for users who are both 'active' and 'vegetarian'.

#### 7. **Updating Data**

ZomboDB automatically updates the Elasticsearch index whenever data in PostgreSQL is modified. For example, updating a user’s health information:

```sql
UPDATE user_health_info
SET activity_level = 'sedentary'
WHERE user_id = 1;
```

#### 8. **Deleting Data**

When you delete data from a table indexed by ZomboDB, the corresponding documents in Elasticsearch are also deleted:

```sql
DELETE FROM user_health_info WHERE user_id = 1;
```

### Troubleshooting and Tips

- **Check Elasticsearch Logs**: If something doesn’t work as expected, check Elasticsearch logs for any errors.
- **Use PostgreSQL Logs**: PostgreSQL logs can also provide insight into any errors related to ZomboDB.
- **Elasticsearch URL**: Ensure that the URL provided in the ZomboDB index creation matches your Elasticsearch server URL.
- **Data Types**: Ensure that PostgreSQL data types are compatible with Elasticsearch indexing.

