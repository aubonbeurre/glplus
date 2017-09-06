package glplus

import "sync/atomic"

const gPositionAttr int = 1 // prog.GetAttribLocation("position")
const gUVsAttr int = 2      // prog.GetAttribLocation("uvs")
const gNormalsAttr int = 3  // prog.GetAttribLocation("normal")

// Ogl2ShaderCompat ...
var Ogl2ShaderCompat = false

// ReferenceCountable ...
type ReferenceCountable interface {
	Decr() bool
	Incr()
}

// ReleasingReferenceCount ...
type ReleasingReferenceCount struct {
	count *int32
}

// Incr ...
func (rrc ReleasingReferenceCount) Incr() {
	atomic.AddInt32(rrc.count, 1)
}

// Decr ...
func (rrc ReleasingReferenceCount) Decr() bool {
	if atomic.AddInt32(rrc.count, -1) == 0 {
		return true
	}
	return false
}

// NewReferenceCount ...
func NewReferenceCount() ReleasingReferenceCount {
	var cnt = new(int32)
	*cnt = 1
	return ReleasingReferenceCount{
		count: cnt,
	}
}
