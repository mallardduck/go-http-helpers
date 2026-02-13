// Package query provides a type-safe API for extracting values from HTTP request
// query strings with automatic parsing and sensible defaults.
//
// # Overview
//
// Standard Go requires multiple steps to safely extract typed values from query
// parameters: checking existence, handling empty strings, parsing with strconv,
// and error handling. This package reduces that boilerplate to single-line calls
// with automatic fallback to default values.
//
// All extraction functions are panic-safe and follow a consistent pattern:
// if a parameter is missing, empty, or cannot be parsed, the provided default
// value is returned.
//
// # Basic Usage
//
// Extract typed values from query parameters:
//
//	// URL: /products?page=2&limit=50&active=true&sort=price
//	page   := query.Int(r, "page", 1)           // 2
//	limit  := query.Int(r, "limit", 25)         // 50
//	active := query.Bool(r, "active", false)    // true
//	sort   := query.String(r, "sort", "name")   // "price"
//	offset := query.Int(r, "offset", 0)         // 0 (missing key)
//
// # Multiple Values
//
// For parameters that appear multiple times (?tag=go&tag=rust&tag=python):
//
//	tags := query.Strings(r, "tag")  // []string{"go", "rust", "python"}
//
//	// Empty slice if parameter is missing
//	filters := query.Strings(r, "filter")  // []string{}
//
// # Single vs Multiple Values
//
// When you don't know if a parameter appears once or multiple times, you have options:
//
// Option 1 - Check first, then extract:
//
//	if query.IsMultiple(r, "id") {
//	    ids := query.Ints(r, "id", 0)      // Handle multiple
//	} else {
//	    id := query.Int(r, "id", 0)        // Handle single
//	}
//
// Option 2 - Always use slice functions, extract first if needed:
//
//	ids := query.Ints(r, "id", 0)           // Always returns slice
//	id := query.First(ids, 0)               // Get first value or default
//
// Option 3 - Single-value functions always get the first value:
//
//	// Works for both /api?id=1 and /api?id=1&id=2&id=3
//	id := query.Int(r, "id", 0)             // Always gets first value
//
// # Generic Slices with Type Conversion
//
// Use Slice with a parser function to convert multiple values to any type.
// Invalid values use the provided default instead of being skipped:
//
//	// URL: /api/items?id=1&id=2&id=invalid&id=5
//	ids := query.Slice(r, "id", 0, strconv.Atoi)
//	// Returns []int{1, 2, 0, 5} - "invalid" becomes 0
//
//	// URL: /products?price=19.99&price=bad&price=39.99
//	prices := query.Slice(r, "price", 0.0, func(s string) (float64, error) {
//	    return strconv.ParseFloat(s, 64)
//	})
//	// Returns []float64{19.99, 0.0, 39.99}
//
// For convenience, typed slice helpers are provided:
//
//	ids := query.Ints(r, "id", 0)           // []int with default 0
//	vals := query.Int64s(r, "val", -1)      // []int64 with default -1
//	prices := query.Float64s(r, "price", 0.0) // []float64 with default 0.0
//	flags := query.Bools(r, "enabled", false) // []bool with default false
//
// # Numeric Types
//
// The package supports various numeric types with automatic parsing:
//
//	// URL: /api/stats?count=42&ratio=3.14&id=123
//	count := query.Int(r, "count", 0)          // int: 42
//	ratio := query.Float64(r, "ratio", 0.0)    // float64: 3.14
//	id    := query.Int64(r, "id", 0)           // int64: 123
//
// # Boolean Parsing
//
// Booleans are parsed flexibly, accepting common true/false representations:
//
//	// URL: /settings?enabled=true&admin=1&debug=yes
//	enabled := query.Bool(r, "enabled", false)  // true ("true")
//	admin   := query.Bool(r, "admin", false)    // true ("1")
//	debug   := query.Bool(r, "debug", false)    // true ("yes")
//	feature := query.Bool(r, "feature", false)  // false (missing)
//
// Recognized as true: "true", "1", "yes", "on", "y"
// Recognized as false: "false", "0", "no", "off", "n"
// Case-insensitive parsing
//
// # Error Handling
//
// Invalid values safely fall back to defaults without panicking:
//
//	// URL: /search?page=invalid&limit=-5&active=maybe
//	page   := query.Int(r, "page", 1)       // 1 (invalid parse)
//	limit  := query.Int(r, "limit", 25)     // -5 (parsed successfully)
//	active := query.Bool(r, "active", false) // false (unrecognized bool)
//
// Note: Negative numbers and zero are valid parse results. Use validation
// logic after extraction if you need to enforce constraints.
//
// # Common Patterns
//
// Pagination:
//
//	page  := query.Int(r, "page", 1)
//	limit := query.Int(r, "limit", 25)
//	if page < 1 {
//	    page = 1
//	}
//	if limit < 1 || limit > 100 {
//	    limit = 25
//	}
//
// Filtering:
//
//	status   := query.String(r, "status", "all")
//	category := query.String(r, "category", "")
//	tags     := query.Strings(r, "tag")
//	activeOnly := query.Bool(r, "active_only", false)
//
// Search:
//
//	q        := query.String(r, "q", "")
//	page     := query.Int(r, "page", 1)
//	perPage  := query.Int(r, "per_page", 20)
//	sortBy   := query.String(r, "sort_by", "relevance")
//	sortDir  := query.String(r, "sort_dir", "desc")
//
// # Design Principles
//
// 1. Fail-safe: Never panic on invalid input
// 2. Predictable: Missing or invalid values always return the default
// 3. Concise: One function call per parameter
// 4. Type-safe: Return values match the requested type
// 5. Zero dependencies: Uses only the standard library
//
// # Performance Note
//
// Each function call parses r.URL.Query() independently. For high-frequency
// extraction of many parameters, consider parsing once and storing the result:
//
//	values := r.URL.Query()
//	// Then use values.Get() directly if performance is critical
//
// However, for typical web applications, the convenience outweighs the minimal
// performance overhead.
package query
