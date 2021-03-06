package glplus

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	ifont "golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	sUseGray = false
)

var (
	// fragment shader
	fragShaderFont = `#version 330
  VARYINGIN vec2 out_uvs;
  uniform sampler2D tex1;
  uniform vec4 color;
  uniform vec4 bg;
	COLOROUT

  void main()
  {
    vec4 col0 = TEXTURE2D(tex1, out_uvs);
    FRAGCOLOR = col0.r * color;
    // Porter duff gl.ONE, gl.ONE_MINUS_SRC_ALPHA
    FRAGCOLOR = vec4(FRAGCOLOR.r + bg.r * (1.0-FRAGCOLOR.a), FRAGCOLOR.g + bg.g * (1.0-FRAGCOLOR.a), FRAGCOLOR.b + bg.b * (1.0-FRAGCOLOR.a), FRAGCOLOR.a + bg.a * (1-FRAGCOLOR.a));
  }`

	// vertex shader
	vertShaderFont = `#version 330
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

// Char ...
type Char struct {
	Index int
	X     int
	Y     int
}

// String ...
type String struct {
	Chars []Char
	Size  image.Point
	vbo   *VBO
	font  *Font
}

// DeleteString ...
func (s *String) DeleteString() {
	if s.vbo != nil {
		s.vbo.DeleteVBO()
	}
}

// Draw ...
func (s *String) Draw(f *Font, color [4]float32, bg [4]float32, mat mgl32.Mat3, scale float32, offsetX float32, offsetY float32) (err error) {
	if f.program == nil {
		var attribs = []string{
			"position",
			"uvs",
		}
		if f.program, err = LoadShaderProgram(vertShaderFont, fragShaderFont, attribs); err != nil {
			return (err)
		}
	}
	if s.vbo == nil {
		s.createVertexBuffer(f)
	}

	f.program.UseProgram()
	f.texture.BindTexture(0)
	s.vbo.Bind(f.program)

	var matrixfont = mat.Mul3(mgl32.Scale2D(scale, scale))
	matrixfont = matrixfont.Mul3(mgl32.Translate2D(offsetX, offsetY))
	f.program.ProgramUniformMatrix3fv("ModelviewMatrix", matrixfont)
	f.program.ProgramUniform1i("tex1", 0)
	f.program.ProgramUniform4fv("color", color)
	f.program.ProgramUniform4fv("bg", bg)

	if err = f.program.ValidateProgram(); err != nil {
		return err
	}

	s.vbo.Draw()

	f.texture.UnbindTexture(0)
	s.vbo.Unbind(f.program)
	f.program.UnuseProgram()

	return nil
}

func (s *String) createVertexBuffer(f *Font) {
	n := len(s.Chars)

	verts := make([]float32, n*20)
	indices := make([]uint32, n*6)

	/*
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
	*/
	var curX float32
	i := 0
	ii := 0
	var jj uint32
	var dv = float32(f.cellssize) / float32(f.texture.Size.Y)
	for j := 0; j < n; j++ {
		var c = s.Chars[j]
		var x = curX
		var y float32
		var w = float32(f.advances[c.Index])
		var h = float32(f.cellssize)
		var u = float32(c.X*f.cellssize) / float32(f.texture.Size.X)
		var v = float32(c.Y*f.cellssize) / float32(f.texture.Size.Y)
		var du = float32(w) / float32(f.texture.Size.X)

		verts[i+0] = x
		verts[i+1] = y
		verts[i+2] = 0
		verts[i+3] = u
		verts[i+4] = v
		i += 5

		verts[i+0] = x + w
		verts[i+1] = y
		verts[i+2] = 0
		verts[i+3] = u + du
		verts[i+4] = v
		i += 5

		verts[i+0] = x + w
		verts[i+1] = y + h
		verts[i+2] = 0
		verts[i+3] = u + du
		verts[i+4] = v + dv
		i += 5

		verts[i+0] = x
		verts[i+1] = y + h
		verts[i+2] = 0
		verts[i+3] = u
		verts[i+4] = v + dv
		i += 5

		indices[ii+0] = 0 + jj
		indices[ii+1] = 1 + jj
		indices[ii+2] = 2 + jj
		indices[ii+3] = 2 + jj
		indices[ii+4] = 3 + jj
		indices[ii+5] = 0 + jj
		ii += 6
		jj += 4

		curX += w
	}

	opt := DefaultVBOOptions()
	opt.Quads = n
	s.vbo = NewVBO(f.program, opt, verts[:], indices[:])
}

// Font ...
type Font struct {
	texture   *GPTexture
	cellssize int
	advances  []int

	program *GPProgram
	rows    int
}

// DeleteFont ...
func (f *Font) DeleteFont() {
	if f.texture != nil {
		f.texture.DeleteTexture()
	}
	if f.program != nil {
		f.program.DeleteProgram()
	}
}

// BindTexture ...
func (f *Font) BindTexture(unit int) {
	f.texture.BindTexture(unit)
}

// UnbindTexture ...
func (f *Font) UnbindTexture(unit int) {
	f.texture.UnbindTexture(unit)
}

// GetUVForChar ...
func (f *Font) GetUVForChar(c rune) mgl32.Vec4 {
	var ascii = int(c)
	var index = ascii - 32
	var xoff = index % f.rows
	var yoff = index / f.rows
	var dv = float32(f.cellssize) / float32(f.texture.Size.Y)
	var w = float32(f.advances[index])
	var u = float32(xoff*f.cellssize) / float32(f.texture.Size.X)
	var v = float32(yoff*f.cellssize) / float32(f.texture.Size.Y)
	var du = float32(w) / float32(f.texture.Size.X)
	return mgl32.Vec4{u, v, du, dv}
}

// NewString ...
func (f *Font) NewString(s string) *String {
	var result = &String{
		Chars: make([]Char, len(s)),
		Size:  image.Point{0, 0},
		font:  f,
	}
	var width int
	for i := 0; i < len(s); i++ {
		var ascii = int(s[i])
		var index = ascii - 32
		var xoff = index % f.rows
		var yoff = index / f.rows
		width += f.advances[index]

		//fmt.Printf("ascii: %d, x: %d, y: %d\n", ascii, xoff, yoff)
		result.Chars[i].Index = index
		result.Chars[i].X = xoff
		result.Chars[i].Y = yoff

	}
	result.Size = image.Point{width, f.cellssize}
	return result
}

// FreeSerif ...
func FreeSerif() (*os.File, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("%s", "No caller information")
	}

	return os.Open(path.Join(path.Dir(filename), "FreeSerif.ttf"))
}

// NewFont ...
func NewFont(reader io.Reader) (font *Font, err error) {

	// Read the font data.
	var fontBytes []byte
	if fontBytes, err = ioutil.ReadAll(reader); err != nil {
		return nil, err
	}
	var f *truetype.Font
	if f, err = freetype.ParseFont(fontBytes); err != nil {
		return nil, err
	}
	const fontSize = 48
	var face ifont.Face
	face = truetype.NewFace(f, &truetype.Options{Size: fontSize})
	height := face.Metrics().Height.Round()
	descent := face.Metrics().Descent.Round()
	//fmt.Printf("Height: %d\n", height)

	dst := image.NewRGBA(image.Rect(0, 0, height*16, height*16))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(dst, dst.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	d := &ifont.Drawer{
		Dst:  dst,
		Src:  image.White,
		Face: face,
	}

	var advances = make([]int, 256-32)
	var offx int
	var offy = height
	for i := 32; i < 255; i++ {
		d.Dot = fixed.P(offx, offy-descent)
		var strc = string(rune(i))
		d.DrawString(strc)
		if advance, ok := face.GlyphAdvance(rune(strc[0])); ok {
			advances[i-32] = advance.Round()
		} else {
			advances[i-32] = 0
		}

		offx += height
		if offx >= height*16 {
			offy += height
			offx = 0
		}
	}

	if false {
		w, _ := os.Create("font.png")
		defer w.Close()
		png.Encode(w, dst) //Encode writes the Image m to w in PNG format.
	}

	texture := GenTexture(dst.Rect.Size())
	texture.BindTexture(0)
	Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_MIN_FILTER, Gl.LINEAR)
	Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_MAG_FILTER, Gl.LINEAR)
	Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_WRAP_S, Gl.CLAMP_TO_EDGE)
	Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_WRAP_T, Gl.CLAMP_TO_EDGE)

	if sUseGray {
		gray := image.NewGray(dst.Bounds())
		if gray.Stride != gray.Rect.Size().X {
			return nil, fmt.Errorf("unsupported stride")
		}
		draw.Draw(gray, gray.Bounds(), dst, image.Point{0, 0}, draw.Src)

		Gl.TexImage2D(
			Gl.TEXTURE_2D,
			0,
			Gl.R8,
			gray.Rect.Size().X,
			gray.Rect.Size().Y,
			Gl.RED,
			Gl.UNSIGNED_BYTE,
			gray.Pix)
	} else {
		Gl.TexImage2D(
			Gl.TEXTURE_2D,
			0,
			Gl.RGBA,
			dst.Rect.Size().X,
			dst.Rect.Size().Y,
			Gl.RGBA,
			Gl.UNSIGNED_BYTE,
			dst.Pix)
	}
	texture.UnbindTexture(0)

	font = &Font{
		texture:   texture,
		rows:      16,
		cellssize: dst.Rect.Size().X / 16,
		advances:  advances,
	}

	return font, nil
}
