package glplus

import (
	"os"
	"testing"

	"github.com/go-gl/glfw3/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func subtestRenderOBJ(t *testing.T) {
	var err error
	var fd *os.File
	if fd, err = os.Open("windarrow.obj"); err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	var objs []*Obj
	if objs, err = LoadObj(fd, nil); err != nil {
		t.Fatal(err)
	}
	objrender := NewObjVBO(objs[0])
	defer objrender.Delete()

	mat := objrender.NormalizedMat()

	objrender.Draw([4]float32{1, 1, 1, 1}, mgl32.Ident4(), mgl32.Ident4(), mat, mgl32.Vec3{1, 0, 0}, 0)

	checkGlError(t)
}

func TestRenderOBJ(t *testing.T) {
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
	t.Run("RenderOBJ", subtestRenderOBJ)
}
