package glplus

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png" // just because
	"os"
)

// GPTexture ...
type GPTexture struct {
	texture *Texture
	Size    image.Point
}

// GenTexture ...
func GenTexture(size image.Point) (texture *GPTexture) {
	texture = &GPTexture{texture: Gl.CreateTexture(), Size: size}
	return texture
}

// Handle ...
func (t *GPTexture) Handle() *Texture {
	return t.texture
}

// DeleteTexture ...
func (t *GPTexture) DeleteTexture() {
	if t.texture != nil {
		Gl.DeleteTexture(t.texture)
	}
}

// BindTexture ...
func (t *GPTexture) BindTexture(unit int) {
	Gl.ActiveTexture(Gl.TEXTURE0 + unit)
	Gl.BindTexture(Gl.TEXTURE_2D, t.texture)
}

// UnbindTexture ...
func (t *GPTexture) UnbindTexture(unit int) {
	Gl.ActiveTexture(Gl.TEXTURE0 + unit)
	Gl.BindTexture(Gl.TEXTURE_2D, nil)
}

// NewRGBATexture ...
func NewRGBATexture(rgba *image.RGBA, linear, repeat bool) (texture *GPTexture, err error) {
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}

	texture = GenTexture(rgba.Rect.Size())

	texture.BindTexture(0)
	if linear {
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_MIN_FILTER, Gl.LINEAR)
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_MAG_FILTER, Gl.LINEAR)
	} else {
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_MIN_FILTER, Gl.NEAREST)
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_MAG_FILTER, Gl.NEAREST)
	}
	if repeat {
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_WRAP_S, Gl.REPEAT)
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_WRAP_T, Gl.REPEAT)
	} else {
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_WRAP_S, Gl.CLAMP_TO_EDGE)
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_WRAP_T, Gl.CLAMP_TO_EDGE)
	}
	Gl.TexImage2D(
		Gl.TEXTURE_2D,
		0,
		Gl.RGBA,
		rgba.Rect.Size().X,
		rgba.Rect.Size().Y,
		Gl.RGBA,
		Gl.UNSIGNED_BYTE,
		Gl.Ptr(rgba.Pix))
	texture.UnbindTexture(0)

	return texture, nil
}

// LoadTexture ...
func LoadTexture(file string, linear, repeat bool) (texture *GPTexture, img image.Image, err error) {
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
