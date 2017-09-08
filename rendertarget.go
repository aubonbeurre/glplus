//+build !netgo,!android

package glplus

import (
	"fmt"
	"image"
	"image/color"
	"unsafe"
)

// RenderTarget ...
type RenderTarget struct {
	fbuffer  *FrameBuffer
	rbuffer  *RenderBuffer
	zbuffer  *RenderBuffer
	hasDepth bool
	Tex      *GPTexture
}

// Delete ...
func (r *RenderTarget) Delete() {
	if r.fbuffer != nil {
		Gl.DeleteFrameBuffer(r.fbuffer)
	}
	if r.rbuffer != nil {
		Gl.DeleteRenderBuffer(r.rbuffer)
	}
	if r.zbuffer != nil {
		Gl.DeleteRenderBuffer(r.zbuffer)
	}
	if r.Tex != nil {
		r.Tex.DeleteTexture()
	}
}

// EnsureSize ...
func (r *RenderTarget) EnsureSize(size image.Point) {
	if r.Tex == nil || r.Tex.Size.X != size.X || r.Tex.Size.Y != size.Y {
		if r.Tex != nil {
			r.Tex.DeleteTexture()
		}
		r.Tex = GenTexture(image.Point{size.X, size.Y})
		r.Tex.BindTexture(0)
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_MIN_FILTER, Gl.NEAREST)
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_MAG_FILTER, Gl.NEAREST)
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_WRAP_S, Gl.CLAMP_TO_EDGE)
		Gl.TexParameteri(Gl.TEXTURE_2D, Gl.TEXTURE_WRAP_T, Gl.CLAMP_TO_EDGE)
		Gl.TexImage2D(
			Gl.TEXTURE_2D,
			0,
			Gl.RGBA,
			size.X,
			size.Y,
			Gl.RGBA,
			Gl.UNSIGNED_BYTE,
			nil)
		r.Tex.UnbindTexture(0)
	}
}

func checkFramebufferStatus() string {
	status := Gl.CheckFramebufferStatus(Gl.FRAMEBUFFER)
	switch status {
	case Gl.FRAMEBUFFER_COMPLETE:
		return "OK"
	case Gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT:
		return "Framebuffer incomplete, incomplete attachment!"
	case Gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:
		return "Framebuffer incomplete, missing attachment!"
	case Gl.FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER:
		return "Framebuffer incomplete, missing draw buffer!"
	case Gl.FRAMEBUFFER_INCOMPLETE_READ_BUFFER:
		return "Framebuffer incomplete, missing read buffer!"
	case Gl.FRAMEBUFFER_UNSUPPORTED:
		return "Unsupported framebuffer format!"
	}
	panic(fmt.Errorf("Unknown error %d", status))
}

// Bind ...
func (r *RenderTarget) Bind(tex *GPTexture) {
	// Bind the frame-buffer object and attach to it a render-buffer object set up as a depth-buffer.
	Gl.BindFrameBuffer(Gl.FRAMEBUFFER, r.fbuffer)

	if r.hasDepth {
		Gl.BindRenderBuffer(Gl.RENDERBUFFER, r.zbuffer)
		Gl.RenderbufferStorage(Gl.RENDERBUFFER, Gl.DEPTH_COMPONENT, tex.Size.X, tex.Size.Y)
	}

	Gl.FramebufferTexture2D(Gl.FRAMEBUFFER, Gl.COLOR_ATTACHMENT0, Gl.TEXTURE_2D, tex.Handle(), 0)

	// Set the render target - primary surface
	Gl.DrawBuffer(Gl.COLOR_ATTACHMENT0)

	status := checkFramebufferStatus()
	if status != "OK" {
		Gl.BindFrameBuffer(Gl.FRAMEBUFFER, nil)
		panic(fmt.Errorf("Canot continue error %s", status))
	}
}

// Unbind ...
func (r *RenderTarget) Unbind(tex *GPTexture) {
	Gl.BindFrameBuffer(Gl.FRAMEBUFFER, nil)
	Gl.Flush()
}

// ReadBuffer ...
func (r *RenderTarget) ReadBuffer(w, h int) (newImage *image.RGBA) {
	Gl.Flush()
	Gl.ReadBuffer(Gl.COLOR_ATTACHMENT0)

	var pix = make([]float32, w*h*4)
	Gl.ReadPixels(0, 0, w, h, Gl.RGBA, Gl.FLOAT, unsafe.Pointer(&pix[0]))

	newImage = image.NewRGBA(image.Rect(0, 0, w, h))
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			ind := j*w*4 + i*4
			newImage.SetRGBA(i, j, color.RGBA{uint8(pix[ind+0] * 255), uint8(pix[ind+1] * 255), uint8(pix[ind+2] * 255), uint8(pix[ind+3] * 255)})
		}
	}
	return newImage
}

// NewRenderTarget ...
func NewRenderTarget(hasDepth bool) (r *RenderTarget) {
	const msaa = 4
	var rbuffer, zbuffer *RenderBuffer
	//now create the color render buffer
	rbuffer = Gl.CreateRenderBuffer()
	Gl.BindRenderBuffer(Gl.RENDERBUFFER, rbuffer)
	//Gl.RenderbufferStorageMultisample(Gl.RENDERBUFFER, msaa, GL_RGB8, width, height);

	if hasDepth {
		zbuffer = Gl.CreateRenderBuffer()
		Gl.BindRenderBuffer(Gl.RENDERBUFFER, zbuffer)
		//Gl.RenderbufferStorageMultisample(Gl.RENDERBUFFER, msaa, GL_DEPTH_COMPONENT, width, height);
	}

	var fbuffer *FrameBuffer
	//create the color buffer to render to
	fbuffer = Gl.CreateFrameBuffer()
	Gl.BindFrameBuffer(Gl.FRAMEBUFFER, fbuffer)

	Gl.FramebufferRenderbuffer(Gl.FRAMEBUFFER, Gl.COLOR_ATTACHMENT0, Gl.RENDERBUFFER, rbuffer)
	if hasDepth {
		Gl.FramebufferRenderbuffer(Gl.FRAMEBUFFER, Gl.DEPTH_ATTACHMENT, Gl.RENDERBUFFER, zbuffer)
	}

	Gl.BindFrameBuffer(Gl.FRAMEBUFFER, nil)

	r = &RenderTarget{
		fbuffer:  fbuffer,
		rbuffer:  rbuffer,
		zbuffer:  zbuffer,
		hasDepth: hasDepth,
	}
	return r
}
