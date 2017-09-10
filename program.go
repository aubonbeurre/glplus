package glplus

import (
	"fmt"
	"runtime"
	"strings"
)

// GPProgram ...
type GPProgram struct {
	prog *Program

	uniforms map[string]*UniformLocation
	attribs  []string
	hash     string
}

// DeleteProgram ...
func (p *GPProgram) DeleteProgram() {
	if cache := sProgCache[p.hash]; cache != nil {
		if cache.Decr() {
			//fmt.Printf("Delete %s\n", p.hash)
			cache.Delete()
		}
	}
}

// GetProgramInfoLog ...
func (p *GPProgram) GetProgramInfoLog() string {
	return Gl.GetProgramInfoLog(p.prog)
}

// ValidateProgram ...
func (p *GPProgram) ValidateProgram() error {
	Gl.ValidateProgram(p.prog)

	if !Gl.GetProgramParameterb(p.prog, Gl.VALIDATE_STATUS) {
		return fmt.Errorf("Failed to validate the program!\n%s", p.GetProgramInfoLog())
	}
	return nil
}

// GetUniformLocation ...
func (p *GPProgram) GetUniformLocation(s string) *UniformLocation {
	var res *UniformLocation
	var ok bool
	if res, ok = p.uniforms[s]; ok {
		return res
	}

	res = Gl.GetUniformLocation(p.prog, s)
	p.uniforms[s] = res
	return res
}

// GetAttribLocation ...
func (p *GPProgram) GetAttribLocation(s string) int {
	return Gl.GetAttribLocation(p.prog, s)
}

// UseProgram ...
func (p *GPProgram) UseProgram() {
	Gl.UseProgram(p.prog)
}

// UnuseProgram ...
func (p *GPProgram) UnuseProgram() {
	Gl.UseProgram(nil)
}

// ProgramUniform1f ...
func (p *GPProgram) ProgramUniform1f(uniform string, value float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	Gl.Uniform1f(uniformloc, value)
}

// ProgramUniform2f ...
func (p *GPProgram) ProgramUniform2f(uniform string, v0 float32, v1 float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	Gl.Uniform2f(uniformloc, v0, v1)
}

// ProgramUniform4fv ...
func (p *GPProgram) ProgramUniform4fv(uniform string, value [4]float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	Gl.Uniform4f(uniformloc, value[0], value[1], value[2], value[3])
}

// ProgramUniform3fv ...
func (p *GPProgram) ProgramUniform3fv(uniform string, value [3]float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	Gl.Uniform3f(uniformloc, value[0], value[1], value[2])
}

// ProgramUniform1i ...
func (p *GPProgram) ProgramUniform1i(uniform string, value int) {
	var uniformloc = p.GetUniformLocation(uniform)
	Gl.Uniform1i(uniformloc, value)
}

// ProgramUniformMatrix4fv ...
func (p *GPProgram) ProgramUniformMatrix4fv(uniform string, matrix [16]float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	Gl.UniformMatrix4fv(uniformloc, false, matrix[:])
}

// ProgramUniformMatrix3fv ...
func (p *GPProgram) ProgramUniformMatrix3fv(uniform string, matrix [9]float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	Gl.UniformMatrix3fv(uniformloc, false, matrix[:])
}

// GetShaderInfoLog ...
func GetShaderInfoLog(shader *Shader) string {
	return Gl.GetShaderInfoLog(shader)
}

// ShaderSource ...
func ShaderSource(shader *Shader, src string) {
	Gl.ShaderSource(shader, src)
}

// LoadShaderProgram ... loads shader objects and then attaches them to a program
func LoadShaderProgram(vertShader string, fragShader string, attribs []string) (*GPProgram, error) {
	// query program cache
	if sProgCache == nil {
		sProgCache = make(map[string]*ProgramCache)
	}
	hash := getMD5Hash(vertShader, fragShader)
	if cache := sProgCache[hash]; cache != nil {
		cache.Incr()
		//fmt.Printf("Hit %s\n", cache.program.hash)
		return cache.program, nil
	}

	// create the program
	var prog = Gl.CreateProgram()
	var p = &GPProgram{
		prog:     prog,
		uniforms: make(map[string]*UniformLocation),
		attribs:  attribs,
		hash:     hash,
	}

	if runtime.GOARCH == "js" || runtime.GOOS == "android" {
		vertShader = strings.Replace(vertShader, "#version 330\n", "precision mediump float;\n", -1)
		vertShader = strings.Replace(vertShader, "ATTRIBUTE", "attribute", -1)
		vertShader = strings.Replace(vertShader, "VARYINGOUT", "varying", -1)
		vertShader = strings.Replace(vertShader, "TEXTURE2D", "texture2D", -1)
	} else {
		vertShader = strings.Replace(vertShader, "ATTRIBUTE", "in", -1)
		vertShader = strings.Replace(vertShader, "VARYINGOUT", "out", -1)
		vertShader = strings.Replace(vertShader, "TEXTURE2D", "texture", -1)
	}

	if runtime.GOARCH == "js" || runtime.GOOS == "android" {
		fragShader = strings.Replace(fragShader, "#version 330\n", "precision mediump float;\n", -1)
		fragShader = strings.Replace(fragShader, "VARYINGIN", "varying", -1)
		fragShader = strings.Replace(fragShader, "COLOROUT", "", -1)
		fragShader = strings.Replace(fragShader, "FRAGCOLOR", "gl_FragColor", -1)
		fragShader = strings.Replace(fragShader, "TEXTURE2D", "texture2D", -1)
	} else {
		fragShader = strings.Replace(fragShader, "ATTRIBUTE", "attribute", -1)
		fragShader = strings.Replace(fragShader, "VARYINGIN", "in", -1)
		fragShader = strings.Replace(fragShader, "COLOROUT", "out vec4 colourOut;", -1)
		fragShader = strings.Replace(fragShader, "FRAGCOLOR", "colourOut", -1)
		fragShader = strings.Replace(fragShader, "TEXTURE2D", "texture", -1)
	}

	// create the vertex shader
	var vs = Gl.CreateShader(Gl.VERTEX_SHADER)
	ShaderSource(vs, vertShader)
	Gl.CompileShader(vs)

	if !Gl.GetShaderiv(vs, Gl.COMPILE_STATUS) {
		return nil, fmt.Errorf("Failed to compile the vertex shader!\n%s", GetShaderInfoLog(vs))
	}

	// create the fragment shader
	var fs = Gl.CreateShader(Gl.FRAGMENT_SHADER)
	ShaderSource(fs, fragShader)
	Gl.CompileShader(fs)

	if !Gl.GetShaderiv(fs, Gl.COMPILE_STATUS) {
		return nil, fmt.Errorf("Failed to compile the fragment shader!\n%s", GetShaderInfoLog(fs))
	}

	// attach the shaders to the program and link
	Gl.AttachShader(prog, vs)
	Gl.AttachShader(prog, fs)

	Gl.LinkProgram(prog)

	if !Gl.GetProgramParameterb(p.prog, Gl.LINK_STATUS) {
		return nil, fmt.Errorf("Failed to link the program!\n%s", p.GetProgramInfoLog())
	}

	// at this point the shaders can be deleted
	Gl.DeleteShader(vs)
	Gl.DeleteShader(fs)

	// insert in prog cache
	sProgCache[hash] = &ProgramCache{
		ReleasingReferenceCount: NewReferenceCount(),
		program:                 p,
	}

	//fmt.Printf("Create %s\n", p.hash)
	return p, nil
}

// GetAttribs ...
func (p *GPProgram) GetAttribs() (attribs map[string]int) {
	attribs = make(map[string]int)
	for _, attr := range p.attribs {
		attribs[attr] = Gl.GetAttribLocation(p.prog, attr)
	}
	return attribs
}
