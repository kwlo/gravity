package physics

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

type boundingSphere struct{}

func (box *boundingSphere) TestPoint(pt Coord3D) bool {
	return false
}

func NewBoundingAxisAlignedBox(coords []Coord3D) BoundingVolume {
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

	return &boundingAxisAlignedBox{min, max}
}

func NewBoundingSphere(coords []Coord3D) BoundingVolume {
	return &boundingSphere{}
}
