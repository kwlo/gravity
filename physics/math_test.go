package physics

import (
	"math"
	"testing"
)

func TestMatrixFrom4Coord3D(t *testing.T) {
	m := MatrixFrom4Coord3D(
		NewCoord3D(1, 2, 3),
		NewCoord3D(4, 5, 6),
		NewCoord3D(7, 8, 9),
		NewCoord3D(10, 11, 12),
	)

	if m.At(0, 1) != 4 {
		t.Fatal("1st row, 2nd item should be 4")
	}

	if m.At(1, 2) != 8 {
		t.Fatal("2nd row, 3rd item should be 8")
	}

	if m.At(2, 3) != 12 {
		t.Fatal("3rd row, 4th item should be 12")
	}

	if m.At(3, 0) != 1 {
		t.Fatal("4th row, 1st item should be 1")
	}
}

func TestDet(t *testing.T) {
	tests := []struct {
		rows int
		cols int
		data []float64
		want float64
	}{
		{2, 2, []float64{2, 3, 0, 4}, 8},
		{3, 3, []float64{4, -2, 6, 2, 5, 0, -2, 1, -4}, -24},
	}

	for idx, tc := range tests {
		got := Det(NewMatrix(tc.rows, tc.cols, tc.data))
		if math.Abs(got-tc.want) > float64EqualityThreshold {
			t.Fatalf("FAILED: case #%d %v, got: %v", idx, tc.want, got)
		}
	}
}
