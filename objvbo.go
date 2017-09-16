package glplus

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	sVertShaderObj = `#version 330
  ATTRIBUTE vec3 position;
	ATTRIBUTE float uvs;
	ATTRIBUTE vec3 normal;
	VARYINGOUT float out_uvs;
	VARYINGOUT vec3 out_normal;
	VARYINGOUT vec4 out_color;
	uniform mat4 mProjViewModel;
	uniform mat4 mViewModel;
	uniform mat4 mView;
	uniform vec4 ambient;
	uniform float shininess;
	uniform vec4 specular;
	uniform vec4 diffuse;

  void main()
  {
      out_normal = normalize(mViewModel * vec4(normal, 0.0)).xyz;
			out_uvs = uvs;
			out_color = diffuse;
			gl_Position = mProjViewModel * vec4(position, 1.0);
  }`

	sFragShaderObj = `#version 330
  uniform vec3 light;
	uniform mat4 mView;
	VARYINGIN float out_uvs;
	VARYINGIN vec3 out_normal;
	VARYINGIN vec4 out_color;
  COLOROUT

  void main(void)
  {
		// "out_uvs / 1000" is there to trick the driver to use it
  	float cosTheta = clamp(dot(normalize(mView * vec4(light, 0)).xyz, normalize(out_normal)), 0.3 + out_uvs / 1000.0, 1.0);
  	FRAGCOLOR = out_color * cosTheta;
  }`

	sVertShaderObjTex = `#version 330
	ATTRIBUTE vec3 position;
	ATTRIBUTE vec2 uvs;
	ATTRIBUTE vec3 normal;
	VARYINGOUT vec2 out_uvs;
	VARYINGOUT vec3 out_normal;
	VARYINGOUT vec4 out_color;
	uniform mat4 mProjViewModel;
	uniform mat4 mViewModel;
	uniform mat4 mView;
	uniform vec4 ambient;
	uniform float shininess;
	uniform vec4 specular;
	uniform vec4 diffuse;

	void main()
	{
			out_uvs = uvs;
			out_normal = normalize(mViewModel * vec4(normal, 0.0)).xyz;
			out_color = diffuse;
			gl_Position = mProjViewModel * vec4(position, 1.0);
	}`

	sFragShaderObjTex = `#version 330
  uniform vec3 light;
	uniform mat4 mView;
	uniform sampler2D tex1;
	uniform mat3 matuv;
  VARYINGIN vec2 out_uvs;
  VARYINGIN vec3 out_normal;
	VARYINGIN vec4 out_color;
  COLOROUT

  void main(void)
  {
		vec2 new_uvs = vec2(1.0-out_uvs.x, out_uvs.y);
		new_uvs = (matuv * vec3(new_uvs, 1)).xy;
		vec4 texcolor = TEXTURE2D(tex1, new_uvs);
  	float cosTheta = clamp(dot(normalize(mView * vec4(light, 0)).xyz, normalize(out_normal)), 0.3, 1.0);
		FRAGCOLOR = mix(out_color, texcolor, texcolor.w) * cosTheta;
  }`

	sVertShaderObjColorTable = `#version 330
	ATTRIBUTE vec3 position;
	ATTRIBUTE float uvs;
	ATTRIBUTE vec3 normal;
	VARYINGOUT float out_uvs;
	VARYINGOUT vec3 out_normal;
	VARYINGOUT vec4 out_color;
	uniform mat4 mProjViewModel;
	uniform mat4 mViewModel;
	uniform mat4 mView;
	uniform vec4 ambient;
	uniform float shininess;
	uniform vec4 specular;
	uniform vec4 diffuse;

	void main()
	{
			out_uvs = uvs;
			out_normal = normalize(mViewModel * vec4(normal, 0.0)).xyz;
			out_color = diffuse;
			gl_Position = mProjViewModel * vec4(position, 1.0);
	}`

	sFragShaderObjColorTable = `#version 330
  uniform vec3 light;
	uniform mat4 mView;
	uniform sampler2D tex1;
  VARYINGIN float out_uvs;
  VARYINGIN vec3 out_normal;
	VARYINGIN vec4 out_color;
  COLOROUT

  void main(void)
  {
		vec4 texcolor = TEXTURE2D(tex1, vec2(out_uvs, 0));
  	float cosTheta = clamp(dot(normalize(mView * vec4(light, 0)).xyz, normalize(out_normal)), 0.3, 1.0);
  	FRAGCOLOR = texcolor * cosTheta;
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
	length := m.Bounds().Length()
	scale := 1 / length

	mres = mgl32.HomogRotate3DX(math.Pi / 2)
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
func (m *ObjsRender) Draw(material *Material, camera, projection, model mgl32.Mat4, light mgl32.Vec3, uvAngle float64, tex *GPTexture) {
	for _, obj := range m.Objs {
		obj.Draw(material, camera, projection, model, light, uvAngle, tex)
	}
}

// NewObjsVBO ...
func NewObjsVBO(objs []*Obj, hasColorTable bool) (m *ObjsRender) {
	m = &ObjsRender{}
	for _, obj := range objs {
		m.Objs = append(m.Objs, NewObjVBO(obj, hasColorTable))
	}
	return m
}

// NewObjVBO ...
func NewObjVBO(obj *Obj, hasColorTable bool) (m *ObjRender) {
	var err error

	m = &ObjRender{
		Obj: obj,
	}

	var attribs = []string{
		"position",
		"uvs",
		"normal",
	}
	if obj.TexImg != nil {
		if m.progCoord, err = LoadShaderProgram(sVertShaderObjTex, sFragShaderObjTex, attribs); err != nil {
			panic(err)
		}

		if m.tex, err = NewRGBATexture(obj.TexImg, true, false); err != nil {
			panic(err)
		}
	} else if hasColorTable {
		if m.progCoord, err = LoadShaderProgram(sVertShaderObjColorTable, sFragShaderObjColorTable, attribs); err != nil {
			panic(err)
		}
	} else {
		if m.progCoord, err = LoadShaderProgram(sVertShaderObj, sFragShaderObj, attribs); err != nil {
			panic(err)
		}
	}
	opt := DefaultVBOOptions()
	opt.Normals = 3
	if obj.TexImg == nil {
		opt.UV = 1
	}
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
	return m.Obj.NormalizedMat()
}

// Draw ...
func (m *ObjRender) Draw(material *Material, camera, projection, model mgl32.Mat4, light mgl32.Vec3, uvAngle float64, tex *GPTexture) {
	m.progCoord.UseProgram()

	matuv := mgl32.Translate2D(0.5, 0.5)
	matuv = matuv.Mul3(mgl32.Rotate3DZ(-float32(uvAngle)))
	matuv = matuv.Mul3(mgl32.Translate2D(-0.5, -0.5))

	mViewModel := camera.Mul4(model)
	mProjViewModel := projection.Mul4(mViewModel)
	m.progCoord.ProgramUniformMatrix4fv("mViewModel", mViewModel)
	m.progCoord.ProgramUniformMatrix4fv("mProjViewModel", mProjViewModel)
	m.progCoord.ProgramUniformMatrix4fv("mView", camera)
	m.progCoord.ProgramUniform3fv("light", light)
	m.progCoord.ProgramUniformMatrix3fv("matuv", matuv)

	if m.tex != nil {
		m.tex.BindTexture(0)
		m.progCoord.ProgramUniform1i("tex1", 0)
	} else if tex != nil {
		tex.BindTexture(0)
		m.progCoord.ProgramUniform1i("tex1", 0)
	}

	m.vbo.Bind(m.progCoord)
	m.progCoord.Material(material)

	var err error
	if err = m.progCoord.ValidateProgram(); err != nil {
		panic(err)
	}
	m.vbo.Draw()
	m.vbo.Unbind(m.progCoord)

	if m.tex != nil {
		m.tex.UnbindTexture(0)
	} else if tex != nil {
		m.tex.UnbindTexture(0)
	}

	m.progCoord.UnuseProgram()
}
