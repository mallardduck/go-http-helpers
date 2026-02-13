package query

import (
	"net/http"
	"strconv"
	"strings"
)

// String extracts a string value from the query parameter with the given key.
// Returns defaultValue if the key is missing or empty.
func String(r *http.Request, key string, defaultValue string) string {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultValue
	}
	return val
}

// Int extracts an integer value from the query parameter with the given key.
// Returns defaultValue if the key is missing, empty, or cannot be parsed as an int.
func Int(r *http.Request, key string, defaultValue int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return parsed
}

// Int64 extracts an int64 value from the query parameter with the given key.
// Returns defaultValue if the key is missing, empty, or cannot be parsed as an int64.
func Int64(r *http.Request, key string, defaultValue int64) int64 {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultValue
	}

	parsed, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultValue
	}
	return parsed
}

// Float64 extracts a float64 value from the query parameter with the given key.
// Returns defaultValue if the key is missing, empty, or cannot be parsed as a float64.
func Float64(r *http.Request, key string, defaultValue float64) float64 {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultValue
	}

	parsed, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultValue
	}
	return parsed
}

// Bool extracts a boolean value from the query parameter with the given key.
// Returns defaultValue if the key is missing, empty, or cannot be parsed as a bool.
//
// Recognized as true (case-insensitive): "true", "1", "yes", "on", "y"
// Recognized as false (case-insensitive): "false", "0", "no", "off", "n"
func Bool(r *http.Request, key string, defaultValue bool) bool {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultValue
	}

	parsed, err := parseBool(val)
	if err != nil {
		return defaultValue
	}
	return parsed
}

// Strings extracts all values for a query parameter that appears multiple times.
// Returns an empty slice if the key is not present.
//
// Example: For URL "?tag=go&tag=rust&tag=python"
//
//	tags := query.Strings(r, "tag")  // []string{"go", "rust", "python"}
func Strings(r *http.Request, key string) []string {
	vals := r.URL.Query()[key]
	if vals == nil {
		return []string{}
	}
	return vals
}

// Parser is a function that converts a string to type T, returning an error if conversion fails.
type Parser[T any] func(string) (T, error)

// Slice extracts all values for a query parameter and converts them using the provided parser.
// If a value cannot be parsed, the defaultValue is used for that element.
// Returns an empty slice if the key is not present.
//
// Example:
//
//	// URL: /api/items?id=1&id=2&id=invalid&id=5
//	ids := query.Slice(r, "id", 0, strconv.Atoi)
//	// Returns []int{1, 2, 0, 5}
func Slice[T any](r *http.Request, key string, defaultValue T, parser Parser[T]) []T {
	vals := r.URL.Query()[key]
	if len(vals) == 0 {
		return []T{}
	}

	result := make([]T, len(vals))
	for i, val := range vals {
		parsed, err := parser(val)
		if err != nil {
			result[i] = defaultValue
		} else {
			result[i] = parsed
		}
	}
	return result
}

// Ints extracts all integer values for a query parameter.
// Invalid values are replaced with defaultValue.
//
// Example:
//
//	// URL: /filter?id=1&id=2&id=invalid&id=5
//	ids := query.Ints(r, "id", 0)  // []int{1, 2, 0, 5}
func Ints(r *http.Request, key string, defaultValue int) []int {
	return Slice(r, key, defaultValue, strconv.Atoi)
}

// Int64s extracts all int64 values for a query parameter.
// Invalid values are replaced with defaultValue.
func Int64s(r *http.Request, key string, defaultValue int64) []int64 {
	return Slice(r, key, defaultValue, func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	})
}

// Float64s extracts all float64 values for a query parameter.
// Invalid values are replaced with defaultValue.
func Float64s(r *http.Request, key string, defaultValue float64) []float64 {
	return Slice(r, key, defaultValue, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
}

// Bools extracts all boolean values for a query parameter.
// Invalid values are replaced with defaultValue.
// Uses the same flexible parsing as Bool (true/1/yes/on, false/0/no/off).
func Bools(r *http.Request, key string, defaultValue bool) []bool {
	return Slice(r, key, defaultValue, parseBool)
}

// parseBool is the internal bool parser that can return an error.
func parseBool(s string) (bool, error) {
	lower := strings.ToLower(strings.TrimSpace(s))
	switch lower {
	case "true", "1", "yes", "on", "y":
		return true, nil
	case "false", "0", "no", "off", "n":
		return false, nil
	default:
		return false, strconv.ErrSyntax
	}
}

// Has checks if a query parameter exists (even if empty).
// Returns true if the key is present in the query string, false otherwise.
//
// Example: For URL "?search=&active"
//
//	query.Has(r, "search")   // true (empty value)
//	query.Has(r, "active")   // true (no value)
//	query.Has(r, "missing")  // false
func Has(r *http.Request, key string) bool {
	_, exists := r.URL.Query()[key]
	return exists
}

// Count returns the number of times a query parameter appears.
// Returns 0 if the parameter is not present.
//
// Example: For URL "?id=1&id=2&id=3"
//
//	query.Count(r, "id")  // 3
func Count(r *http.Request, key string) int {
	return len(r.URL.Query()[key])
}

// IsMultiple checks if a query parameter appears more than once.
// Returns false if the parameter is not present or appears only once.
//
// Example: For URL "?single=1&multi=a&multi=b"
//
//	query.IsMultiple(r, "single")  // false
//	query.IsMultiple(r, "multi")   // true
//	query.IsMultiple(r, "missing") // false
func IsMultiple(r *http.Request, key string) bool {
	return len(r.URL.Query()[key]) > 1
}

// First returns the first element of a slice, or defaultValue if the slice is empty.
// This is useful when you always use slice functions but sometimes expect single values.
//
// Example:
//
//	ids := query.Ints(r, "id", 0)
//	id := query.First(ids, 0)  // Gets first ID or 0 if none
func First[T any](slice []T, defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}
	return slice[0]
}

// All returns the entire parsed query string as a map.
// This is useful when you need to iterate over all parameters or
// when you need to extract many values and want to parse once.
//
// The returned map is a copy and can be safely modified.
func All(r *http.Request) map[string][]string {
	values := r.URL.Query()
	result := make(map[string][]string, len(values))
	for k, v := range values {
		result[k] = v
	}
	return result
}
