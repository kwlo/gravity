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

func TestBarycentricCoords(t *testing.T) {
	gotu, gotv, gotw, gotx := BarycentricCoords(
		NewCoord3D(10, 0, 10),
		NewCoord3D(-10, 0, 10),
		NewCoord3D(0, 0, -10),
		NewCoord3D(0, 10, 0),
		NewCoord3D(2, 2, 2),
	)

	if gotu < 0 && gotv < 0 && gotw < 0 && gotx < 0 {
		// Point P should be inside the ABCD shape, therefore
		// u, v, w, x should be all positive
		values := []float64{gotu, gotv, gotw, gotx}
		t.Fatalf("FAILED: %#v", values)
	}
}
