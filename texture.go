package glplus

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png" // just because
	"os"

	gl "github.com/go-gl/gl/v4.1-core/gl"
)

// Texture ...
type Texture struct {
	texture uint32
	Size    image.Point
}

// GenTexture ...
func GenTexture(size image.Point) (texture *Texture) {
	var t uint32
	gl.GenTextures(1, &t)
	texture = &Texture{t, size}
	return texture
}

// Handle ...
func (t *Texture) Handle() uint32 {
	return t.texture
}

// DeleteTexture ...
func (t *Texture) DeleteTexture() {
	if t.texture != 0 {
		gl.DeleteTextures(1, &t.texture)
	}
}

// BindTexture ...
func (t *Texture) BindTexture(unit uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + unit)
	gl.BindTexture(gl.TEXTURE_2D, t.texture)
}

// UnbindTexture ...
func (t *Texture) UnbindTexture(unit uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + unit)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

// NewRGBATexture ...
func NewRGBATexture(rgba *image.RGBA, linear, repeat bool) (texture *Texture, err error) {
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}

	texture = GenTexture(rgba.Rect.Size())

	texture.BindTexture(0)
	if linear {
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	} else {
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	}
	if repeat {
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	} else {
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	}
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	texture.UnbindTexture(0)

	return texture, nil
}

// LoadTexture ...
func LoadTexture(file string, linear, repeat bool) (texture *Texture, img image.Image, err error) {
	var imgFile *os.File
	if imgFile, err = os.Open(file); err != nil {
		return nil, img, err
	}
	defer imgFile.Close()

	if img, _, err = image.Decode(imgFile); err != nil {
		return nil, img, err
	}

	var rgba *image.RGBA

	rgba = image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	if texture, err = NewRGBATexture(rgba, linear, repeat); err != nil {
		return nil, img, err
	}

	return texture, img, nil
}
