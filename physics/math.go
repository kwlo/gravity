package physics

import "gonum.org/v1/gonum/mat"

// Coord3D is a struct for x, y, z 3D coordinate
type Coord3D mat.Vector

// Matrix is a struct holding the NxM matrix values
type Matrix mat.Matrix

// NewCoord3D return new Coord3D
func NewCoord3D(
	x float64,
	y float64,
	z float64,
) mat.Vector {
	return mat.NewVecDense(3, []float64{x, y, z})
}

// NewMatrix returns new matrix given the data
func NewMatrix(rows int, cols int, data []float64) Matrix {
	return mat.NewDense(rows, cols, data)
}

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
	// detPBCD := Det(MatrixFrom4Coord3D(p, b, c, d))
	return 0, 0, 0, 0
}
