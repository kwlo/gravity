package physics

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-6

func TestConvertDistance(t *testing.T) {
	if ConvertDistance(10, Yards, Inches) != 360 {
		t.Errorf("wrong result")
	}
	if ConvertDistance(72, Inches, Yards) != 2 {
		t.Errorf("wrong result")
	}
	if ConvertDistance(3.5, Normalized, Yards) != 35000 {
		t.Errorf("wrong result")
	}
	if math.Abs(ConvertDistance(827, Yards, Normalized)-0.0827) > float64EqualityThreshold {
		t.Errorf("wrong result")
	}
	if math.Abs(ConvertDistance(958, Yards, Miles)-0.544318) > float64EqualityThreshold {
		t.Errorf("wrong result")
	}
	if math.Abs(ConvertDistance(6.41, Miles, Yards)-11281.6) > float64EqualityThreshold {
		t.Errorf("wrong result")
	}
	if math.Abs(ConvertDistance(5.21, Normalized, Inches)-1875600) > float64EqualityThreshold {
		t.Errorf("wrong result")
	}
	if math.Abs(ConvertDistance(71, Inches, Normalized)-0.000197222) > float64EqualityThreshold {
		t.Errorf("wrong result")
	}
}
