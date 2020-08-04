package physics

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

func TestConvertDistance(t *testing.T) {
	tests := []struct {
		num    float64
		source DistanceEnum
		target DistanceEnum
		want   float64
	}{
		{10, Yards, Inches, 360},
		{72, Inches, Yards, 2},
		{3.5, Normalized, Yards, 35000},
		{-3.5, Normalized, Yards, -35000},
		{827, Yards, Normalized, 0.0827},
		{958, Yards, Miles, 0.544318182},
		{6.41, Miles, Yards, 11281.6},
		{5.21, Normalized, Inches, 1875600},
		{71, Inches, Normalized, 0.000197222},
	}

	for idx, tc := range tests {
		got := ConvertDistance(tc.num, tc.source, tc.target)
		if math.Abs(got-tc.want) > float64EqualityThreshold {
			t.Fatalf("FAILED: case #%d %v, %v, got: %v", idx, tc.num, tc.want, got)
		}
	}
}
