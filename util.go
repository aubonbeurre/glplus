package glplus

import "sync/atomic"

// Gl may become engo.Gl (Gl = glplus.NewContext())
var Gl *Context

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
