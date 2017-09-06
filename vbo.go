package glplus

import "math"

// VBOOptions ...
type VBOOptions struct {
	Vertex  int
	Normals int
	UV      int
	IsStrip bool
	Quads   int
}

func DefaultVBOOptions() VBOOptions {
	return VBOOptions{
		Vertex: 3,
		UV:     2,
	}
}

// VBO ...
type VBO struct {
	vao        *ENGOGLVertexArray
	vboVerts   *ENGOGLBuffer
	vboIndices *ENGOGLBuffer
	numElem    int
	isShort    bool

	options VBOOptions
}

// DeleteVBO ...
func (v *VBO) DeleteVBO() {
	if v.vboVerts != nil {
		Gl.DeleteBuffer(v.vboVerts)
	}
	if v.vboIndices != nil {
		Gl.DeleteBuffer(v.vboIndices)
	}
	if v.vao != nil {
		Gl.DeleteVertexArray(v.vao)
	}
}

// Bind ...
func (v *VBO) Bind() {
	Gl.BindVertexArray(v.vao)
	if v.options.Vertex != 0 {
		Gl.EnableVertexAttribArray(gPositionAttr)
	}
	if v.options.UV != 0 {
		Gl.EnableVertexAttribArray(gUVsAttr)
	}
	if v.options.Normals != 0 {
		Gl.EnableVertexAttribArray(gNormalsAttr)
	}
	if v.vboIndices != nil {
		Gl.BindBuffer(Gl.ELEMENT_ARRAY_BUFFER, v.vboIndices)
	} else {
		Gl.BindBuffer(Gl.ARRAY_BUFFER, v.vboVerts)
	}
}

// Unbind ...
func (v *VBO) Unbind() {
	if v.vboIndices != nil {
		Gl.BindBuffer(Gl.ELEMENT_ARRAY_BUFFER, nil)
	} else {
		Gl.BindBuffer(Gl.ARRAY_BUFFER, nil)
	}
	if v.options.Vertex != 0 {
		Gl.DisableVertexAttribArray(gPositionAttr)
	}
	if v.options.UV != 0 {
		Gl.DisableVertexAttribArray(gUVsAttr)
	}
	if v.options.Normals != 0 {
		Gl.DisableVertexAttribArray(gNormalsAttr)
	}
	Gl.BindVertexArray(nil)
}

func (v *VBO) elemType() int {
	if v.isShort {
		return Gl.UNSIGNED_SHORT
	}

	return Gl.UNSIGNED_INT
}

// Draw ...
func (v *VBO) Draw() {
	if v.vboIndices != nil {
		if v.options.Quads != 0 {
			Gl.DrawElements(Gl.TRIANGLES, v.options.Quads*6, v.elemType(), 0)
		} else if v.options.IsStrip {
			Gl.DrawElements(Gl.TRIANGLE_STRIP, v.numElem, v.elemType(), 0)
		} else {
			Gl.DrawElements(Gl.TRIANGLES, v.numElem, v.elemType(), 0)
		}
	} else {
		if v.options.Quads != 0 {
			Gl.DrawArrays(Gl.TRIANGLES, 0, v.options.Quads*6)
		} else if v.options.IsStrip {
			Gl.DrawArrays(Gl.TRIANGLE_STRIP, 0, v.numElem)
		} else {
			Gl.DrawArrays(Gl.TRIANGLES, 0, v.numElem)
		}
	}
}

// Load ...
func (v *VBO) load(verts []float32, indices []uint32) {
	Gl.BindVertexArray(v.vao)
	Gl.BindBuffer(Gl.ARRAY_BUFFER, v.vboVerts)
	if v.vboIndices != nil {
		Gl.BindBuffer(Gl.ELEMENT_ARRAY_BUFFER, v.vboIndices)
	}

	// load our data up and bind it to the 'position' shader attribute
	Gl.BufferData(Gl.ARRAY_BUFFER, verts, Gl.STATIC_DRAW)

	if v.vboIndices != nil {
		if len(indices) < math.MaxUint16 {
			v.isShort = true
			uindices := make([]uint16, len(indices))
			for i, ind := range indices {
				uindices[i] = uint16(ind)
			}
			Gl.BufferData(Gl.ELEMENT_ARRAY_BUFFER, uindices, Gl.STATIC_DRAW)
		} else {
			Gl.BufferData(Gl.ELEMENT_ARRAY_BUFFER, indices, Gl.STATIC_DRAW)
		}
	}

	var numElemsPerVertex = v.options.Vertex + v.options.UV + v.options.Normals
	var totalSize = numElemsPerVertex * 4
	var offset int
	if v.options.Vertex != 0 {
		Gl.VertexAttribPointer(gPositionAttr, v.options.Vertex, Gl.FLOAT, false, totalSize, offset)
		offset += v.options.Vertex * 4
	}
	if v.options.UV != 0 {
		Gl.VertexAttribPointer(gUVsAttr, v.options.UV, Gl.FLOAT, false, totalSize, offset)
		offset += v.options.UV * 4
	}
	if v.options.Normals != 0 {
		Gl.VertexAttribPointer(gNormalsAttr, v.options.Normals, Gl.FLOAT, false, totalSize, offset)
	}

	Gl.BindBuffer(Gl.ARRAY_BUFFER, nil)
	if v.vboIndices != nil {
		Gl.BindBuffer(Gl.ELEMENT_ARRAY_BUFFER, nil)
	}
	Gl.BindVertexArray(nil)

	if v.vboIndices != nil {
		v.numElem = len(indices)
	} else {
		v.numElem = len(verts) / numElemsPerVertex
	}
}

// NewVBO ...
func NewVBO(options VBOOptions, verts []float32, indices []uint32) (vbo *VBO) {
	// create and bind the required VAO object
	var vao *ENGOGLVertexArray
	vao = Gl.CreateVertexArray()
	Gl.BindVertexArray(vao)

	// create a VBO to hold the vertex data
	var vboVerts *ENGOGLBuffer
	var vboIndices *ENGOGLBuffer
	vboVerts = Gl.CreateBuffer()
	if indices != nil {
		vboIndices = Gl.CreateBuffer()
	}

	vbo = &VBO{vao: vao,
		vboVerts:   vboVerts,
		vboIndices: vboIndices,
		numElem:    0,
		options:    options,
	}
	Gl.BindVertexArray(nil)

	vbo.load(verts, indices)

	return vbo
}

// NewVBOQuad ...
func NewVBOQuad(x float32, y float32, w float32, h float32) (vbo *VBO) {

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

	vbo = NewVBO(DefaultVBOOptions(), verts[:], indices[:])
	return vbo
}

// NewVBOCube ...
func NewVBOCube(x float32, y float32, z float32, u float32, v float32, w float32) (vbo *VBO) {
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

	vbo = NewVBO(DefaultVBOOptions(), verts[:], indices[:])
	return vbo
}

// NewVBOCubeNormal ...
func NewVBOCubeNormal(x float32, y float32, z float32, u float32, v float32, w float32) (vbo *VBO) {
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

	opt := DefaultVBOOptions()
	opt.IsStrip = true
	opt.Normals = 3
	vbo = NewVBO(opt, verts[:], indices[:])

	return vbo
}
