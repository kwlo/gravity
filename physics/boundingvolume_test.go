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
		t.Fatalf("%v shound be inside %v", test, got)
	}
}

func TestNewBoundingSphereWithRadius(t *testing.T) {
	tests := []struct {
		center    Coord3D
		radius    float64
		testPt    Coord3D
		assertion bool
	}{
		{NewCoord3D(0, 0, 0), 2, NewCoord3D(0, 0, 0), true},
		{NewCoord3D(0, 0, 0), 2, NewCoord3D(3, 0, 0), false},
		{NewCoord3D(1, 1, 0), 2, NewCoord3D(0, 0, 0), true},
		{NewCoord3D(3, 3, 0), 2, NewCoord3D(0, 0, 0), false},
	}

	for idx, test := range tests {
		bv := NewBoundingSphereWithRadius(test.center, test.radius)
		got := bv.TestPoint(test.testPt)
		if got != test.assertion {
			t.Fatalf("Assertion for #%v failed, %v", idx, bv)
		}
	}
}

func TestNewBoundingSphere(t *testing.T) {
	tests := []struct {
		coords    []Coord3D
		testPt    Coord3D
		assertion bool
	}{
		{
			[]Coord3D{
				NewCoord3D(-2, 0, 0),
				NewCoord3D(2, 0, 0),
			},
			NewCoord3D(0, 0, 0),
			true,
		},
		{
			[]Coord3D{
				NewCoord3D(-1, 0, 0),
				NewCoord3D(3, 0, 0),
			},
			NewCoord3D(0, 0, 0),
			true,
		},
		{
			[]Coord3D{
				NewCoord3D(-2, 0, 0),
				NewCoord3D(2, 0, 0),
			},
			NewCoord3D(0, 2, 0),
			true,
		},
		{
			[]Coord3D{
				NewCoord3D(-2, 0, 0),
				NewCoord3D(2, 0, 0),
			},
			NewCoord3D(2, 2, 0),
			false,
		},
		{
			[]Coord3D{
				NewCoord3D(-2, 0, 0),
				NewCoord3D(2, 0, 0),
			},
			NewCoord3D(0, 3, 0),
			false,
		},
		{
			[]Coord3D{
				NewCoord3D(-2, 0, 0),
				NewCoord3D(2, 0, 0),
				NewCoord3D(0, 3, 0),
			},
			NewCoord3D(3, 0, 0),
			true,
		},
	}

	for idx, test := range tests {
		bv := NewBoundingSphere(test.coords)
		got := bv.TestPoint(test.testPt)
		if got != test.assertion {
			t.Fatalf("Assertion for #%v failed, %v", idx, bv)
		}
	}
}
