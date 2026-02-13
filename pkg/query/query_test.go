package query

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		key          string
		defaultValue string
		expected     string
	}{
		{"present value", "/?name=Alice", "name", "Bob", "Alice"},
		{"missing key", "/?other=value", "name", "Bob", "Bob"},
		{"empty value", "/?name=", "name", "Bob", "Bob"},
		{"empty default", "/?name=Alice", "name", "", "Alice"},
		{"special chars", "/?msg=hello%20world", "msg", "", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := String(r, tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		key          string
		defaultValue int
		expected     int
	}{
		{"valid int", "/?page=42", "page", 1, 42},
		{"missing key", "/?other=value", "page", 1, 1},
		{"empty value", "/?page=", "page", 1, 1},
		{"invalid int", "/?page=abc", "page", 1, 1},
		{"negative int", "/?offset=-5", "offset", 0, -5},
		{"zero", "/?count=0", "count", 10, 0},
		{"large number", "/?id=999999", "id", 0, 999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Int(r, tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf("Int() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		key          string
		defaultValue int64
		expected     int64
	}{
		{"valid int64", "/?id=9223372036854775807", "id", 0, 9223372036854775807},
		{"missing key", "/?other=value", "id", 100, 100},
		{"invalid int64", "/?id=not-a-number", "id", 100, 100},
		{"negative", "/?offset=-9999999999", "offset", 0, -9999999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Int64(r, tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf("Int64() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		key          string
		defaultValue float64
		expected     float64
	}{
		{"valid float", "/?price=19.99", "price", 0.0, 19.99},
		{"integer as float", "/?price=20", "price", 0.0, 20.0},
		{"missing key", "/?other=value", "price", 9.99, 9.99},
		{"invalid float", "/?price=abc", "price", 9.99, 9.99},
		{"scientific notation", "/?val=1.23e-4", "val", 0.0, 1.23e-4},
		{"negative", "/?temp=-15.5", "temp", 0.0, -15.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Float64(r, tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf("Float64() = %f, want %f", got, tt.expected)
			}
		})
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		key          string
		defaultValue bool
		expected     bool
	}{
		// True values
		{"true", "/?active=true", "active", false, true},
		{"TRUE", "/?active=TRUE", "active", false, true},
		{"1", "/?active=1", "active", false, true},
		{"yes", "/?active=yes", "active", false, true},
		{"YES", "/?active=YES", "active", false, true},
		{"on", "/?active=on", "active", false, true},
		{"ON", "/?active=ON", "active", false, true},
		{"y", "/?active=y", "active", false, true},
		{"Y", "/?active=Y", "active", false, true},

		// False values
		{"false", "/?active=false", "active", true, false},
		{"FALSE", "/?active=FALSE", "active", true, false},
		{"0", "/?active=0", "active", true, false},
		{"no", "/?active=no", "active", true, false},
		{"NO", "/?active=NO", "active", true, false},
		{"off", "/?active=off", "active", true, false},
		{"OFF", "/?active=OFF", "active", true, false},
		{"n", "/?active=n", "active", true, false},
		{"N", "/?active=N", "active", true, false},

		// Edge cases
		{"missing key", "/?other=value", "active", false, false},
		{"empty value", "/?active=", "active", true, true},
		{"invalid value", "/?active=maybe", "active", false, false},
		{"whitespace", "/?active=" + url.QueryEscape(" true "), "active", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Bool(r, tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf("Bool() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStrings(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		key      string
		expected []string
	}{
		{
			"multiple values",
			"/?tag=go&tag=rust&tag=python",
			"tag",
			[]string{"go", "rust", "python"},
		},
		{
			"single value",
			"/?tag=go",
			"tag",
			[]string{"go"},
		},
		{
			"missing key",
			"/?other=value",
			"tag",
			[]string{},
		},
		{
			"empty values",
			"/?tag=&tag=",
			"tag",
			[]string{"", ""},
		},
		{
			"mixed empty and values",
			"/?tag=go&tag=&tag=rust",
			"tag",
			[]string{"go", "", "rust"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Strings(r, tt.key)

			if len(got) != len(tt.expected) {
				t.Fatalf("Strings() length = %d, want %d", len(got), len(tt.expected))
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Strings()[%d] = %q, want %q", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestInts(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		key          string
		defaultValue int
		expected     []int
	}{
		{
			"all valid",
			"/?id=1&id=2&id=3",
			"id",
			0,
			[]int{1, 2, 3},
		},
		{
			"with invalid",
			"/?id=1&id=invalid&id=3",
			"id",
			0,
			[]int{1, 0, 3},
		},
		{
			"custom default",
			"/?id=1&id=bad&id=3",
			"id",
			-1,
			[]int{1, -1, 3},
		},
		{
			"negative numbers",
			"/?id=-5&id=10&id=-3",
			"id",
			0,
			[]int{-5, 10, -3},
		},
		{
			"missing key",
			"/?other=value",
			"id",
			0,
			[]int{},
		},
		{
			"all invalid",
			"/?id=bad&id=worse",
			"id",
			99,
			[]int{99, 99},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Ints(r, tt.key, tt.defaultValue)

			if len(got) != len(tt.expected) {
				t.Fatalf("Ints() length = %d, want %d", len(got), len(tt.expected))
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Ints()[%d] = %d, want %d", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestInt64s(t *testing.T) {
	r := httptest.NewRequest("GET", "/?id=1&id=9223372036854775807&id=invalid&id=3", nil)
	got := Int64s(r, "id", -1)

	expected := []int64{1, 9223372036854775807, -1, 3}
	if len(got) != len(expected) {
		t.Fatalf("Int64s() length = %d, want %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("Int64s()[%d] = %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestFloat64s(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		key          string
		defaultValue float64
		expected     []float64
	}{
		{
			"all valid",
			"/?price=19.99&price=29.99&price=39.99",
			"price",
			0.0,
			[]float64{19.99, 29.99, 39.99},
		},
		{
			"with invalid",
			"/?price=19.99&price=bad&price=39.99",
			"price",
			0.0,
			[]float64{19.99, 0.0, 39.99},
		},
		{
			"custom default",
			"/?price=10.5&price=invalid",
			"price",
			-1.0,
			[]float64{10.5, -1.0},
		},
		{
			"scientific notation",
			"/?val=1.23e-4&val=5.67e2",
			"val",
			0.0,
			[]float64{1.23e-4, 5.67e2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Float64s(r, tt.key, tt.defaultValue)

			if len(got) != len(tt.expected) {
				t.Fatalf("Float64s() length = %d, want %d", len(got), len(tt.expected))
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Float64s()[%d] = %f, want %f", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestBools(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		key          string
		defaultValue bool
		expected     []bool
	}{
		{
			"all true variations",
			"/?flag=true&flag=1&flag=yes&flag=on",
			"flag",
			false,
			[]bool{true, true, true, true},
		},
		{
			"all false variations",
			"/?flag=false&flag=0&flag=no&flag=off",
			"flag",
			true,
			[]bool{false, false, false, false},
		},
		{
			"mixed valid",
			"/?flag=true&flag=false&flag=yes&flag=no",
			"flag",
			false,
			[]bool{true, false, true, false},
		},
		{
			"with invalid",
			"/?flag=true&flag=maybe&flag=false",
			"flag",
			false,
			[]bool{true, false, false},
		},
		{
			"invalid with true default",
			"/?flag=invalid",
			"flag",
			true,
			[]bool{true},
		},
		{
			"case insensitive",
			"/?flag=TRUE&flag=FALSE&flag=YES",
			"flag",
			false,
			[]bool{true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Bools(r, tt.key, tt.defaultValue)

			if len(got) != len(tt.expected) {
				t.Fatalf("Bools() length = %d, want %d", len(got), len(tt.expected))
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Bools()[%d] = %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestSliceGeneric(t *testing.T) {
	// Test with custom parser
	r := httptest.NewRequest("GET", "/?val=abc&val=def&val=ghi", nil)

	// Custom parser that returns string length
	lengthParser := func(s string) (int, error) {
		return len(s), nil
	}

	got := Slice(r, "val", 0, lengthParser)
	expected := []int{3, 3, 3}

	if len(got) != len(expected) {
		t.Fatalf("Slice() length = %d, want %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("Slice()[%d] = %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		key      string
		expected bool
	}{
		{"present with value", "/?name=Alice", "name", true},
		{"present empty", "/?name=", "name", true},
		{"present no equals", "/?active", "active", true},
		{"missing", "/?other=value", "name", false},
		{"empty query", "/", "name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Has(r, tt.key)
			if got != tt.expected {
				t.Errorf("Has() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		key      string
		expected int
	}{
		{"not present", "/", "id", 0},
		{"single value", "/?id=1", "id", 1},
		{"multiple values", "/?id=1&id=2&id=3", "id", 3},
		{"empty value", "/?id=", "id", 1},
		{"mixed", "/?id=1&id=&id=3", "id", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := Count(r, tt.key)
			if got != tt.expected {
				t.Errorf("Count() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestIsMultiple(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		key      string
		expected bool
	}{
		{"not present", "/", "id", false},
		{"single value", "/?id=1", "id", false},
		{"two values", "/?id=1&id=2", "id", true},
		{"three values", "/?id=1&id=2&id=3", "id", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", tt.url, nil)
			got := IsMultiple(r, tt.key)
			if got != tt.expected {
				t.Errorf("IsMultiple() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		name         string
		slice        []int
		defaultValue int
		expected     int
	}{
		{"non-empty slice", []int{1, 2, 3}, 0, 1},
		{"single element", []int{42}, 0, 42},
		{"empty slice", []int{}, 99, 99},
		{"nil slice", nil, 99, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := First(tt.slice, tt.defaultValue)
			if got != tt.expected {
				t.Errorf("First() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestAll(t *testing.T) {
	r := httptest.NewRequest("GET", "/?name=Alice&age=30&tag=go&tag=rust", nil)
	got := All(r)

	// Check it's a map with the right keys
	if len(got) != 3 {
		t.Errorf("All() returned %d keys, want 3", len(got))
	}

	// Check single values
	if nameVals, ok := got["name"]; !ok || len(nameVals) != 1 || nameVals[0] != "Alice" {
		t.Errorf("All()['name'] = %v, want [Alice]", nameVals)
	}

	// Check multi-values
	if tagVals, ok := got["tag"]; !ok || len(tagVals) != 2 {
		t.Errorf("All()['tag'] = %v, want 2 values", tagVals)
	}
}

// Benchmark tests
func BenchmarkInt(b *testing.B) {
	r := httptest.NewRequest("GET", "/?page=42", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Int(r, "page", 1)
	}
}

func BenchmarkString(b *testing.B) {
	r := httptest.NewRequest("GET", "/?name=Alice", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = String(r, "name", "")
	}
}

func BenchmarkBool(b *testing.B) {
	r := httptest.NewRequest("GET", "/?active=true", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Bool(r, "active", false)
	}
}
