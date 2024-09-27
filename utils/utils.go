package utils

import (
	"encoding/json"
	"math"
	"reflect"
	"strconv"

	"gorm.io/datatypes"
)

// Helper function to compare JSON fields
func CompareJSONFields(a, b datatypes.JSON) bool {
	var aMap, bMap map[string]interface{}
	_ = json.Unmarshal(a, &aMap)
	_ = json.Unmarshal(b, &bMap)
	return reflect.DeepEqual(aMap, bMap)
}

// Helper function to determine the course type based on the number of courses
func GetCourseType(numCourses int) string {
	switch numCourses {
	case 1:
		return "single-course"
	case 2:
		return "2-course"
	case 3:
		return "3-course"
	default:
		return "multi-course" // For any other number of courses
	}
}

func GramsToOunces(grams float64) float64 {
	ounces := grams / 28.35
	return math.Round(ounces*100) / 100
}

func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
