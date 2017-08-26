package glplus

// Matrix2x3 ...
type Matrix2x3 struct {
	a float32
	b float32
	c float32
	d float32
	e float32
	f float32
}

// Concat ...
func (m *Matrix2x3) Concat(m1 Matrix2x3) Matrix2x3 {
	return Matrix2x3{m1.a*m.a + m1.b*m.c,
		m1.a*m.b + m1.b*m.d,
		m1.c*m.a + m1.d*m.c,
		m1.c*m.b + m1.d*m.d,
		m1.e*m.a + m1.f*m.c + m.e,
		m1.e*m.b + m1.f*m.d + m.f}
}

// Translate ...
func (m *Matrix2x3) Translate(x float32, y float32) Matrix2x3 {
	return Matrix2x3{m.a, m.b, m.c, m.d,
		m.e + (x*m.a + y*m.c),
		m.f + (x*m.b + y*m.d)}
}

// Scale ...
func (m *Matrix2x3) Scale(x float32, y float32) Matrix2x3 {
	return Matrix2x3{m.a * x,
		m.b * x,
		m.c * y,
		m.d * y,
		m.e,
		m.f}
}

// Array ...
func (m *Matrix2x3) Array() [16]float32 {
	return [...]float32{
		m.a, m.b, 0.0, 0.0,
		m.c, m.d, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		m.e, m.f, 0.0, 1.0,
	}
}

// IdentityMatrix2x3 ...
func IdentityMatrix2x3() Matrix2x3 {
	return Matrix2x3{1, 0, 0, 1, 0, 0}
}
