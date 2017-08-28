package glplus

import (
	"fmt"
	"image"
	"image/color"
	"unsafe"

	gl "github.com/go-gl/gl/v4.1-core/gl"
)

// RenderTarget ...
type RenderTarget struct {
	fbuffer uint32
	rbuffer uint32
}

// Delete ...
func (r *RenderTarget) Delete() {
	if r.fbuffer != 0 {
		gl.DeleteFramebuffers(1, &r.fbuffer)
	}
	if r.rbuffer != 0 {
		gl.DeleteRenderbuffers(1, &r.rbuffer)
	}
}

func checkFramebufferStatus() string {
	status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER)
	switch status {
	case gl.FRAMEBUFFER_COMPLETE:
		return "OK"
	case gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT:
		return "Framebuffer incomplete, incomplete attachment!"
	case gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:
		return "Framebuffer incomplete, missing attachment!"
	case gl.FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER:
		return "Framebuffer incomplete, missing draw buffer!"
	case gl.FRAMEBUFFER_INCOMPLETE_READ_BUFFER:
		return "Framebuffer incomplete, missing read buffer!"
	case gl.FRAMEBUFFER_UNSUPPORTED:
		return "Unsupported framebuffer format!"
	}
	panic(fmt.Errorf("Unknown error %d", status))
}

// Bind ...
func (r *RenderTarget) Bind(tex *Texture) {
	// Bind the frame-buffer object and attach to it a render-buffer object set up as a depth-buffer.
	gl.BindFramebuffer(gl.FRAMEBUFFER, r.fbuffer)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, tex.Handle(), 0)

	// Set the render target - primary surface
	gl.DrawBuffer(gl.COLOR_ATTACHMENT0)

	status := checkFramebufferStatus()
	if status != "OK" {
		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
		panic(fmt.Errorf("Canot continue error %s", status))
	}
}

// Unbind ...
func (r *RenderTarget) Unbind(tex *Texture) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Flush()
}

// ReadBuffer ...
func (r *RenderTarget) ReadBuffer(w, h int) (newImage *image.RGBA) {
	gl.Flush()
	gl.ReadBuffer(gl.COLOR_ATTACHMENT0)

	var pix = make([]float32, w*h*4)
	gl.ReadPixels(0, 0, int32(w), int32(h), gl.RGBA, gl.FLOAT, unsafe.Pointer(&pix[0]))

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
func NewRenderTarget() (r *RenderTarget) {
	var fbuffer uint32
	//create the color buffer to render to
	gl.GenFramebuffers(1, &fbuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbuffer)

	var rbuffer uint32
	//now create the color render buffer
	gl.GenRenderbuffers(1, &rbuffer)
	gl.BindRenderbuffer(gl.RENDERBUFFER, rbuffer)

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	r = &RenderTarget{
		fbuffer: fbuffer,
		rbuffer: rbuffer,
	}
	return r
}
