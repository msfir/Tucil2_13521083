package point

import "math"

type Point3D struct {
	x float32
	y float32
	z float32
}

var NumOfCalls uint32 = 0

func (p Point3D) ToSlice() []float32 {
	return []float32{p.x, p.y, p.z}
}

func Distance(a Point3D, b Point3D) float32 {
	NumOfCalls++;
	deltaX := float64(a.x - b.x)
	deltaY := float64(a.y - b.y)
	deltaZ := float64(a.z - b.z)
	return float32(math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ))
}

func CreatePoint3D(x float32, y float32, z float32) Point3D {
	return Point3D{x, y, z}
}

func (p Point3D) GetX() float32 {
	return p.x
}

func (p *Point3D) SetX(x float32) {
	p.x = x
}

func (p Point3D) GetY() float32 {
	return p.y
}

func (p *Point3D) SetY(y float32) {
	p.y = y
}

func (p Point3D) GetZ() float32 {
	return p.z
}

func (p *Point3D) SetZ(z float32) {
	p.z = z
}
