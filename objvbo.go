package glplus

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	sVertShaderObj = `#version 330
  ATTRIBUTE vec3 position;
  ATTRIBUTE vec2 uvs;
  ATTRIBUTE vec3 normal;
  VARYINGOUT vec2 out_uvs;
  VARYINGOUT vec3 out_normal;
  uniform mat4 projection;
  uniform mat4 camera;
  uniform mat4 model;

  void main()
  {
      gl_Position = projection * camera * model * vec4(position, 1.0);
      out_uvs = uvs;
      out_normal = normalize(model * vec4(normal, 0.0)).xyz;
  }`

	sFragShaderObj = `#version 330
  uniform vec4 color1;
  uniform vec3 light;
  VARYINGIN vec2 out_uvs;
  VARYINGIN vec3 out_normal;
  COLOROUT

  void main(void)
  {
  	float cosTheta = clamp(dot(light, normalize(out_normal)), 0.3, 1.0);
  	FRAGCOLOR = color1 * cosTheta;
  }`

	sFragShaderObjTex = `#version 330
  uniform vec4 color1;
  uniform vec3 light;
	uniform sampler2D tex1;
	uniform mat3 matuv;
  VARYINGIN vec2 out_uvs;
  VARYINGIN vec3 out_normal;
  COLOROUT

  void main(void)
  {
		vec2 new_uvs = vec2(1.0-out_uvs.x, out_uvs.y);
		new_uvs = (matuv * vec3(new_uvs, 1)).xy;
		vec4 texcolor = TEXTURE2D(tex1, new_uvs);
  	float cosTheta = clamp(dot(light, normalize(out_normal)), 0.3, 1.0);
		FRAGCOLOR = mix(color1, texcolor, texcolor.w);
  	FRAGCOLOR = FRAGCOLOR * cosTheta;
  }`
)

// ObjRender ...
type ObjRender struct {
	Obj *Obj

	progCoord *GPProgram
	vbo       *VBO
	tex       *GPTexture
}

// ObjsRender ...
type ObjsRender struct {
	Objs []*ObjRender
}

// Bounds ...
func (m *ObjsRender) Bounds() (b Bounds) {
	var build BoundBuilder
	build.reset()

	for _, obj := range m.Objs {
		build.include64(obj.Obj.Bounds.X.Lo, obj.Obj.Bounds.Y.Lo, obj.Obj.Bounds.Z.Lo)
		build.include64(obj.Obj.Bounds.X.Hi, obj.Obj.Bounds.Y.Hi, obj.Obj.Bounds.Z.Hi)
	}
	return build.build()
}

// NormalizedMat ...
func (m *ObjsRender) NormalizedMat() (mres mgl32.Mat4) {
	center := m.Bounds().Center()
	length := math.Max(m.Bounds().X.Length(), m.Bounds().Y.Length())
	length = math.Max(length, m.Bounds().Z.Length())
	scale := float32(1 / length)

	mres = mgl32.HomogRotate3DX(-math.Pi / 2)
	mres = mres.Mul4(mgl32.Scale3D(scale, scale, scale))
	mres = mres.Mul4(mgl32.Translate3D(-center[0], -center[1], -center[2]))
	return mres
}

// Delete ...
func (m *ObjsRender) Delete() {
	for _, obj := range m.Objs {
		obj.Delete()
	}
}

// Draw ...
func (m *ObjsRender) Draw(color1 [4]float32, camera, projection, model mgl32.Mat4, light mgl32.Vec3, time float64) {
	for _, obj := range m.Objs {
		obj.Draw(color1, camera, projection, model, light, time)
	}
}

// NewObjsVBO ...
func NewObjsVBO(objs []*Obj) (m *ObjsRender) {
	m = &ObjsRender{}
	for _, obj := range objs {
		m.Objs = append(m.Objs, NewObjVBO(obj))
	}
	return m
}

// NewObjVBO ...
func NewObjVBO(obj *Obj) (m *ObjRender) {
	var err error

	m = &ObjRender{
		Obj: obj,
	}

	var attribsNormal = []string{
		"position",
		"uvs",
		"normal",
	}
	if obj.TexImg != nil {
		if m.progCoord, err = LoadShaderProgram(sVertShaderObj, sFragShaderObjTex, attribsNormal); err != nil {
			panic(err)
		}

		if m.tex, err = NewRGBATexture(obj.TexImg, true, false); err != nil {
			panic(err)
		}
	} else {
		if m.progCoord, err = LoadShaderProgram(sVertShaderObj, sFragShaderObj, attribsNormal); err != nil {
			panic(err)
		}
	}

	opt := DefaultVBOOptions()
	opt.Normals = 3
	m.vbo = NewVBO(m.progCoord, opt, obj.ObjVertices, nil)

	return m
}

// Delete ...
func (m *ObjRender) Delete() {
	m.progCoord.DeleteProgram()
	m.vbo.DeleteVBO()
	if m.tex != nil {
		m.tex.DeleteTexture()
	}
}

// NormalizedMat ...
func (m *ObjRender) NormalizedMat() (mres mgl32.Mat4) {
	center := m.Obj.Bounds.Center()
	length := math.Max(m.Obj.Bounds.X.Length(), m.Obj.Bounds.Y.Length())
	length = math.Max(length, m.Obj.Bounds.Z.Length())
	scale := float32(1 / length)

	mres = mgl32.HomogRotate3DX(-math.Pi / 2)
	mres = mres.Mul4(mgl32.Scale3D(scale, scale, scale))
	mres = mres.Mul4(mgl32.Translate3D(-center[0], -center[1], -center[2]))
	return mres
}

// Draw ...
func (m *ObjRender) Draw(color1 [4]float32, camera, projection, model mgl32.Mat4, light mgl32.Vec3, time float64) {
	m.progCoord.UseProgram()

	matuv := mgl32.Translate2D(0.5, 0.5)
	matuv = matuv.Mul3(mgl32.Rotate3DZ(-float32(time)))
	matuv = matuv.Mul3(mgl32.Translate2D(-0.5, -0.5))

	m.progCoord.ProgramUniformMatrix4fv("projection", projection)
	m.progCoord.ProgramUniformMatrix4fv("camera", camera)
	m.progCoord.ProgramUniformMatrix4fv("model", model)
	m.progCoord.ProgramUniform3fv("light", light)
	m.progCoord.ProgramUniformMatrix3fv("matuv", matuv)

	if m.tex != nil {
		m.tex.BindTexture(0)
		m.progCoord.ProgramUniform1i("tex1", 0)
	}

	m.vbo.Bind(m.progCoord)
	m.progCoord.ProgramUniform4fv("color1", color1)
	var err error
	if err = m.progCoord.ValidateProgram(); err != nil {
		panic(err)
	}
	m.vbo.Draw()
	m.vbo.Unbind(m.progCoord)

	if m.tex != nil {
		m.tex.UnbindTexture(0)
	}

	m.progCoord.UnuseProgram()
}
