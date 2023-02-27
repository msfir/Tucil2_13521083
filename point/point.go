package point

import "math"

type Point3D struct {
	x, y, z float64
}

var NumOfCalls uint32 = 0

func (p Point3D) ToSlice() []float64 {
	return []float64{p.x, p.y, p.z}
}

func Distance(a, b Point3D) float64 {
	NumOfCalls++
	deltaX := float64(a.x - b.x)
	deltaY := float64(a.y - b.y)
	deltaZ := float64(a.z - b.z)
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ)
}

func CreatePoint3D(x, y, z float64) Point3D {
	return Point3D{x, y, z}
}

func (p Point3D) GetX() float64 {
	return p.x
}

func (p *Point3D) SetX(x float64) {
	p.x = x
}

func (p Point3D) GetY() float64 {
	return p.y
}

func (p *Point3D) SetY(y float64) {
	p.y = y
}

func (p Point3D) GetZ() float64 {
	return p.z
}

func (p *Point3D) SetZ(z float64) {
	p.z = z
}
