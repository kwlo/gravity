package physics

import (
	"math"
	"testing"
)

func TestNewCoord(t *testing.T) {
	got := NewCoord3D(1, 2, 3)

	if got.Len() != 3 {
		t.Fatal("Len() should always be 3")
	}
	if got.AtVec(0) != 1 {
		t.Fatal("AtVec(0) should retrieve the first value (x axis)")
	}
	if got.AtVec(2) != 3 {
		t.Fatal("AtVec(2) should retrieve the first value (z axis)")
	}
	if got.AtX() != 1 {
		t.Fatal("AtX() should retrieve the first value (x axis)")
	}
	if got.AtY() != 2 {
		t.Fatal("AtY() should retrieve the first value (y axis)")
	}
	if got.AtZ() != 3 {
		t.Fatal("AtZ() should retrieve the first value (z axis)")
	}
	got.X(4)
	got.Y(5)
	got.Z(6)
	if got.AtX() != 4 {
		t.Fatal("AtX() should retrieve the updated value (x axis)")
	}
	if got.AtY() != 5 {
		t.Fatal("AtY() should retrieve the updated value (y axis)")
	}
	if got.AtZ() != 6 {
		t.Fatal("AtZ() should retrieve the updated value (z axis)")
	}
}

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

func TestDot(t *testing.T) {
	got := Dot(
		NewCoord3D(1, 2, 3),
		NewCoord3D(4, 5, 6),
	)

	want := 32.0
	if got != want {
		t.Fatalf("FAILED: Got %v instead of %v", got, want)
	}
}

func TestAdd(t *testing.T) {
	got := Add(
		NewCoord3D(1, 2, 3),
		NewCoord3D(4, 5, 6),
	)

	want := NewCoord3D(5, 7, 9)
	if got.AtX() != want.AtX() ||
		got.AtY() != want.AtY() ||
		got.AtZ() != want.AtZ() {
		t.Fatalf("FAILED: Got %v instead of %v", got, want)
	}
}

func TestSub(t *testing.T) {
	got := Sub(
		NewCoord3D(1, 2, 3),
		NewCoord3D(4, 5, 6),
	)

	want := NewCoord3D(-3, -3, -3)
	if got.AtX() != want.AtX() ||
		got.AtY() != want.AtY() ||
		got.AtZ() != want.AtZ() {
		t.Fatalf("FAILED: Got %v instead of %v", got, want)
	}
}

func TestMult(t *testing.T) {
	got := Mult(
		NewCoord3D(1, 2, 3),
		NewCoord3D(4, 5, 6),
	)

	want := NewCoord3D(4, 10, 18)
	if got.AtX() != want.AtX() ||
		got.AtY() != want.AtY() ||
		got.AtZ() != want.AtZ() {
		t.Fatalf("FAILED: Got %v instead of %v", got, want)
	}
}
