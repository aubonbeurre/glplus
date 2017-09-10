package glplus

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/golang/geo/r1"
)

// Bounds ...
type Bounds struct {
	X, Y, Z r1.Interval
}

// BoundBuilder ...
type BoundBuilder struct {
	minX float64
	minY float64
	minZ float64
	maxX float64
	maxY float64
	maxZ float64
}

func (r *BoundBuilder) reset() {
	r.minX = math.MaxFloat64
	r.minY = math.MaxFloat64
	r.minZ = math.MaxFloat64
	r.maxX = -math.MaxFloat64
	r.maxY = -math.MaxFloat64
	r.maxZ = -math.MaxFloat64
}

func (r *BoundBuilder) build() Bounds {
	return Bounds{
		X: r1.Interval{Lo: r.minX, Hi: r.maxX},
		Y: r1.Interval{Lo: r.minY, Hi: r.maxY},
		Z: r1.Interval{Lo: r.minZ, Hi: r.maxZ},
	}
}

func (r *BoundBuilder) include32(x, y, z float32) {
	r.minX = math.Min(r.minX, float64(x))
	r.minY = math.Min(r.minY, float64(y))
	r.minZ = math.Min(r.minZ, float64(z))
	r.maxX = math.Max(r.maxX, float64(x))
	r.maxY = math.Max(r.maxY, float64(y))
	r.maxZ = math.Max(r.maxZ, float64(z))
}

func (r *BoundBuilder) include64(x, y, z float64) {
	r.minX = math.Min(r.minX, float64(x))
	r.minY = math.Min(r.minY, float64(y))
	r.minZ = math.Min(r.minZ, float64(z))
	r.maxX = math.Max(r.maxX, float64(x))
	r.maxY = math.Max(r.maxY, float64(y))
	r.maxZ = math.Max(r.maxZ, float64(z))
}

// Center ...
func (b Bounds) Center() mgl32.Vec3 {
	return mgl32.Vec3{float32(b.X.Center()), float32(b.Y.Center()), float32(b.Z.Center())}
}

// Length ...
func (b Bounds) Length() float32 {
	length := math.Max(b.X.Length(), b.Y.Length())
	return float32(math.Max(length, b.Z.Length()))
}
