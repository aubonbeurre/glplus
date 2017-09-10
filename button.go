package glplus

import (
	"image"

	"github.com/go-gl/mathgl/mgl32"
)

// TODO: https://github.com/memononen/nanovg

var (
	// https://thebookofshaders.com/07/
	testFontShader = `
#ifdef GL_ES
precision mediump float;
#endif

uniform vec2 u_resolution;
uniform vec2 u_mouse;
uniform float u_time;

void main(){
  vec2 st = gl_FragCoord.xy/u_resolution.xy;
  st.x *= u_resolution.x/u_resolution.y;
  vec3 color = vec3(0.0);
  float d = 0.0;

  // Remap the space to -1. to 1.
  st = st *2.-1.;

  // Make the distance field
   d = length( max(abs(st)-0.392,0.) );

  // Drawing with the distance field
	float intensity = smoothstep(0.444,0.480,d)* smoothstep(0.592,0.548,d);
	float intensity2 = 1.0-smoothstep(0.0,1.180,d);
	vec4 col1 = intensity * vec4(0.286,0.669,0.800,1.000);
	vec4 col2 = intensity2 * vec4(0.291,0.212,1.000,1.000);
	gl_FragColor = mix(col1, col2, col2.a);
}`

	// fragment shader
	fragShaderButton = `#version 330
  VARYINGIN vec2 out_uvs;
	uniform float animTime;
  COLOROUT

  void main(){
    vec2 st = out_uvs;

    // Remap the space to -1. to 1.
    st = st *2.-1.;

    // Make the distance field
    float d = length( max(abs(st)-0.392+animTime,0.) );

    // Drawing with the distance field
    float intensity = smoothstep(0.444,0.480,d)* smoothstep(0.592,0.548,d);
    float intensity2 = 1.0-smoothstep(0.0,1.180,d);
    vec4 col1 = intensity * vec4(0.286,0.669,0.800,1.000);
    vec4 col2 = intensity2 * vec4(0.291,0.212,1.000,1.000);
    FRAGCOLOR = mix(col1, col2, col2.a);
  }
  `
	// vertex shader
	vertShaderButton = `#version 330
  ATTRIBUTE vec4 position;
  ATTRIBUTE vec2 uvs;
  VARYINGOUT vec2 out_uvs;
  uniform mat3 ModelviewMatrix;
  void main()
  {
		gl_Position = vec4(ModelviewMatrix * vec3(position.xy, 1.0), 0.0).xywz;
  	out_uvs = uvs;
  }`
)

// ButtonProgram ...
type ButtonProgram struct {
	ReleasingReferenceCount

	vbo     *VBO
	program *GPProgram
}

// Delete ...
func (b *ButtonProgram) Delete() {
	if b.vbo != nil {
		b.vbo.DeleteVBO()
	}
	if b.program != nil {
		b.program.DeleteProgram()
	}
}

var sButtonProgramSingleton *ButtonProgram

// Button ...
type Button struct {
	Size image.Point
}

// Delete ...
func (b *Button) Delete() {
	if sButtonProgramSingleton.Decr() {
		sButtonProgramSingleton.Delete()
		sButtonProgramSingleton = nil
	}
}

// Draw ...
func (b *Button) Draw(color [4]float32, bg [4]float32, animTime float32, mat mgl32.Mat3) (err error) {
	program := sButtonProgramSingleton.program
	vbo := sButtonProgramSingleton.vbo

	program.UseProgram()
	vbo.Bind(program)

	var m = mat.Mul3(mgl32.Scale2D(float32(b.Size.X), float32(b.Size.Y)))

	program.ProgramUniformMatrix3fv("ModelviewMatrix", m)
	program.ProgramUniform4fv("color", color)
	program.ProgramUniform4fv("bg", bg)
	program.ProgramUniform1f("animTime", animTime/3)

	if err = program.ValidateProgram(); err != nil {
		return err
	}

	Gl.Disable(Gl.DEPTH_TEST)
	Gl.Enable(Gl.BLEND)

	vbo.Draw()

	Gl.Enable(Gl.DEPTH_TEST)
	Gl.Disable(Gl.BLEND)

	vbo.Unbind(program)
	program.UnuseProgram()

	return nil
}

// NewButton ...
func NewButton(size image.Point) (btn *Button, err error) {
	if sButtonProgramSingleton == nil {
		var attribs = []string{
			"position",
			"uvs",
		}
		var err error
		var program *GPProgram
		if program, err = LoadShaderProgram(vertShaderButton, fragShaderButton, attribs); err != nil {
			return nil, err
		}
		sButtonProgramSingleton = &ButtonProgram{
			vbo: NewVBOQuad(program, 0, 0, 1, 1),
			ReleasingReferenceCount: NewReferenceCount(),
			program:                 program,
		}

	} else {
		sButtonProgramSingleton.Incr()
	}

	var result = &Button{
		Size: size,
	}
	return result, nil
}
