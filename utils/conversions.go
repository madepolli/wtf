package utils

import (
	"strconv"
)

/* -------------------- Map Conversion -------------------- */

// MapToStrs takes a map of interfaces and returns a map of strings
func MapToStrs(aMap map[string]interface{}) map[string]string {
	results := make(map[string]string, len(aMap))

	for key, val := range aMap {
		results[key] = val.(string)
	}

	return results
}

/* -------------------- Slice Conversion -------------------- */

// IntsToUints takes a slice of ints and returns a slice of uints
func IntsToUints(slice []int) []uint {
	results := make([]uint, len(slice))

	for i, val := range slice {
		results[i] = uint(val)
	}

	return results
}

// ToInts takes a slice of interfaces and returns a slice of ints
func ToInts(slice []interface{}) []int {
	results := make([]int, len(slice))

	for i, val := range slice {
		results[i] = val.(int)
	}

	return results
}

// ToStrs takes a slice of interfaces and returns a slice of strings
func ToStrs(slice []interface{}) []string {
	results := make([]string, len(slice))

	for i, val := range slice {
		switch t := val.(type) {
		case int:
			results[i] = strconv.Itoa(t)
		case string:
			results[i] = t
		}
	}

	return results
}

// ToUints takes a slice of interfaces and returns a slice of ints
func ToUints(slice []interface{}) []uint {
	results := make([]uint, len(slice))

	for i, val := range slice {
		results[i] = val.(uint)
	}

	return results
}
