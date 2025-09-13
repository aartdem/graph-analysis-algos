package mathutil

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"zeros", 0, 0, 0},
		{"positives", 2, 3, 5},
		{"negatives", -2, -3, -5},
		{"mix", -10, 7, -3},
		{"big", 1 << 20, 1, 1<<20 + 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Sum(tc.a, tc.b); got != tc.expected {
				t.Fatalf("Sum(%d,%d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
