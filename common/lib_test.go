package common

import (
	"reflect"
	"testing"
)

// TestSumInts testar funktionen som summerar en slice av heltal.
func TestSumInts(t *testing.T) {
	// Ett testfall representeras av en struct som innehåller
	// in-data, förväntat resultat och ett namn.
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Positive numbers", []int{1, 2, 3}, 6},
		{"Negative numbers", []int{-1, -2, -3}, -6},
		{"Mixed numbers", []int{-1, 2, 5, -8}, -2},
		{"Empty slice", []int{}, 0},
		{"Single number", []int{42}, 42},
	}

	for _, tt := range tests {
		// tt.name gör att vi får ut ett tydligt namn för varje testfall i output.
		t.Run(tt.name, func(t *testing.T) {
			got := SumInts(tt.input)
			if got != tt.expected {
				t.Errorf("SumInts(%v) = %v; vill ha %v", tt.input, got, tt.expected)
			}
		})
	}
}

// TestMinMaxInt testar både MinInt och MaxInt.
func TestMinMaxInt(t *testing.T) {
	t.Run("MinInt", func(t *testing.T) {
		got := MinInt(5, 10)
		if got != 5 {
			t.Errorf("MinInt(5, 10) = %v; vill ha %v", got, 5)
		}
		got = MinInt(10, 5)
		if got != 5 {
			t.Errorf("MinInt(10, 5) = %v; vill ha %v", got, 5)
		}
		got = MinInt(5, 5)
		if got != 5 {
			t.Errorf("MinInt(5, 5) = %v; vill ha %v", got, 5)
		}
	})

	t.Run("MaxInt", func(t *testing.T) {
		got := MaxInt(5, 10)
		if got != 10 {
			t.Errorf("MaxInt(5, 10) = %v; vill ha %v", got, 10)
		}
		got = MaxInt(10, 5)
		if got != 10 {
			t.Errorf("MaxInt(10, 5) = %v; vill ha %v", got, 10)
		}
		got = MaxInt(5, 5)
		if got != 5 {
			t.Errorf("MaxInt(5, 5) = %v; vill ha %v", got, 5)
		}
	})
}

// TestMinMaxInts testar funktionerna som hittar minsta och största värdet i en slice.
func TestMinMaxInts(t *testing.T) {
	// MinInts
	t.Run("MinInts", func(t *testing.T) {
		input := []int{5, 2, 8, 1, 9}
		expectedVal, expectedKey := 1, 3
		gotVal, gotKey := MinInts(input)

		if gotVal != expectedVal || gotKey != expectedKey {
			t.Errorf("MinInts(%v) = (%v, %v); vill ha (%v, %v)", input, gotVal, gotKey, expectedVal, expectedKey)
		}

		input = []int{1, 2, 8, 1, 9} // Testar dubbletter av minsta värdet
		expectedVal, expectedKey = 1, 0
		gotVal, gotKey = MinInts(input)
		if gotVal != expectedVal || gotKey != expectedKey {
			t.Errorf("MinInts(%v) = (%v, %v); vill ha (%v, %v)", input, gotVal, gotKey, expectedVal, expectedKey)
		}
	})

	// MaxInts
	t.Run("MaxInts", func(t *testing.T) {
		input := []int{5, 2, 8, 1, 9}
		expectedVal, expectedKey := 9, 4
		gotVal, gotKey := MaxInts(input)

		if gotVal != expectedVal || gotKey != expectedKey {
			t.Errorf("MaxInts(%v) = (%v, %v); vill ha (%v, %v)", input, gotVal, gotKey, expectedVal, expectedKey)
		}

		input = []int{5, 2, 9, 1, 9} // Testar dubbletter av största värdet
		expectedVal, expectedKey = 9, 2
		gotVal, gotKey = MaxInts(input)
		if gotVal != expectedVal || gotKey != expectedKey {
			t.Errorf("MaxInts(%v) = (%v, %v); vill ha (%v, %v)", input, gotVal, gotKey, expectedVal, expectedKey)
		}
	})
}


// TestContainsString testar funktionen som kontrollerar om en sträng finns i en slice.
func TestContainsString(t *testing.T) {
	slice := []string{"apple", "banana", "orange"}

	t.Run("Contains existing string", func(t *testing.T) {
		if !ContainsString(slice, "banana") {
			t.Errorf("ContainsString should return true for 'banana'")
		}
	})

	t.Run("Does not contain string", func(t *testing.T) {
		if ContainsString(slice, "grape") {
			t.Errorf("ContainsString should return false for 'grape'")
		}
	})

	t.Run("Empty slice", func(t *testing.T) {
		if ContainsString([]string{}, "anything") {
			t.Errorf("ContainsString should return false for empty slice")
		}
	})
}

// TestStrToInt testar konverteringen av en sträng till ett heltal.
func TestStrToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{" 456 ", 456}, // Testar med mellanslag
	}

	for _, tt := range tests {
		got := StrToInt(tt.input)
		if got != tt.expected {
			t.Errorf("StrToInt(%q) = %v; vill ha %v", tt.input, got, tt.expected)
		}
	}
}

// TestStringsToInt testar konverteringen av en slice av strängar till en slice av heltal.
func TestStringsToInt(t *testing.T) {
	input := []string{"1", "2", "3", " ", "", "4"}
	expected := []int{1, 2, 3, 4}

	got := StringsToInt(input)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("StringsToInt(%v) = %v; vill ha %v", input, got, expected)
	}
}

// TestIntAbs testar funktionen som returnerar absolutvärdet av ett heltal.
func TestIntAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
	}

	for _, tt := range tests {
		got := IntAbs(tt.input)
		if got != tt.expected {
			t.Errorf("IntAbs(%v) = %v; vill ha %v", tt.input, got, tt.expected)
		}
	}
}
// Test IntToStr som returnerar strängen av ett heltal
func TestIntToStr(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{5, "5"},
		{-5, "-5"},
		{0, "0"},
	}

	for _, tt := range tests {
		got := IntToStr(tt.input)
		if got != tt.expected {
			t.Errorf("IntAbs(%v) = %v; vill ha %v", tt.input, got, tt.expected)
		}
	}
}
