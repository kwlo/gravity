package physics

import (
	"gonum.org/v1/gonum/mat"
)

// Coord3D is a struct for x, y, z 3D coordinate
type Coord3D interface {
	AtVec(int) float64
	Len() int
	AtX() float64
	AtY() float64
	AtZ() float64
	X(float64)
	Y(float64)
	Z(float64)
	Vector() mat.Vector
}

type coord3DImpl struct {
	vec *mat.VecDense
}

// AtVec returns the value at i index
func (c *coord3DImpl) Vector() mat.Vector {
	return c.vec
}

// AtVec returns the value at i index
func (c *coord3DImpl) AtVec(i int) float64 {
	return c.vec.AtVec(i)
}

// Len returns the length of the struct (should be 3)
func (c *coord3DImpl) Len() int {
	return c.vec.Len()
}

// AtX returns the X coord value
func (c *coord3DImpl) AtX() float64 {
	return c.vec.AtVec(0)
}

// AtY returns the Y coord value
func (c *coord3DImpl) AtY() float64 {
	return c.vec.AtVec(1)
}

// AtZ returns the Z coord value
func (c *coord3DImpl) AtZ() float64 {
	return c.vec.AtVec(2)
}

// X updates the input value of the x axis
func (c *coord3DImpl) X(val float64) {
	c.vec.SetVec(0, val)
}

// Y updates the input value of the y axis
func (c *coord3DImpl) Y(val float64) {
	c.vec.SetVec(1, val)
}

// Z updates the input value of the z axis
func (c *coord3DImpl) Z(val float64) {
	c.vec.SetVec(2, val)
}

// Matrix is a struct holding the NxM matrix values
type Matrix mat.Matrix

// NewCoord3D return new Coord3D
func NewCoord3D(
	x float64,
	y float64,
	z float64,
) Coord3D {
	return &coord3DImpl{mat.NewVecDense(3, []float64{x, y, z})}
}

// NewMatrix returns new matrix given the data
func NewMatrix(rows int, cols int, data []float64) Matrix {
	return mat.NewDense(rows, cols, data)
}

// Det returns the determinant of the given matrix
func Det(m Matrix) float64 {
	return mat.Det(m)
}

// MatrixFrom4Coord3D return 4x4 matrix from 4 vectors
// Returns as:
// {a1, b1, c1, d1}
// {a2, b2, c2, d2}
// {a3, b3, c3, d3}
// { 1,  1,  1,  1}
func MatrixFrom4Coord3D(
	a Coord3D,
	b Coord3D,
	c Coord3D,
	d Coord3D,
) Matrix {
	m := mat.NewDense(4, 4, []float64{
		a.AtVec(0), b.AtVec(0), c.AtVec(0), d.AtVec(0),
		a.AtVec(1), b.AtVec(1), c.AtVec(1), d.AtVec(1),
		a.AtVec(2), b.AtVec(2), c.AtVec(2), d.AtVec(2),
		1, 1, 1, 1,
	})

	return m
}

// BarycentricCoords returns the barycentric coords given
// tetrahedron ABCD and point P
func BarycentricCoords(
	a Coord3D,
	b Coord3D,
	c Coord3D,
	d Coord3D,
	p Coord3D,
) (float64, float64, float64, float64) {
	detPBCD := Det(MatrixFrom4Coord3D(p, b, c, d))
	detAPCD := Det(MatrixFrom4Coord3D(a, p, c, d))
	detABPD := Det(MatrixFrom4Coord3D(a, b, p, d))
	detABCP := Det(MatrixFrom4Coord3D(a, b, c, p))
	detABCD := Det(MatrixFrom4Coord3D(a, b, c, d))

	return detPBCD / detABCD,
		detAPCD / detABCD,
		detABPD / detABCD,
		detABCP / detABCD
}

// Dot returns the dot product of Coord3D a and b
func Dot(a Coord3D, b Coord3D) float64 {
	return mat.Dot(a.Vector(), b.Vector())
}

// Add returns the coord a + b
func Add(a Coord3D, b Coord3D) Coord3D {
	return NewCoord3D(
		a.AtX()+b.AtX(),
		a.AtY()+b.AtY(),
		a.AtZ()+b.AtZ(),
	)
}

// Sub returns the coord a - b
func Sub(a Coord3D, b Coord3D) Coord3D {
	return NewCoord3D(
		a.AtX()-b.AtX(),
		a.AtY()-b.AtY(),
		a.AtZ()-b.AtZ(),
	)
}

// Mult returns the coord a * b
func Mult(a Coord3D, b Coord3D) Coord3D {
	return NewCoord3D(
		a.AtX()*b.AtX(),
		a.AtY()*b.AtY(),
		a.AtZ()*b.AtZ(),
	)
}
