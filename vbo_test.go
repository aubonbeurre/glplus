package glplus

import (
	"testing"

	"github.com/go-gl/glfw3/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	sVertShaderCoordMarker = `#version 330
  ATTRIBUTE vec3 position;
  ATTRIBUTE vec3 normal;
  VARYINGOUT vec3 out_normal;
  uniform mat4 projection;
  uniform mat4 camera;
  uniform mat4 model;

  void main()
  {
      gl_Position = projection * camera * model * vec4(position, 1);
      out_normal = normalize(model * vec4(normal, 0)).xyz;
  }`

	sFragShaderCoordMarker = `#version 330
  uniform vec4 color1;
	uniform vec3 light;
  VARYINGIN vec3 out_normal;
  COLOROUT

  void main(void)
  {
  	float cosTheta = clamp(dot(light, normalize(out_normal)), 0.3, 1);
  	FRAGCOLOR = color1 * cosTheta;
  }`
)

func checkGlError(t *testing.T) {
	if Gl.GetError() != Gl.NO_ERROR {
		t.Fail()
	}
}

func subtestRenderVBO(t *testing.T) {
	var attribsNormal = []string{
		"position",
		"normal",
	}
	var err error
	var progCoord *Program
	if progCoord, err = LoadShaderProgram(sVertShaderCoordMarker, sFragShaderCoordMarker, attribsNormal); err != nil {
		t.Fatal(err)
	}
	defer progCoord.DeleteProgram()

	vbo := NewVBOCubeNormal(progCoord, 0, 0, 0, 1, 1, 1)
	defer vbo.DeleteVBO()
	checkGlError(t)

	progCoord.UseProgram()

	progCoord.ProgramUniformMatrix4fv("projection", mgl32.Ident4())
	progCoord.ProgramUniformMatrix4fv("camera", mgl32.Ident4())
	progCoord.ProgramUniformMatrix4fv("model", mgl32.Ident4())
	progCoord.ProgramUniform3fv("light", mgl32.Vec3{})
	progCoord.ProgramUniform4fv("color1", [4]float32{})

	vbo.Bind(progCoord)
	if err = progCoord.ValidateProgram(); err != nil {
		t.Fatal(err)
	}
	vbo.Draw()
	vbo.Unbind(progCoord)

	progCoord.UnuseProgram()
}

func TestVBO(t *testing.T) {
	var err error
	if err = glfw.Init(); err != nil {
		t.Fatal(err)
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Visible, glfw.False)

	glfw.WindowHint(glfw.Samples, 4)

	// do the actual window creation
	var window *glfw.Window
	window, err = glfw.CreateWindow(1024, 768, "test", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	window.MakeContextCurrent()

	Gl = NewContext()

	t.Run("RenderVBO", subtestRenderVBO)
}
