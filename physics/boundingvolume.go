package physics

import (
	"math"
)

// BoundingVolume is the basic bounding volume interface type
type BoundingVolume interface {
	TestPoint(Coord3D) bool
}

type boundingAxisAlignedBox struct {
	min Coord3D
	max Coord3D
}

func (box *boundingAxisAlignedBox) TestPoint(pt Coord3D) bool {
	if box.min.AtX() > pt.AtX() || box.max.AtX() < pt.AtX() {
		return false
	}
	if box.min.AtY() > pt.AtY() || box.max.AtY() < pt.AtY() {
		return false
	}
	if box.min.AtZ() > pt.AtZ() || box.max.AtZ() < pt.AtZ() {
		return false
	}
	return true
}

type boundingSphere struct {
	center Coord3D
	radius float64
}

func (box *boundingSphere) TestPoint(pt Coord3D) bool {
	if math.Pow(box.radius, 2) >= (math.Pow(box.center.AtX()-pt.AtX(), 2) +
		math.Pow(box.center.AtY()-pt.AtY(), 2) +
		math.Pow(box.center.AtZ()-pt.AtZ(), 2)) {
		return true
	}
	return false
}

func findMinMax(coords []Coord3D) (Coord3D, Coord3D) {
	min := NewCoord3D(coords[0].AtX(), coords[0].AtY(), coords[0].AtZ())
	max := NewCoord3D(coords[0].AtX(), coords[0].AtY(), coords[0].AtZ())

	for _, coord := range coords {
		if coord.AtX() < min.AtX() {
			min.X(coord.AtX())
		}
		if coord.AtY() < min.AtY() {
			min.Y(coord.AtY())
		}
		if coord.AtZ() < min.AtZ() {
			min.Z(coord.AtZ())
		}

		if coord.AtX() > max.AtX() {
			max.X(coord.AtX())
		}
		if coord.AtY() > max.AtY() {
			max.Y(coord.AtY())
		}
		if coord.AtZ() > max.AtZ() {
			max.Z(coord.AtZ())
		}
	}
	return min, max
}

// NewBoundingAxisAlignedBox returns a BoundingVolume interface with axis aligned bounding box
func NewBoundingAxisAlignedBox(coords []Coord3D) BoundingVolume {
	min, max := findMinMax(coords)
	return &boundingAxisAlignedBox{min, max}
}

// NewBoundingSphereWithRadius returns a BoundingVolume with a bounding sphere from center point and radius
func NewBoundingSphereWithRadius(center Coord3D, radius float64) BoundingVolume {
	return &boundingSphere{center, radius}
}

// NewBoundingSphere returns a BoundingVolume with a bounding sphere from input coordinates. Implemented by using simple Ritter's bounding sphere
func NewBoundingSphere(coords []Coord3D) BoundingVolume {
	// Create initial sphere from min/max of point
	min, max := findMinMax(coords)

	center := NewCoord3D(
		(max.AtX()-min.AtX())*0.5,
		(max.AtY()-min.AtY())*0.5,
		(max.AtZ()-min.AtZ())*0.5,
	)
	radius := Dot(
		Sub(max, center),
		Sub(max, center),
	)
	radius = math.Sqrt(radius)

	// Rebound the sphere if there are any points outside the current sphere
	for _, coord := range coords {
		d := Sub(coord, center)
		dist := Dot(d, d)

		if dist > radius*radius {
			dist = math.Sqrt(dist)
			newRadius := (radius + dist) * 0.5
			k := (newRadius - radius) / dist
			radius = newRadius
			center = Add(
				center,
				NewCoord3D(
					d.AtX()*k,
					d.AtY()*k,
					d.AtZ()*k,
				),
			)
		}
	}

	return &boundingSphere{center, radius}
}
