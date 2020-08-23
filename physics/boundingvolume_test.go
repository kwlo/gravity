package physics

import (
	"testing"
)

func TestNewBoundingAxisAlignedBox(t *testing.T) {
	coords := []Coord3D{
		NewCoord3D(-1, -2, -3),
		NewCoord3D(2, 3, 4),
	}

	got := NewBoundingAxisAlignedBox(coords)

	test := NewCoord3D(0, 0, 0)
	if !got.TestPoint(test) {
		t.Fatalf("%v, %v", coords[0].AtX(), got)
	}
}
