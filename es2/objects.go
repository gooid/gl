// openGL 常用的类型对象化
// 如 Attrib Uniform Program Shader Texture 等
package gl

import (
	"errors"
	"unsafe"
)

// Attrib
type Attrib uint32

func (a Attrib) c() uint32 {
	return uint32(a)
}

// VertexAttrib1f writes a float vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func (indx Attrib) F(x float32) {
	VertexAttrib1f(uint32(indx), x)
}

// VertexAttrib1fv writes a float vertex attribute.
func (indx Attrib) Fv(values [1]float32) {
	VertexAttrib1fv(uint32(indx), &values[0])
}

// VertexAttrib2f writes a vec2 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func (indx Attrib) F2(x float32, y float32) {
	VertexAttrib2f(uint32(indx), x, y)
}

// VertexAttrib2fv writes a vec2 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func (indx Attrib) F2v(values [2]float32) {
	VertexAttrib2fv(uint32(indx), &values[0])
}

// VertexAttrib3f writes a vec3 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func (indx Attrib) F3(x float32, y float32, z float32) {
	VertexAttrib3f(uint32(indx), x, y, z)
}

// VertexAttrib3f writes a vec3 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func (indx Attrib) F3v(values [3]float32) {
	VertexAttrib3fv(uint32(indx), &values[0])
}

// VertexAttrib4f writes a vec4 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func (indx Attrib) F4(x float32, y float32, z float32, w float32) {
	VertexAttrib4f(uint32(indx), x, y, z, w)
}

// VertexAttrib4f writes a vec4 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func (indx Attrib) F4v(values [4]float32) {
	VertexAttrib4fv(uint32(indx), &values[0])
}

// GetVertexAttribf reads the float value of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func (src Attrib) GetF(pname uint32) float32 {
	var result float32
	GetVertexAttribfv(src.c(), uint32(pname), &result)
	return result
}

// GetVertexAttribfv reads float values of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func (src Attrib) GetFv(pname uint32, dst []float32) {
	GetVertexAttribfv(src.c(), uint32(pname), &dst[0])
}

// GetVertexAttribi reads the int value of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func (src Attrib) Get(pname uint32) int32 {
	var result int32
	GetVertexAttribiv(src.c(), uint32(pname), &result)
	return result
}

// GetVertexAttribiv reads int values of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func (src Attrib) GetIv(pname uint32, dst []int32) {
	GetVertexAttribiv(src.c(), uint32(pname), &dst[0])
}

// VertexAttribPointer uses a bound buffer to define vertex attribute data.
//
// Direct use of VertexAttribPointer to load data into OpenGL is not
// supported via the Go bindings. Instead, use BindBuffer with an
// ARRAY_BUFFER and then fill it using BufferData.
//
// The size argument specifies the number of components per attribute,
// between 1-4. The stride argument specifies the byte offset between
// consecutive vertex attributes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttribPointer.xhtml
func (indx Attrib) Pointer(size uint, typ uint32, normalized bool, stride int, pointer interface{}) {
	VertexAttribPointer(uint32(indx), int32(size), typ, normalized, int32(stride), Ptr(pointer))
}

// EnableVertexAttribArray enables a vertex attribute array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glEnableVertexAttribArray.xhtml
func (indx Attrib) Enable() {
	EnableVertexAttribArray(uint32(indx))
}

// DisableVertexAttribArray disables a vertex attribute array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDisableVertexAttribArray.xhtml
func (indx Attrib) Disable() {
	DisableVertexAttribArray(uint32(indx))
}

// Uniform
type Uniform int32

func (u Uniform) c() int32 {
	return int32(u)
}

// Uniform1f writes a float uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) F(v float32) {
	Uniform1f(dst.c(), v)
}

// Uniform1fv writes a [len(src)]float uniform array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) Fv(src []float32) {
	Uniform1fv(dst.c(), int32(len(src)), &src[0])
}

// Uniform1i writes an int uniform variable.
//
// Uniform1i and Uniform1iv are the only two functions that may be used
// to load uniform variables defined as sampler types. Loading samplers
// with any other function will result in a INVALID_OPERATION error.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) I(v int) {
	Uniform1i(dst.c(), int32(v))
}

// Uniform1iv writes a int uniform array of len(src) elements.
//
// Uniform1i and Uniform1iv are the only two functions that may be used
// to load uniform variables defined as sampler types. Loading samplers
// with any other function will result in a INVALID_OPERATION error.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) Iv(src []int32) {
	Uniform1iv(dst.c(), int32(len(src)), &src[0])
}

// Uniform2f writes a vec2 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) F2(v0, v1 float32) {
	Uniform2f(dst.c(), v0, v1)
}

// Uniform2fv writes a vec2 uniform array of len(src)/2 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) F2v(src []float32) {
	Uniform2fv(dst.c(), int32(len(src)/2), &src[0])
}

// Uniform2i writes an ivec2 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) I2(v0, v1 int) {
	Uniform2i(dst.c(), int32(v0), int32(v1))
}

// Uniform2iv writes an ivec2 uniform array of len(src)/2 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) I2v(src []int32) {
	Uniform2iv(dst.c(), int32(len(src)/2), &src[0])
}

// Uniform3f writes a vec3 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) F3(v0, v1, v2 float32) {
	Uniform3f(dst.c(), v0, v1, v2)
}

// Uniform3fv writes a vec3 uniform array of len(src)/3 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) F3v(src []float32) {
	Uniform3fv(dst.c(), int32(len(src)/3), &src[0])
}

// Uniform3i writes an ivec3 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) I3(v0, v1, v2 int32) {
	Uniform3i(dst.c(), v0, v1, v2)
}

// Uniform3iv writes an ivec3 uniform array of len(src)/3 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) I3v(src []int32) {
	Uniform3iv(dst.c(), int32(len(src)/3), &src[0])
}

// Uniform4f writes a vec4 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) F4(v0, v1, v2, v3 float32) {
	Uniform4f(dst.c(), v0, v1, v2, v3)
}

// Uniform4fv writes a vec4 uniform array of len(src)/4 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) F4v(src []float32) {
	Uniform4fv(dst.c(), int32(len(src)/4), &src[0])
}

// Uniform4i writes an ivec4 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) I4(v0, v1, v2, v3 int32) {
	Uniform4i(dst.c(), v0, v1, v2, v3)
}

// Uniform4i writes an ivec4 uniform array of len(src)/4 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) I4v(src []int32) {
	Uniform4iv(dst.c(), int32(len(src)/4), &src[0])
}

// UniformMatrix2fv writes 2x2 matrices. Each matrix uses four
// float32 values, so the number of matrices written is len(src)/4.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) Matrix2fv(src []float32) {
	UniformMatrix2fv(dst.c(), int32(len(src)/(2*2)), false, &src[0])
}

// UniformMatrix3fv writes 3x3 matrices. Each matrix uses nine
// float32 values, so the number of matrices written is len(src)/9.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) Matrix3fv(src []float32) {
	UniformMatrix3fv(dst.c(), int32(len(src)/(3*3)), false, &src[0])
}

// UniformMatrix4fv writes 4x4 matrices. Each matrix uses 16
// float32 values, so the number of matrices written is len(src)/16.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func (dst Uniform) Matrix4fv(src []float32) {
	UniformMatrix4fv(dst.c(), int32(len(src)/(4*4)), false, &src[0])
}

///
type Buffer uint32

func (o Buffer) c() uint32 {
	return uint32(o)
}

// CreateBuffer creates a buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenBuffers.xhtml
func CreateBuffer() Buffer {
	var b uint32
	GenBuffers(1, &b)
	return Buffer(b)
}

// DeleteBuffer deletes the given buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteBuffers.xhtml
func (buffer *Buffer) Delete() {
	b := buffer.c()
	DeleteBuffers(1, &b)
	*buffer = 0
}

// BindBuffer binds a buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindBuffer.xhtml
func (buffer Buffer) Bind(target uint32) {
	BindBuffer(target, buffer.c())
}

// Remove buffer binding
func (buffer Buffer) Unbind(target uint32) {
	BindBuffer(target, 0)
}

// IsBuffer reports if b is a valid buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsBuffer.xhtml
func (buffer Buffer) IsValid() bool {
	return IsBuffer(buffer.c())
}

///
type RenderBuffer uint32

func (o RenderBuffer) c() uint32 {
	return uint32(o)
}

// CreateRenderbuffer create a renderbuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenRenderbuffers.xhtml
func CreateRenderbuffer() RenderBuffer {
	var v uint32
	GenRenderbuffers(1, &v)
	return RenderBuffer(v)
}

// DeleteRenderbuffer deletes the given render buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteRenderbuffers.xhtml
func (rb *RenderBuffer) Delete() {
	v := rb.c()
	DeleteRenderbuffers(1, &v)
	*rb = 0
}

// BindRenderbuffer binds a render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindRenderbuffer.xhtml
func (rb RenderBuffer) Bind() {
	BindRenderbuffer(RENDERBUFFER, rb.c())
}

// Unbind the render buffer
func (rb RenderBuffer) Unbind() {
	BindRenderbuffer(RENDERBUFFER, 0)
}

// IsRenderbuffer reports if rb is a valid render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsRenderbuffer.xhtml
func (rb RenderBuffer) IsValid() bool {
	return IsRenderbuffer(rb.c())
}

///
type FrameBuffer uint32

func (o FrameBuffer) c() uint32 {
	return uint32(o)
}

// CreateFramebuffer creates a framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenFramebuffers.xhtml
func CreateFramebuffer() FrameBuffer {
	var v uint32
	GenFramebuffers(1, &v)
	return FrameBuffer(v)
}

// DeleteFramebuffer deletes the given framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteFramebuffers.xhtml
func (fb *FrameBuffer) Delete() {
	v := fb.c()
	DeleteFramebuffers(1, &v)
	*fb = 0
}

// BindFramebuffer binds a framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindFramebuffer.xhtml
func (fb FrameBuffer) Bind(target uint32) {
	BindFramebuffer(target, fb.c())
}

// Unbinds the framebuffer.
func (fb FrameBuffer) Unbind(target uint32) {
	BindFramebuffer(target, 0)
}

// IsFramebuffer reports if fb is a valid frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsFramebuffer.xhtml
func (fb FrameBuffer) IsValid() bool {
	return IsFramebuffer(fb.c())
}

// FramebufferTexture2D attaches the t to the current frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFramebufferTexture2D.xhtml
func (fb FrameBuffer) Texture2D(attachment, texTarget uint32, t Texture, level int) {
	FramebufferTexture2D(fb.c(), uint32(attachment), uint32(texTarget), t.c(), int32(level))
}

// Texture
type Texture uint32

func (o Texture) c() uint32 {
	return uint32(o)
}

// Create single texture object
func CreateTexture() Texture {
	var v uint32
	GenTextures(1, &v)
	return Texture(v)
}

// Delete texture object
func (texture *Texture) Delete() {
	v := texture.c()
	DeleteTextures(1, &v)
	*texture = 0
}

// BindTexture binds a texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindTexture.xhtml
func (texture Texture) Bind(target uint32) {
	BindTexture(target, texture.c())
}

// Unbind this texture.
func (texture Texture) Unbind(target uint32) {
	BindTexture(target, 0)
}

// IsTexture reports if t is a valid texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsTexture.xhtml
func (t Texture) IsValid() bool {
	return IsTexture(t.c())
}

// Shader
type shader uint32

func (o shader) c() uint32 {
	return uint32(o)
}

// DeleteShader deletes shader s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteShader.xhtml
func (shader *shader) Delete() {
	DeleteShader(shader.c())
	*shader = 0
}

// GetShaderInfoLog returns the information log for a shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderInfoLog.xhtml
func (shader shader) GetInfoLog() string {
	length := shader.Get(INFO_LOG_LENGTH)

	// length is buffer size including null character
	if length > 1 {
		buf := make([]byte, length+1)
		GetShaderInfoLog(shader.c(), length, &length, (*uint8)(unsafe.Pointer(&buf[0])))
		return string(buf[:length])
	}
	return ""
}

// GetShaderSource returns source code of shader s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderSource.xhtml
func (shader shader) GetSource() string {
	length := shader.Get(SHADER_SOURCE_LENGTH)

	if length > 1 {
		buf := make([]byte, length+1)
		GetShaderSource(shader.c(), length, &length, (*uint8)(unsafe.Pointer(&buf[0])))
		return string(buf[:length])
	}
	return ""
}

// ShaderSource sets the source code of s to the given source code.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glShaderSource.xhtml
func (shader shader) Source(source ...string) {
	count := int32(len(source))
	lengths := make([]int32, count)

	for i, s := range source {
		lengths[i] = int32(len(s))
	}
	xstring, free := Strs(source...)
	defer free()
	ShaderSource(shader.c(), count, xstring, &lengths[0])
}

// CompileShader compiles the source code of s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompileShader.xhtml
func (shader shader) Compile() { CompileShader(shader.c()) }

// GetShaderiv returns a parameter value for a shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderiv.xhtml
func (shader shader) Get(pname uint32) int32 {
	var v int32
	GetShaderiv(shader.c(), pname, &v)
	return v
}

// IsShader reports if s is valid shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsShader.xhtml
func (shader shader) IsValid() bool {
	return IsShader(shader.c())
}

// Program
type Program uint32

func (p Program) c() uint32 {
	return uint32(p)
}

// DeleteProgram deletes the given program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteProgram.xhtml
func (p *Program) Delete() {
	DeleteProgram(p.c())
	*p = 0
}

// AttachShader attaches a shader to a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glAttachShader.xhtml
func (p Program) AttachShader(shader shader) {
	AttachShader(p.c(), shader.c())
}

// GetAttachedShaders returns the shader objects attached to program p.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetAttachedShaders.xhtml
func (p Program) GetAttachedShaders() []shader {
	cnt := p.Get(ATTACHED_SHADERS)
	outs := make([]shader, cnt+1)
	GetAttachedShaders(p.c(), cnt, &cnt, (*uint32)(unsafe.Pointer(&outs[0])))
	return outs[:cnt]
}

// DetachShader detaches the shader s from the program p.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDetachShader.xhtml
func (p Program) DetachShader(shader shader) {
	DetachShader(p.c(), shader.c())
}

// LinkProgram links the specified program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glLinkProgram.xhtml
func (p Program) Link() { LinkProgram(p.c()) }

// ValidateProgram checks to see whether the executables contained in
// program can execute given the current OpenGL state.
//
// Typically only used for debugging.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glValidateProgram.xhtml
func (p Program) Validate() { ValidateProgram(p.c()) }

// UseProgram sets the active program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUseProgram.xhtml
func (p Program) Use() { UseProgram(p.c()) }

// Unuse the active program.
func (p Program) Unuse() { UseProgram(uint32(0)) }

// GetProgramInfoLog returns the information log for a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetProgramInfoLog.xhtml
func (p Program) GetInfoLog() string {
	length := p.Get(INFO_LOG_LENGTH)

	// length is buffer size including null character
	if length > 1 {
		buf := make([]byte, length+1)
		GetProgramInfoLog(p.c(), length, &length, (*uint8)(unsafe.Pointer(&buf[0])))
		return string(buf[:length])
	}
	return ""
}

// GetProgramiv returns a parameter value for a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetProgramiv.xhtml
func (p Program) Get(pname uint32) int32 {
	var v int32
	GetProgramiv(p.c(), pname, &v)
	return v
}

// GetActiveUniform returns details about an active uniform variable.
// A value of 0 for index selects the first active uniform variable.
// Permissible values for index range from 0 to the number of active
// uniform variables minus 1.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniform.xhtml
func (p Program) GetActiveUniform(index int) (int32, uint32, string) {
	// Maximum length of active uniform name in program
	bufSize := p.Get(ACTIVE_UNIFORM_MAX_LENGTH)
	var size int32
	var xtype uint32
	buf := make([]byte, bufSize+1)
	GetActiveUniform(p.c(), uint32(index),
		bufSize, &bufSize, &size, &xtype, (*uint8)(unsafe.Pointer(&buf[0])))
	return size, xtype, string(buf[:bufSize])
}

// GetUniformiv returns the float values of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniform.xhtml
func (p Program) GetUniformiv(location Uniform, params []int32) {
	GetUniformiv(p.c(), location.c(), &params[0])
}

// GetUniformfv returns the float values of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniform.xhtml
func (p Program) GetUniformfv(location Uniform, params []float32) {
	GetUniformfv(p.c(), location.c(), &params[0])
}

// GetUniformLocation returns the location of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniformLocation.xhtml
func (p Program) GetUniformLocation(name string) Uniform {
	return Uniform(GetUniformLocation(p.c(), Str(name+"\x00")))
}

// GetActiveAttrib returns details about an active attribute variable.
// A value of 0 for index selects the first active attribute variable.
// Permissible values for index range from 0 to the number of active
// attribute variables minus 1.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveAttrib.xhtml
func (p Program) GetActiveAttrib(index int) (int32, uint32, string) {
	// Maximum length of active uniform name in program
	bufSize := p.Get(ACTIVE_ATTRIBUTE_MAX_LENGTH)
	var size int32
	var xtype uint32
	buf := make([]byte, bufSize+1)
	GetActiveAttrib(p.c(), uint32(index),
		bufSize, &bufSize, &size, &xtype, (*uint8)(unsafe.Pointer(&buf[0])))
	return size, xtype, string(buf[:bufSize])
}

// GetAttribLocation returns the location of an attribute variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetAttribLocation.xhtml
func (p Program) GetAttribLocation(name string) Attrib {
	return Attrib(GetAttribLocation(p.c(), Str(name+"\x00")))
}

// BindAttribLocation binds a vertex attribute index with a named
// variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindAttribLocation.xhtml
func (p Program) BindAttribLocation(index Attrib, name string) {
	BindAttribLocation(p.c(), index.c(), Str(name+"\x00"))
}

// IsProgram reports if p is a valid program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsProgram.xhtml
func (p Program) IsValid() bool {
	return IsProgram(p.c())
}

//
func compileShader(t uint32, strs ...string) (shader, error) {
	s := shader(CreateShader(t))
	if !s.IsValid() {
		return shader(0), errors.New(GoStr(GetString(GetError())))
	}

	s.Source(strs...)
	s.Compile()
	if s.Get(COMPILE_STATUS) == FALSE {
		s.Delete()
		return shader(0), errors.New(s.GetInfoLog())
	}
	return s, nil
}

// NewProgram
func NewProgram(vertexSrc, fragmentSrc []string) (Program, error) {
	vert, err := compileShader(VERTEX_SHADER, vertexSrc...)
	if err != nil {
		return Program(0), err
	}
	defer vert.Delete() // 如果 Attach，则只是标记为删除

	frag, err := compileShader(FRAGMENT_SHADER, fragmentSrc...)
	if err != nil {
		return Program(0), err
	}
	defer frag.Delete()

	p := Program(CreateProgram())
	if !p.IsValid() {
		return p, errors.New(GoStr(GetString(GetError())))
	}

	p.AttachShader(vert)
	p.AttachShader(frag)
	p.Link()
	if p.Get(LINK_STATUS) == FALSE {
		p.DetachShader(vert)
		p.DetachShader(frag)
		p.Delete()
		return Program(0), errors.New(p.GetInfoLog())
	}
	return p, nil
}

//
type Shader struct {
	Program
}

func (s *Shader) Handle() uint32 {
	return s.c()
}

func NewShader(vertexSrc, fragmentSrc []string) *Shader {
	p, _ := NewProgram(vertexSrc, fragmentSrc)
	return &Shader{Program: p}
}
