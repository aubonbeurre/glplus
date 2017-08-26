package glplus

import (
	"unsafe"

	gl "github.com/go-gl/gl/v4.1-core/gl"
)

// VBO ...
type VBO struct {
	vao        uint32
	vboVerts   uint32
	vboIndices uint32
	numElem    int
	hasNormals bool
	isStrip    bool
}

// DeleteVBO ...
func (v *VBO) DeleteVBO() {
	if v.vboVerts != 0 {
		gl.DeleteBuffers(1, &v.vboVerts)
	}
	if v.vboIndices != 0 {
		gl.DeleteBuffers(1, &v.vboIndices)
	}
	if v.vao != 0 {
		gl.DeleteVertexArrays(1, &v.vao)
	}
}

// Bind ...
func (v *VBO) Bind() {
	gl.BindVertexArray(v.vao)
	gl.EnableVertexAttribArray(gPositionAttr)
	gl.EnableVertexAttribArray(gUVsAttr)
	if v.hasNormals {
		gl.EnableVertexAttribArray(gNormalsAttr)
	}
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.vboIndices)
}

// Unbind ...
func (v *VBO) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.DisableVertexAttribArray(gPositionAttr)
	gl.DisableVertexAttribArray(gUVsAttr)
	if v.hasNormals {
		gl.DisableVertexAttribArray(gNormalsAttr)
	}
	gl.BindVertexArray(0)
}

// Draw ...
func (v *VBO) Draw() {
	if v.isStrip {
		gl.DrawElements(gl.TRIANGLE_STRIP, int32(v.numElem), gl.UNSIGNED_INT, nil)
	} else {
		gl.DrawElements(gl.TRIANGLES, int32(v.numElem), gl.UNSIGNED_INT, nil)
	}
}

// DrawQuads ...
func (v *VBO) DrawQuads(nquads int) {
	gl.DrawElements(gl.TRIANGLES, int32(nquads*6), gl.UNSIGNED_INT, nil)
}

// Load ...
func (v *VBO) Load(verts *float32, vsize int, indices *uint32, isize int) {
	// calculate the memory size of floats used to calculate total memory size of float arrays
	var floatSize = int(unsafe.Sizeof(float32(1.0)))
	var intSize = int(unsafe.Sizeof(uint32(1)))

	gl.BindBuffer(gl.ARRAY_BUFFER, v.vboVerts)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.vboIndices)

	// load our data up and bind it to the 'position' shader attribute
	gl.BufferData(gl.ARRAY_BUFFER, floatSize*vsize, unsafe.Pointer(verts), gl.STATIC_DRAW)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, intSize*isize, unsafe.Pointer(indices), gl.STATIC_DRAW)

	if v.hasNormals {
		gl.VertexAttribPointer(gPositionAttr, 3, gl.FLOAT, false, 32, gl.PtrOffset(0))
		gl.VertexAttribPointer(gUVsAttr, 2, gl.FLOAT, false, 32, gl.PtrOffset(12))
		gl.VertexAttribPointer(gNormalsAttr, 3, gl.FLOAT, false, 32, gl.PtrOffset(20))
	} else {
		gl.VertexAttribPointer(gPositionAttr, 3, gl.FLOAT, false, 20, gl.PtrOffset(0))
		gl.VertexAttribPointer(gUVsAttr, 2, gl.FLOAT, false, 20, gl.PtrOffset(12))

	}

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	v.numElem = isize
}

// NewVBO ...
func NewVBO(isStrip bool) (vbo *VBO) {
	// create and bind the required VAO object
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// create a VBO to hold the vertex data
	var vboVerts uint32
	var vboIndices uint32
	gl.GenBuffers(1, &vboVerts)
	gl.GenBuffers(1, &vboIndices)

	vbo = &VBO{vao: vao,
		vboVerts:   vboVerts,
		vboIndices: vboIndices,
		numElem:    0,
		hasNormals: false,
		isStrip:    isStrip,
	}
	return vbo
}

// NewVBOQuad ...
func NewVBOQuad(x float32, y float32, w float32, h float32) (vbo *VBO) {
	vbo = NewVBO(false)

	verts := [...]float32{
		x, y, 0.0, 0, 0,
		x + w, y, 0.0, 1, 0,
		x + w, y + h, 0.0, 1, 1,
		x, y + h, 0.0, 0, 1,
	}

	indices := [...]uint32{
		0, 1, 2,
		2, 3, 0,
	}

	vbo.Load(&verts[0], len(verts), &indices[0], len(indices))

	return vbo
}

// NewVBOCube ...
func NewVBOCube(x float32, y float32, z float32, u float32, v float32, w float32) (vbo *VBO) {
	vbo = NewVBO(false)

	verts := [...]float32{
		// front
		-1.0, -1.0, 1.0, 0, 0,
		1.0, -1.0, 1.0, 0, 0,
		1.0, 1.0, 1.0, 0, 0,
		-1.0, 1.0, 1.0, 0, 0,
		// back
		-1.0, -1.0, -1.0, 0, 0,
		1.0, -1.0, -1.0, 0, 0,
		1.0, 1.0, -1.0, 0, 0,
		-1.0, 1.0, -1.0, 0, 0,
	}

	var i uint32
	for i = 0; i < 8; i++ {
		var ind = i * 5
		verts[ind+0] = verts[ind+0]*u + x
		verts[ind+1] = verts[ind+1]*v + y
		verts[ind+2] = verts[ind+2]*w + z
	}

	indices := [...]uint32{
		// front
		0, 1, 2,
		2, 3, 0,
		// top
		1, 5, 6,
		6, 2, 1,
		// back
		7, 6, 5,
		5, 4, 7,
		// bottom
		4, 0, 3,
		3, 7, 4,
		// left
		4, 5, 1,
		1, 0, 4,
		// right
		3, 2, 6,
		6, 7, 3,
	}

	vbo.Load(&verts[0], len(verts), &indices[0], len(indices))

	return vbo
}

// NewVBOCubeNormal ...
func NewVBOCubeNormal(x float32, y float32, z float32, u float32, v float32, w float32) (vbo *VBO) {
	vbo = NewVBO(true)
	vbo.hasNormals = true

	verts := [...]float32{
		// Vertex data for face 0
		-1.0, -1.0, 1.0, 0.0, 0.0, 0, 0, 1, // v0
		1.0, -1.0, 1.0, 0.33, 0.0, 0, 0, 1, // v1
		-1.0, 1.0, 1.0, 0.0, 0.5, 0, 0, 1, // v2
		1.0, 1.0, 1.0, 0.33, 0.5, 0, 0, 1, // v3

		// Vertex data for face 1
		1.0, -1.0, 1.0, 0.0, 0.5, 1, 0, 0, // v4
		1.0, -1.0, -1.0, 0.33, 0.5, 1, 0, 0, // v5
		1.0, 1.0, 1.0, 0.0, 1.0, 1, 0, 0, // v6
		1.0, 1.0, -1.0, 0.33, 1.0, 1, 0, 0, // v7

		// Vertex data for face 2
		1.0, -1.0, -1.0, 0.66, 0.5, 0, 0, -1, // v8
		-1.0, -1.0, -1.0, 1.0, 0.5, 0, 0, -1, // v9
		1.0, 1.0, -1.0, 0.66, 1.0, 0, 0, -1, // v10
		-1.0, 1.0, -1.0, 1.0, 1.0, 0, 0, -1, // v11

		// Vertex data for face 3
		-1.0, -1.0, -1.0, 0.66, 0.0, -1, 0, 0, // v12
		-1.0, -1.0, 1.0, 1.0, 0.0, -1, 0, 0, // v13
		-1.0, 1.0, -1.0, 0.66, 0.5, -1, 0, 0, // v14
		-1.0, 1.0, 1.0, 1.0, 0.5, -1, 0, 0, // v15

		// Vertex data for face 4
		-1.0, -1.0, -1.0, 0.33, 0.0, 0, -1, 0, // v16
		1.0, -1.0, -1.0, 0.66, 0.0, 0, -1, 0, // v17
		-1.0, -1.0, 1.0, 0.33, 0.5, 0, -1, 0, // v18
		1.0, -1.0, 1.0, 0.66, 0.5, 0, -1, 0, // v19

		// Vertex data for face 5
		-1.0, 1.0, 1.0, 0.33, 0.5, 0, 1, 0, // v20
		1.0, 1.0, 1.0, 0.66, 0.5, 0, 1, 0, // v21
		-1.0, 1.0, -1.0, 0.33, 1.0, 0, 1, 0, // v22
		1.0, 1.0, -1.0, 0.66, 1.0, 0, 1, 0, // v23
	}

	var i uint32
	for i = 0; i < 24; i++ {
		var ind = i * 8
		verts[ind+0] = verts[ind+0]*u + x
		verts[ind+1] = verts[ind+1]*v + y
		verts[ind+2] = verts[ind+2]*w + z
	}

	indices := [...]uint32{
		0, 1, 2, 3, 3, // Face 0 - triangle strip ( v0,  v1,  v2,  v3)
		4, 4, 5, 6, 7, 7, // Face 1 - triangle strip ( v4,  v5,  v6,  v7)
		8, 8, 9, 10, 11, 11, // Face 2 - triangle strip ( v8,  v9, v10, v11)
		12, 12, 13, 14, 15, 15, // Face 3 - triangle strip (v12, v13, v14, v15)
		16, 16, 17, 18, 19, 19, // Face 4 - triangle strip (v16, v17, v18, v19)
		20, 20, 21, 22, 23, // Face 5 - triangle strip (v20, v21, v22, v23)
	}

	vbo.Load(&verts[0], len(verts), &indices[0], len(indices))

	return vbo
}
