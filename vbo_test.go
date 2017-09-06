package glplus

import (
	"testing"

	"github.com/go-gl/glfw3/v3.2/glfw"
)

func checkGlError(t *testing.T) {
	if Gl.GetError() != Gl.NO_ERROR {
		t.Fail()
	}
}

func subtestNewVBO(t *testing.T) {
	vbo := NewVBO(VBOOptions{})
	checkGlError(t)
	defer vbo.DeleteVBO()
}

func subtestNewVBOQuad(t *testing.T) {
	vbo := NewVBOQuad(0, 0, 1, 1)
	checkGlError(t)
	defer vbo.DeleteVBO()
}

func subtestNewVBOCube(t *testing.T) {
	vbo := NewVBOCube(0, 0, 0, 1, 1, 1)
	checkGlError(t)
	defer vbo.DeleteVBO()
}

func subtestNewVBOCubeNormal(t *testing.T) {
	vbo := NewVBOCubeNormal(0, 0, 0, 1, 1, 1)
	checkGlError(t)
	defer vbo.DeleteVBO()
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

	t.Run("NewVBO", subtestNewVBO)
	t.Run("NewVBOQuad", subtestNewVBOQuad)
	t.Run("NewVBOCube", subtestNewVBOCube)
	t.Run("NewVBOCubeNormal", subtestNewVBOCubeNormal)
}
