package gl

import (
	"log"
)

type Shader struct {
	shaderHandle uint32
	vertHandle   uint32
	fragHandle   uint32
}

// If you get an error please report on github. You may try different GL context version or GLSL version.
func checkShader(handle uint32, desc string) bool {
	status, logLength := int32(0), int32(0)
	GetShaderiv(handle, COMPILE_STATUS, &status)
	if status == FALSE {
		GetShaderiv(handle, INFO_LOG_LENGTH, &logLength)
		logStr := ""
		if logLength > 0 {
			buf := make([]byte, logLength+1)
			GetShaderInfoLog(handle, logLength+1, &logLength, (*uint8)(Ptr(&buf[0])))
			logStr = string(buf[:logLength])
		}
		log.Println("ERROR: ", desc, ":", logStr)
	}
	return status == TRUE
}

// If you get an error please report on github. You may try different GL context version or GLSL version.
func checkProgram(handle uint32, desc string) bool {
	status, logLength := int32(0), int32(0)
	GetProgramiv(handle, LINK_STATUS, &status)

	if status == FALSE {
		GetProgramiv(handle, INFO_LOG_LENGTH, &logLength)
		logStr := ""
		if logLength > 0 {
			buf := make([]byte, logLength+1)
			GetProgramInfoLog(handle, logLength, &logLength, (*uint8)(Ptr(&buf[0])))
			logStr = string(buf[:logLength])
		}
		log.Println("ERROR: ", desc, ":", logStr)
	}
	return status == TRUE
}

func NewShader(vertexSrc, fragmentSrc []string) *Shader {
	s := &Shader{}
	// Create shaders
	s.vertHandle = CreateShader(VERTEX_SHADER)
	vstrs, free1 := Strs(vertexSrc...)
	defer free1()
	ShaderSource(s.vertHandle, int32(len(vertexSrc)), vstrs, nil)
	CompileShader(s.vertHandle)
	if !checkShader(s.vertHandle, "vertex shader") {
		return nil
	}

	s.fragHandle = CreateShader(FRAGMENT_SHADER)
	fstrs, free2 := Strs(fragmentSrc...)
	defer free2()
	ShaderSource(s.fragHandle, int32(len(fragmentSrc)), fstrs, nil)
	CompileShader(s.fragHandle)
	if !checkShader(s.fragHandle, "fragment shader") {
		return nil
	}

	s.shaderHandle = CreateProgram()
	AttachShader(s.shaderHandle, s.vertHandle)
	AttachShader(s.shaderHandle, s.fragHandle)
	LinkProgram(s.shaderHandle)
	if !checkProgram(s.shaderHandle, "shader program") {
		return nil
	}
	return s
}

func (s *Shader) Delete() {
	if (s.shaderHandle != 0) && (s.vertHandle != 0) {
		DetachShader(s.shaderHandle, s.vertHandle)
	}
	if s.vertHandle != 0 {
		DeleteShader(s.vertHandle)
	}
	s.vertHandle = 0

	if (s.shaderHandle != 0) && (s.fragHandle != 0) {
		DetachShader(s.shaderHandle, s.fragHandle)
	}
	if s.fragHandle != 0 {
		DeleteShader(s.fragHandle)
	}
	s.fragHandle = 0

	if s.shaderHandle != 0 {
		DeleteProgram(s.shaderHandle)
	}
	s.shaderHandle = 0
}

func (s *Shader) Handle() uint32 {
	return s.shaderHandle
}
