package glplus

import (
	"fmt"

	gl "github.com/go-gl/gl/v4.1-core/gl"
)

// Program ...
type Program struct {
	prog uint32

	uniforms map[string]int32
}

// Handle ...
func (p *Program) Handle() uint32 {
	return p.prog
}

// DeleteProgram ...
func (p *Program) DeleteProgram() {
	gl.DeleteProgram(p.prog)
}

// GetProgramInfoLog ...
func (p *Program) GetProgramInfoLog() string {
	var logSize int32
	gl.GetProgramiv(p.prog, gl.INFO_LOG_LENGTH, &logSize)
	var infoLog = make([]uint8, logSize)
	gl.GetProgramInfoLog(p.prog, logSize, &logSize, &infoLog[0])
	return gl.GoStr(&infoLog[0])
}

// ValidateProgram ...
func (p *Program) ValidateProgram() error {
	gl.ValidateProgram(p.prog)

	var success int32 = gl.FALSE
	gl.GetProgramiv(p.prog, gl.VALIDATE_STATUS, &success)

	if success == gl.FALSE {
		return fmt.Errorf("Failed to validate the program!\n%s", p.GetProgramInfoLog())
	}
	return nil
}

// GetUniformLocation ...
func (p *Program) GetUniformLocation(s string) int32 {
	var res int32
	var ok bool
	if res, ok = p.uniforms[s]; ok {
		return res
	}

	var sarray = StringToArray(s)
	res = gl.GetUniformLocation(p.prog, &sarray[0])
	p.uniforms[s] = res
	return res
}

// GetAttribLocation ...
func (p *Program) GetAttribLocation(s string) uint32 {
	var sarray = StringToArray(s)
	return uint32(gl.GetAttribLocation(p.prog, &sarray[0]))
}

// UseProgram ...
func (p *Program) UseProgram() {
	gl.UseProgram(p.prog)
}

// UnuseProgram ...
func (p *Program) UnuseProgram() {
	gl.UseProgram(0)
}

// ProgramUniform1f ...
func (p *Program) ProgramUniform1f(uniform string, value float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	gl.ProgramUniform1f(p.prog, uniformloc, value)
}

// ProgramUniform2f ...
func (p *Program) ProgramUniform2f(uniform string, v0 float32, v1 float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	gl.ProgramUniform2f(p.prog, uniformloc, v0, v1)
}

// ProgramUniform4fv ...
func (p *Program) ProgramUniform4fv(uniform string, value [4]float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	gl.ProgramUniform4fv(p.prog, uniformloc, 1, &value[0])
}

// ProgramUniform3fv ...
func (p *Program) ProgramUniform3fv(uniform string, value [3]float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	gl.ProgramUniform3fv(p.prog, uniformloc, 1, &value[0])
}

// ProgramUniform1i ...
func (p *Program) ProgramUniform1i(uniform string, value int32) {
	var uniformloc = p.GetUniformLocation(uniform)
	gl.ProgramUniform1i(p.prog, uniformloc, value)
}

// ProgramUniformMatrix4fv ...
func (p *Program) ProgramUniformMatrix4fv(uniform string, matrix [16]float32) {
	var uniformloc = p.GetUniformLocation(uniform)
	gl.ProgramUniformMatrix4fv(p.prog, uniformloc, 1, false, &matrix[0])
}

// GetShaderInfoLog ...
func GetShaderInfoLog(shader uint32) string {
	var logSize int32
	gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logSize)
	var infoLog = make([]uint8, logSize)
	gl.GetShaderInfoLog(shader, logSize, &logSize, &infoLog[0])
	return gl.GoStr(&infoLog[0])
}

// ShaderSource ...
func ShaderSource(shader uint32, src string) {
	//var srcArray []uint8 = StringToArray(src)
	//var ptr *uint8 = &srcArray[0]
	src += "\x00"
	csources, free := gl.Strs(src)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
}

// LoadShaderProgram ... loads shader objects and then attaches them to a program
func LoadShaderProgram(vertShader string, fragShader string, attribs []string) (*Program, error) {
	// create the program
	var prog = gl.CreateProgram()
	var p = &Program{
		prog:     prog,
		uniforms: make(map[string]int32),
	}

	// create the vertex shader
	var vs = gl.CreateShader(gl.VERTEX_SHADER)
	ShaderSource(vs, vertShader)
	gl.CompileShader(vs)

	var success int32 = gl.FALSE
	gl.GetShaderiv(vs, gl.COMPILE_STATUS, &success)

	if success == gl.FALSE {
		return nil, fmt.Errorf("Failed to compile the vertex shader!\n%s", GetShaderInfoLog(vs))
	}

	// create the fragment shader
	var fs = gl.CreateShader(gl.FRAGMENT_SHADER)
	ShaderSource(fs, fragShader)
	gl.CompileShader(fs)

	success = gl.FALSE
	gl.GetShaderiv(fs, gl.COMPILE_STATUS, &success)

	if success == gl.FALSE {
		return nil, fmt.Errorf("Failed to compile the fragment shader!\n%s", GetShaderInfoLog(fs))
	}

	// attach the shaders to the program and link
	gl.AttachShader(prog, vs)
	gl.AttachShader(prog, fs)

	for i := 1; i <= len(attribs); i++ {
		var attrarray = StringToArray(attribs[i-1])
		gl.BindAttribLocation(prog, uint32(i), &attrarray[0])
	}

	gl.LinkProgram(prog)

	success = gl.FALSE
	gl.GetProgramiv(prog, gl.LINK_STATUS, &success)

	if success == gl.FALSE {
		return nil, fmt.Errorf("Failed to link the program!\n%s", p.GetProgramInfoLog())
	}

	// at this point the shaders can be deleted
	gl.DeleteShader(vs)
	gl.DeleteShader(fs)

	return p, nil
}
