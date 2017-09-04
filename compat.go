package glplus

import (
	"log"

	gl "github.com/go-gl/gl/v4.1-core/gl"
)

type Context struct {
	ARRAY_BUFFER                                 int
	ARRAY_BUFFER_BINDING                         int
	ATTACHED_SHADERS                             int
	BACK                                         int
	BLEND                                        int
	BLEND_COLOR                                  int
	BLEND_DST_ALPHA                              int
	BLEND_DST_RGB                                int
	BLEND_EQUATION                               int
	BLEND_EQUATION_ALPHA                         int
	BLEND_EQUATION_RGB                           int
	BLEND_SRC_ALPHA                              int
	BLEND_SRC_RGB                                int
	BLUE_BITS                                    int
	BOOL                                         int
	BOOL_VEC2                                    int
	BOOL_VEC3                                    int
	BOOL_VEC4                                    int
	BROWSER_DEFAULT_WEBGL                        int
	BUFFER_SIZE                                  int
	BUFFER_USAGE                                 int
	BYTE                                         int
	CCW                                          int
	CLAMP_TO_EDGE                                int
	CLAMP_TO_BORDER                              int
	COLOR_ATTACHMENT0                            int
	COLOR_BUFFER_BIT                             int
	COLOR_CLEAR_VALUE                            int
	COLOR_WRITEMASK                              int
	COMPILE_STATUS                               uint32
	COMPRESSED_TEXTURE_FORMATS                   int
	CONSTANT_ALPHA                               int
	CONSTANT_COLOR                               int
	CONTEXT_LOST_WEBGL                           int
	CULL_FACE                                    int
	CULL_FACE_MODE                               int
	CURRENT_PROGRAM                              int
	CURRENT_VERTEX_ATTRIB                        int
	CW                                           int
	DECR                                         int
	DECR_WRAP                                    int
	DELETE_STATUS                                int
	DEPTH_ATTACHMENT                             int
	DEPTH_BITS                                   int
	DEPTH_BUFFER_BIT                             int
	DEPTH_CLEAR_VALUE                            int
	DEPTH_COMPONENT                              int
	DEPTH_COMPONENT16                            int
	DEPTH_FUNC                                   int
	DEPTH_RANGE                                  int
	DEPTH_STENCIL                                int
	DEPTH_STENCIL_ATTACHMENT                     int
	DEPTH_TEST                                   int
	DEPTH_WRITEMASK                              int
	DITHER                                       int
	DONT_CARE                                    int
	DST_ALPHA                                    int
	DST_COLOR                                    int
	DYNAMIC_DRAW                                 int
	ELEMENT_ARRAY_BUFFER                         int
	ELEMENT_ARRAY_BUFFER_BINDING                 int
	EQUAL                                        int
	FASTEST                                      int
	FLOAT                                        int
	FLOAT_MAT2                                   int
	FLOAT_MAT3                                   int
	FLOAT_MAT4                                   int
	FLOAT_VEC2                                   int
	FLOAT_VEC3                                   int
	FLOAT_VEC4                                   int
	FRAGMENT_SHADER                              int
	FRAMEBUFFER                                  int
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           int
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         int
	FRAMEBUFFER_BINDING                          int
	FRAMEBUFFER_COMPLETE                         int
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT            int
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS            int
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    int
	FRAMEBUFFER_UNSUPPORTED                      int
	FRONT                                        int
	FRONT_AND_BACK                               int
	FRONT_FACE                                   int
	FUNC_ADD                                     int
	FUNC_REVERSE_SUBTRACT                        int
	FUNC_SUBTRACT                                int
	GENERATE_MIPMAP_HINT                         int
	GEQUAL                                       int
	GREATER                                      int
	GREEN_BITS                                   int
	HIGH_FLOAT                                   int
	HIGH_INT                                     int
	INCR                                         int
	INCR_WRAP                                    int
	INFO_LOG_LENGTH                              uint32
	INT                                          int
	INT_VEC2                                     int
	INT_VEC3                                     int
	INT_VEC4                                     int
	INVALID_ENUM                                 int
	INVALID_FRAMEBUFFER_OPERATION                int
	INVALID_OPERATION                            int
	INVALID_VALUE                                int
	INVERT                                       int
	KEEP                                         int
	LEQUAL                                       int
	LESS                                         int
	LINEAR                                       int
	LINEAR_MIPMAP_LINEAR                         int
	LINEAR_MIPMAP_NEAREST                        int
	LINES                                        int
	LINE_LOOP                                    int
	LINE_STRIP                                   int
	LINE_WIDTH                                   int
	LINK_STATUS                                  int
	LOW_FLOAT                                    int
	LOW_INT                                      int
	LUMINANCE                                    int
	LUMINANCE_ALPHA                              int
	MAX_COMBINED_TEXTURE_IMAGE_UNITS             int
	MAX_CUBE_MAP_TEXTURE_SIZE                    int
	MAX_FRAGMENT_UNIFORM_VECTORS                 int
	MAX_RENDERBUFFER_SIZE                        int
	MAX_TEXTURE_IMAGE_UNITS                      int
	MAX_TEXTURE_SIZE                             int
	MAX_VARYING_VECTORS                          int
	MAX_VERTEX_ATTRIBS                           int
	MAX_VERTEX_TEXTURE_IMAGE_UNITS               int
	MAX_VERTEX_UNIFORM_VECTORS                   int
	MAX_VIEWPORT_DIMS                            int
	MEDIUM_FLOAT                                 int
	MEDIUM_INT                                   int
	MIRRORED_REPEAT                              int
	MULTISAMPLE                                  int
	NEAREST                                      int
	NEAREST_MIPMAP_LINEAR                        int
	NEAREST_MIPMAP_NEAREST                       int
	NEVER                                        int
	NICEST                                       int
	NONE                                         int
	NOTEQUAL                                     int
	NO_ERROR                                     int
	NUM_COMPRESSED_TEXTURE_FORMATS               int
	ONE                                          int
	ONE_MINUS_CONSTANT_ALPHA                     int
	ONE_MINUS_CONSTANT_COLOR                     int
	ONE_MINUS_DST_ALPHA                          int
	ONE_MINUS_DST_COLOR                          int
	ONE_MINUS_SRC_ALPHA                          int
	ONE_MINUS_SRC_COLOR                          int
	OUT_OF_MEMORY                                int
	PACK_ALIGNMENT                               int
	POINTS                                       int
	POLYGON_OFFSET_FACTOR                        int
	POLYGON_OFFSET_FILL                          int
	POLYGON_OFFSET_UNITS                         int
	RED_BITS                                     int
	RENDERBUFFER                                 int
	RENDERBUFFER_ALPHA_SIZE                      int
	RENDERBUFFER_BINDING                         int
	RENDERBUFFER_BLUE_SIZE                       int
	RENDERBUFFER_DEPTH_SIZE                      int
	RENDERBUFFER_GREEN_SIZE                      int
	RENDERBUFFER_HEIGHT                          int
	RENDERBUFFER_INTERNAL_FORMAT                 int
	RENDERBUFFER_RED_SIZE                        int
	RENDERBUFFER_STENCIL_SIZE                    int
	RENDERBUFFER_WIDTH                           int
	RENDERER                                     int
	REPEAT                                       int
	REPLACE                                      int
	RGB                                          int
	RGB5_A1                                      int
	RGB565                                       int
	RGBA                                         int
	RGBA4                                        int
	SAMPLER_2D                                   int
	SAMPLER_CUBE                                 int
	SAMPLES                                      int
	SAMPLE_ALPHA_TO_COVERAGE                     int
	SAMPLE_BUFFERS                               int
	SAMPLE_COVERAGE                              int
	SAMPLE_COVERAGE_INVERT                       int
	SAMPLE_COVERAGE_VALUE                        int
	SCISSOR_BOX                                  int
	SCISSOR_TEST                                 int
	SHADER_COMPILER                              int
	SHADER_SOURCE_LENGTH                         int
	SHADER_TYPE                                  int
	SHADING_LANGUAGE_VERSION                     int
	SHORT                                        int
	SRC_ALPHA                                    int
	SRC_ALPHA_SATURATE                           int
	SRC_COLOR                                    int
	STATIC_DRAW                                  int
	STENCIL_ATTACHMENT                           int
	STENCIL_BACK_FAIL                            int
	STENCIL_BACK_FUNC                            int
	STENCIL_BACK_PASS_DEPTH_FAIL                 int
	STENCIL_BACK_PASS_DEPTH_PASS                 int
	STENCIL_BACK_REF                             int
	STENCIL_BACK_VALUE_MASK                      int
	STENCIL_BACK_WRITEMASK                       int
	STENCIL_BITS                                 int
	STENCIL_BUFFER_BIT                           int
	STENCIL_CLEAR_VALUE                          int
	STENCIL_FAIL                                 int
	STENCIL_FUNC                                 int
	STENCIL_INDEX                                int
	STENCIL_INDEX8                               int
	STENCIL_PASS_DEPTH_FAIL                      int
	STENCIL_PASS_DEPTH_PASS                      int
	STENCIL_REF                                  int
	STENCIL_TEST                                 int
	STENCIL_VALUE_MASK                           int
	STENCIL_WRITEMASK                            int
	STREAM_DRAW                                  int
	SUBPIXEL_BITS                                int
	TEXTURE                                      int
	TEXTURE0                                     int
	TEXTURE1                                     int
	TEXTURE2                                     int
	TEXTURE3                                     int
	TEXTURE4                                     int
	TEXTURE5                                     int
	TEXTURE6                                     int
	TEXTURE7                                     int
	TEXTURE8                                     int
	TEXTURE9                                     int
	TEXTURE10                                    int
	TEXTURE11                                    int
	TEXTURE12                                    int
	TEXTURE13                                    int
	TEXTURE14                                    int
	TEXTURE15                                    int
	TEXTURE16                                    int
	TEXTURE17                                    int
	TEXTURE18                                    int
	TEXTURE19                                    int
	TEXTURE20                                    int
	TEXTURE21                                    int
	TEXTURE22                                    int
	TEXTURE23                                    int
	TEXTURE24                                    int
	TEXTURE25                                    int
	TEXTURE26                                    int
	TEXTURE27                                    int
	TEXTURE28                                    int
	TEXTURE29                                    int
	TEXTURE30                                    int
	TEXTURE31                                    int
	TEXTURE_2D                                   int
	TEXTURE_BINDING_2D                           int
	TEXTURE_BINDING_CUBE_MAP                     int
	TEXTURE_CUBE_MAP                             int
	TEXTURE_CUBE_MAP_NEGATIVE_X                  int
	TEXTURE_CUBE_MAP_NEGATIVE_Y                  int
	TEXTURE_CUBE_MAP_NEGATIVE_Z                  int
	TEXTURE_CUBE_MAP_POSITIVE_X                  int
	TEXTURE_CUBE_MAP_POSITIVE_Y                  int
	TEXTURE_CUBE_MAP_POSITIVE_Z                  int
	TEXTURE_MAG_FILTER                           int
	TEXTURE_MIN_FILTER                           int
	TEXTURE_WRAP_S                               int
	TEXTURE_WRAP_T                               int
	TRIANGLES                                    int
	TRIANGLE_FAN                                 int
	TRIANGLE_STRIP                               int
	UNPACK_ALIGNMENT                             int
	UNPACK_COLORSPACE_CONVERSION_WEBGL           int
	UNPACK_FLIP_Y_WEBGL                          int
	UNPACK_PREMULTIPLY_ALPHA_WEBGL               int
	UNSIGNED_BYTE                                int
	UNSIGNED_INT                                 int
	UNSIGNED_SHORT                               int
	UNSIGNED_SHORT_4_4_4_4                       int
	UNSIGNED_SHORT_5_5_5_1                       int
	UNSIGNED_SHORT_5_6_5                         int
	VALIDATE_STATUS                              int
	VENDOR                                       int
	VERSION                                      int
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING           int
	VERTEX_ATTRIB_ARRAY_ENABLED                  int
	VERTEX_ATTRIB_ARRAY_NORMALIZED               int
	VERTEX_ATTRIB_ARRAY_POINTER                  int
	VERTEX_ATTRIB_ARRAY_SIZE                     int
	VERTEX_ATTRIB_ARRAY_STRIDE                   int
	VERTEX_ATTRIB_ARRAY_TYPE                     int
	VERTEX_SHADER                                int
	VIEWPORT                                     int
	ZERO                                         int
	TRUE                                         int
}

// NewContext ...
func NewContext() *Context {
	if err := gl.Init(); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &Context{
		ARRAY_BUFFER:               gl.ARRAY_BUFFER,
		ARRAY_BUFFER_BINDING:       gl.ARRAY_BUFFER_BINDING,
		ATTACHED_SHADERS:           gl.ATTACHED_SHADERS,
		BACK:                       gl.BACK,
		BLEND:                      gl.BLEND,
		BLEND_COLOR:                gl.BLEND_COLOR,
		BLEND_DST_ALPHA:            gl.BLEND_DST_ALPHA,
		BLEND_DST_RGB:              gl.BLEND_DST_RGB,
		BLEND_EQUATION:             gl.BLEND_EQUATION,
		BLEND_EQUATION_ALPHA:       gl.BLEND_EQUATION_ALPHA,
		BLEND_EQUATION_RGB:         gl.BLEND_EQUATION_RGB,
		BLEND_SRC_ALPHA:            gl.BLEND_SRC_ALPHA,
		BLEND_SRC_RGB:              gl.BLEND_SRC_RGB,
		BOOL:                       gl.BOOL,
		BOOL_VEC2:                  gl.BOOL_VEC2,
		BOOL_VEC3:                  gl.BOOL_VEC3,
		BOOL_VEC4:                  gl.BOOL_VEC4,
		BUFFER_SIZE:                gl.BUFFER_SIZE,
		BUFFER_USAGE:               gl.BUFFER_USAGE,
		BYTE:                       gl.BYTE,
		CCW:                        gl.CCW,
		CLAMP_TO_EDGE:              gl.CLAMP_TO_EDGE,
		CLAMP_TO_BORDER:            gl.CLAMP_TO_BORDER,
		COLOR_ATTACHMENT0:          gl.COLOR_ATTACHMENT0,
		COLOR_BUFFER_BIT:           gl.COLOR_BUFFER_BIT,
		COLOR_CLEAR_VALUE:          gl.COLOR_CLEAR_VALUE,
		COLOR_WRITEMASK:            gl.COLOR_WRITEMASK,
		COMPILE_STATUS:             gl.COMPILE_STATUS,
		COMPRESSED_TEXTURE_FORMATS: gl.COMPRESSED_TEXTURE_FORMATS,
		CONSTANT_ALPHA:             gl.CONSTANT_ALPHA,
		CONSTANT_COLOR:             gl.CONSTANT_COLOR,
		CULL_FACE:                  gl.CULL_FACE,
		CULL_FACE_MODE:             gl.CULL_FACE_MODE,
		CURRENT_PROGRAM:            gl.CURRENT_PROGRAM,
		CURRENT_VERTEX_ATTRIB:      gl.CURRENT_VERTEX_ATTRIB,
		CW:                           gl.CW,
		DECR:                         gl.DECR,
		DECR_WRAP:                    gl.DECR_WRAP,
		DELETE_STATUS:                gl.DELETE_STATUS,
		DEPTH_ATTACHMENT:             gl.DEPTH_ATTACHMENT,
		DEPTH_BUFFER_BIT:             gl.DEPTH_BUFFER_BIT,
		DEPTH_CLEAR_VALUE:            gl.DEPTH_CLEAR_VALUE,
		DEPTH_COMPONENT:              gl.DEPTH_COMPONENT,
		DEPTH_COMPONENT16:            gl.DEPTH_COMPONENT16,
		DEPTH_FUNC:                   gl.DEPTH_FUNC,
		DEPTH_RANGE:                  gl.DEPTH_RANGE,
		DEPTH_STENCIL:                gl.DEPTH_STENCIL,
		DEPTH_STENCIL_ATTACHMENT:     gl.DEPTH_STENCIL_ATTACHMENT,
		DEPTH_TEST:                   gl.DEPTH_TEST,
		DEPTH_WRITEMASK:              gl.DEPTH_WRITEMASK,
		DITHER:                       gl.DITHER,
		DONT_CARE:                    gl.DONT_CARE,
		DST_ALPHA:                    gl.DST_ALPHA,
		DST_COLOR:                    gl.DST_COLOR,
		DYNAMIC_DRAW:                 gl.DYNAMIC_DRAW,
		ELEMENT_ARRAY_BUFFER:         gl.ELEMENT_ARRAY_BUFFER,
		ELEMENT_ARRAY_BUFFER_BINDING: gl.ELEMENT_ARRAY_BUFFER_BINDING,
		EQUAL:                                        gl.EQUAL,
		FASTEST:                                      gl.FASTEST,
		FLOAT:                                        gl.FLOAT,
		FLOAT_MAT2:                                   gl.FLOAT_MAT2,
		FLOAT_MAT3:                                   gl.FLOAT_MAT3,
		FLOAT_MAT4:                                   gl.FLOAT_MAT4,
		FLOAT_VEC2:                                   gl.FLOAT_VEC2,
		FLOAT_VEC3:                                   gl.FLOAT_VEC3,
		FLOAT_VEC4:                                   gl.FLOAT_VEC4,
		FRAGMENT_SHADER:                              gl.FRAGMENT_SHADER,
		FRAMEBUFFER:                                  gl.FRAMEBUFFER,
		FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:           gl.FRAMEBUFFER_ATTACHMENT_OBJECT_NAME,
		FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:           gl.FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE: gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:         gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL,
		FRAMEBUFFER_BINDING:                          gl.FRAMEBUFFER_BINDING,
		FRAMEBUFFER_COMPLETE:                         gl.FRAMEBUFFER_COMPLETE,
		FRAMEBUFFER_INCOMPLETE_ATTACHMENT:            gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT,
		FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:    gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT,
		FRAMEBUFFER_UNSUPPORTED:                      gl.FRAMEBUFFER_UNSUPPORTED,
		FRONT:                         gl.FRONT,
		FRONT_AND_BACK:                gl.FRONT_AND_BACK,
		FRONT_FACE:                    gl.FRONT_FACE,
		FUNC_ADD:                      gl.FUNC_ADD,
		FUNC_REVERSE_SUBTRACT:         gl.FUNC_REVERSE_SUBTRACT,
		FUNC_SUBTRACT:                 gl.FUNC_SUBTRACT,
		GEQUAL:                        gl.GEQUAL,
		GREATER:                       gl.GREATER,
		HIGH_FLOAT:                    gl.HIGH_FLOAT,
		HIGH_INT:                      gl.HIGH_INT,
		INCR:                          gl.INCR,
		INCR_WRAP:                     gl.INCR_WRAP,
		INFO_LOG_LENGTH:               gl.INFO_LOG_LENGTH,
		INT:                           gl.INT,
		INT_VEC2:                      gl.INT_VEC2,
		INT_VEC3:                      gl.INT_VEC3,
		INT_VEC4:                      gl.INT_VEC4,
		INVALID_ENUM:                  gl.INVALID_ENUM,
		INVALID_FRAMEBUFFER_OPERATION: gl.INVALID_FRAMEBUFFER_OPERATION,
		INVALID_OPERATION:             gl.INVALID_OPERATION,
		INVALID_VALUE:                 gl.INVALID_VALUE,
		INVERT:                        gl.INVERT,
		KEEP:                          gl.KEEP,
		LEQUAL:                        gl.LEQUAL,
		LESS:                          gl.LESS,
		LINEAR:                        gl.LINEAR,
		LINEAR_MIPMAP_LINEAR:          gl.LINEAR_MIPMAP_LINEAR,
		LINEAR_MIPMAP_NEAREST:         gl.LINEAR_MIPMAP_NEAREST,
		LINES:       gl.LINES,
		LINE_LOOP:   gl.LINE_LOOP,
		LINE_STRIP:  gl.LINE_STRIP,
		LINE_WIDTH:  gl.LINE_WIDTH,
		LINK_STATUS: gl.LINK_STATUS,
		LOW_FLOAT:   gl.LOW_FLOAT,
		LOW_INT:     gl.LOW_INT,
		MAX_COMBINED_TEXTURE_IMAGE_UNITS: gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS,
		MAX_CUBE_MAP_TEXTURE_SIZE:        gl.MAX_CUBE_MAP_TEXTURE_SIZE,
		MAX_FRAGMENT_UNIFORM_VECTORS:     gl.MAX_FRAGMENT_UNIFORM_VECTORS,
		MAX_RENDERBUFFER_SIZE:            gl.MAX_RENDERBUFFER_SIZE,
		MAX_TEXTURE_IMAGE_UNITS:          gl.MAX_TEXTURE_IMAGE_UNITS,
		MAX_TEXTURE_SIZE:                 gl.MAX_TEXTURE_SIZE,
		MAX_VARYING_VECTORS:              gl.MAX_VARYING_VECTORS,
		MAX_VERTEX_ATTRIBS:               gl.MAX_VERTEX_ATTRIBS,
		MAX_VERTEX_TEXTURE_IMAGE_UNITS:   gl.MAX_VERTEX_TEXTURE_IMAGE_UNITS,
		MAX_VERTEX_UNIFORM_VECTORS:       gl.MAX_VERTEX_UNIFORM_VECTORS,
		MAX_VIEWPORT_DIMS:                gl.MAX_VIEWPORT_DIMS,
		MEDIUM_FLOAT:                     gl.MEDIUM_FLOAT,
		MEDIUM_INT:                       gl.MEDIUM_INT,
		MIRRORED_REPEAT:                  gl.MIRRORED_REPEAT,
		MULTISAMPLE:                      gl.MULTISAMPLE,
		NEAREST:                          gl.NEAREST,
		NEAREST_MIPMAP_LINEAR:            gl.NEAREST_MIPMAP_LINEAR,
		NEAREST_MIPMAP_NEAREST:           gl.NEAREST_MIPMAP_NEAREST,
		NEVER:                          gl.NEVER,
		NICEST:                         gl.NICEST,
		NONE:                           gl.NONE,
		NOTEQUAL:                       gl.NOTEQUAL,
		NO_ERROR:                       gl.NO_ERROR,
		NUM_COMPRESSED_TEXTURE_FORMATS: gl.NUM_COMPRESSED_TEXTURE_FORMATS,
		ONE: gl.ONE,
		ONE_MINUS_CONSTANT_ALPHA:     gl.ONE_MINUS_CONSTANT_ALPHA,
		ONE_MINUS_CONSTANT_COLOR:     gl.ONE_MINUS_CONSTANT_COLOR,
		ONE_MINUS_DST_ALPHA:          gl.ONE_MINUS_DST_ALPHA,
		ONE_MINUS_DST_COLOR:          gl.ONE_MINUS_DST_COLOR,
		ONE_MINUS_SRC_ALPHA:          gl.ONE_MINUS_SRC_ALPHA,
		ONE_MINUS_SRC_COLOR:          gl.ONE_MINUS_SRC_COLOR,
		OUT_OF_MEMORY:                gl.OUT_OF_MEMORY,
		PACK_ALIGNMENT:               gl.PACK_ALIGNMENT,
		POINTS:                       gl.POINTS,
		POLYGON_OFFSET_FACTOR:        gl.POLYGON_OFFSET_FACTOR,
		POLYGON_OFFSET_FILL:          gl.POLYGON_OFFSET_FILL,
		POLYGON_OFFSET_UNITS:         gl.POLYGON_OFFSET_UNITS,
		RENDERBUFFER:                 gl.RENDERBUFFER,
		RENDERBUFFER_ALPHA_SIZE:      gl.RENDERBUFFER_ALPHA_SIZE,
		RENDERBUFFER_BINDING:         gl.RENDERBUFFER_BINDING,
		RENDERBUFFER_BLUE_SIZE:       gl.RENDERBUFFER_BLUE_SIZE,
		RENDERBUFFER_DEPTH_SIZE:      gl.RENDERBUFFER_DEPTH_SIZE,
		RENDERBUFFER_GREEN_SIZE:      gl.RENDERBUFFER_GREEN_SIZE,
		RENDERBUFFER_HEIGHT:          gl.RENDERBUFFER_HEIGHT,
		RENDERBUFFER_INTERNAL_FORMAT: gl.RENDERBUFFER_INTERNAL_FORMAT,
		RENDERBUFFER_RED_SIZE:        gl.RENDERBUFFER_RED_SIZE,
		RENDERBUFFER_STENCIL_SIZE:    gl.RENDERBUFFER_STENCIL_SIZE,
		RENDERBUFFER_WIDTH:           gl.RENDERBUFFER_WIDTH,
		RENDERER:                     gl.RENDERER,
		REPEAT:                       gl.REPEAT,
		REPLACE:                      gl.REPLACE,
		RGB:                          gl.RGB,
		RGB5_A1:                      gl.RGB5_A1,
		RGB565:                       gl.RGB565,
		RGBA:                         gl.RGBA,
		RGBA4:                        gl.RGBA4,
		SAMPLER_2D:                   gl.SAMPLER_2D,
		SAMPLER_CUBE:                 gl.SAMPLER_CUBE,
		SAMPLES:                      gl.SAMPLES,
		SAMPLE_ALPHA_TO_COVERAGE:     gl.SAMPLE_ALPHA_TO_COVERAGE,
		SAMPLE_BUFFERS:               gl.SAMPLE_BUFFERS,
		SAMPLE_COVERAGE:              gl.SAMPLE_COVERAGE,
		SAMPLE_COVERAGE_INVERT:       gl.SAMPLE_COVERAGE_INVERT,
		SAMPLE_COVERAGE_VALUE:        gl.SAMPLE_COVERAGE_VALUE,
		SCISSOR_BOX:                  gl.SCISSOR_BOX,
		SCISSOR_TEST:                 gl.SCISSOR_TEST,
		SHADER_COMPILER:              gl.SHADER_COMPILER,
		SHADER_SOURCE_LENGTH:         gl.SHADER_SOURCE_LENGTH,
		SHADER_TYPE:                  gl.SHADER_TYPE,
		SHADING_LANGUAGE_VERSION:     gl.SHADING_LANGUAGE_VERSION,
		SHORT:                        gl.SHORT,
		SRC_ALPHA:                    gl.SRC_ALPHA,
		SRC_ALPHA_SATURATE:           gl.SRC_ALPHA_SATURATE,
		SRC_COLOR:                    gl.SRC_COLOR,
		STATIC_DRAW:                  gl.STATIC_DRAW,
		STENCIL_ATTACHMENT:           gl.STENCIL_ATTACHMENT,
		STENCIL_BACK_FAIL:            gl.STENCIL_BACK_FAIL,
		STENCIL_BACK_FUNC:            gl.STENCIL_BACK_FUNC,
		STENCIL_BACK_PASS_DEPTH_FAIL: gl.STENCIL_BACK_PASS_DEPTH_FAIL,
		STENCIL_BACK_PASS_DEPTH_PASS: gl.STENCIL_BACK_PASS_DEPTH_PASS,
		STENCIL_BACK_REF:             gl.STENCIL_BACK_REF,
		STENCIL_BACK_VALUE_MASK:      gl.STENCIL_BACK_VALUE_MASK,
		STENCIL_BACK_WRITEMASK:       gl.STENCIL_BACK_WRITEMASK,
		STENCIL_BUFFER_BIT:           gl.STENCIL_BUFFER_BIT,
		STENCIL_CLEAR_VALUE:          gl.STENCIL_CLEAR_VALUE,
		STENCIL_FAIL:                 gl.STENCIL_FAIL,
		STENCIL_FUNC:                 gl.STENCIL_FUNC,
		STENCIL_INDEX:                gl.STENCIL_INDEX,
		STENCIL_INDEX8:               gl.STENCIL_INDEX8,
		STENCIL_PASS_DEPTH_FAIL:      gl.STENCIL_PASS_DEPTH_FAIL,
		STENCIL_PASS_DEPTH_PASS:      gl.STENCIL_PASS_DEPTH_PASS,
		STENCIL_REF:                  gl.STENCIL_REF,
		STENCIL_TEST:                 gl.STENCIL_TEST,
		STENCIL_VALUE_MASK:           gl.STENCIL_VALUE_MASK,
		STENCIL_WRITEMASK:            gl.STENCIL_WRITEMASK,
		STREAM_DRAW:                  gl.STREAM_DRAW,
		SUBPIXEL_BITS:                gl.SUBPIXEL_BITS,
		TEXTURE:                      gl.TEXTURE,
		TEXTURE0:                     gl.TEXTURE0,
		TEXTURE1:                     gl.TEXTURE1,
		TEXTURE2:                     gl.TEXTURE2,
		TEXTURE3:                     gl.TEXTURE3,
		TEXTURE4:                     gl.TEXTURE4,
		TEXTURE5:                     gl.TEXTURE5,
		TEXTURE6:                     gl.TEXTURE6,
		TEXTURE7:                     gl.TEXTURE7,
		TEXTURE8:                     gl.TEXTURE8,
		TEXTURE9:                     gl.TEXTURE9,
		TEXTURE10:                    gl.TEXTURE10,
		TEXTURE11:                    gl.TEXTURE11,
		TEXTURE12:                    gl.TEXTURE12,
		TEXTURE13:                    gl.TEXTURE13,
		TEXTURE14:                    gl.TEXTURE14,
		TEXTURE15:                    gl.TEXTURE15,
		TEXTURE16:                    gl.TEXTURE16,
		TEXTURE17:                    gl.TEXTURE17,
		TEXTURE18:                    gl.TEXTURE18,
		TEXTURE19:                    gl.TEXTURE19,
		TEXTURE20:                    gl.TEXTURE20,
		TEXTURE21:                    gl.TEXTURE21,
		TEXTURE22:                    gl.TEXTURE22,
		TEXTURE23:                    gl.TEXTURE23,
		TEXTURE24:                    gl.TEXTURE24,
		TEXTURE25:                    gl.TEXTURE25,
		TEXTURE26:                    gl.TEXTURE26,
		TEXTURE27:                    gl.TEXTURE27,
		TEXTURE28:                    gl.TEXTURE28,
		TEXTURE29:                    gl.TEXTURE29,
		TEXTURE30:                    gl.TEXTURE30,
		TEXTURE31:                    gl.TEXTURE31,
		TEXTURE_2D:                   gl.TEXTURE_2D,
		TEXTURE_BINDING_2D:           gl.TEXTURE_BINDING_2D,
		TEXTURE_BINDING_CUBE_MAP:     gl.TEXTURE_BINDING_CUBE_MAP,
		TEXTURE_CUBE_MAP:             gl.TEXTURE_CUBE_MAP,
		TEXTURE_CUBE_MAP_NEGATIVE_X:  gl.TEXTURE_CUBE_MAP_NEGATIVE_X,
		TEXTURE_CUBE_MAP_NEGATIVE_Y:  gl.TEXTURE_CUBE_MAP_NEGATIVE_Y,
		TEXTURE_CUBE_MAP_NEGATIVE_Z:  gl.TEXTURE_CUBE_MAP_NEGATIVE_Z,
		TEXTURE_CUBE_MAP_POSITIVE_X:  gl.TEXTURE_CUBE_MAP_POSITIVE_X,
		TEXTURE_CUBE_MAP_POSITIVE_Y:  gl.TEXTURE_CUBE_MAP_POSITIVE_Y,
		TEXTURE_CUBE_MAP_POSITIVE_Z:  gl.TEXTURE_CUBE_MAP_POSITIVE_Z,
		TEXTURE_MAG_FILTER:           gl.TEXTURE_MAG_FILTER,
		TEXTURE_MIN_FILTER:           gl.TEXTURE_MIN_FILTER,
		TEXTURE_WRAP_S:               gl.TEXTURE_WRAP_S,
		TEXTURE_WRAP_T:               gl.TEXTURE_WRAP_T,
		TRIANGLES:                    gl.TRIANGLES,
		TRIANGLE_FAN:                 gl.TRIANGLE_FAN,
		TRIANGLE_STRIP:               gl.TRIANGLE_STRIP,
		UNPACK_ALIGNMENT:             gl.UNPACK_ALIGNMENT,
		UNSIGNED_BYTE:                gl.UNSIGNED_BYTE,
		UNSIGNED_INT:                 gl.UNSIGNED_INT,
		UNSIGNED_SHORT:               gl.UNSIGNED_SHORT,
		UNSIGNED_SHORT_4_4_4_4:       gl.UNSIGNED_SHORT_4_4_4_4,
		UNSIGNED_SHORT_5_5_5_1:       gl.UNSIGNED_SHORT_5_5_5_1,
		UNSIGNED_SHORT_5_6_5:         gl.UNSIGNED_SHORT_5_6_5,
		VALIDATE_STATUS:              gl.VALIDATE_STATUS,
		VENDOR:                       gl.VENDOR,
		VERSION:                      gl.VERSION,
		VERTEX_ATTRIB_ARRAY_BUFFER_BINDING: gl.VERTEX_ATTRIB_ARRAY_BUFFER_BINDING,
		VERTEX_ATTRIB_ARRAY_ENABLED:        gl.VERTEX_ATTRIB_ARRAY_ENABLED,
		VERTEX_ATTRIB_ARRAY_NORMALIZED:     gl.VERTEX_ATTRIB_ARRAY_NORMALIZED,
		VERTEX_ATTRIB_ARRAY_POINTER:        gl.VERTEX_ATTRIB_ARRAY_POINTER,
		VERTEX_ATTRIB_ARRAY_SIZE:           gl.VERTEX_ATTRIB_ARRAY_SIZE,
		VERTEX_ATTRIB_ARRAY_STRIDE:         gl.VERTEX_ATTRIB_ARRAY_STRIDE,
		VERTEX_ATTRIB_ARRAY_TYPE:           gl.VERTEX_ATTRIB_ARRAY_TYPE,
		VERTEX_SHADER:                      gl.VERTEX_SHADER,
		VIEWPORT:                           gl.VIEWPORT,
		ZERO:                               gl.ZERO,
		TRUE:                               gl.TRUE,
	}
}

// Gl may become engo.Gl (Gl = glplus.NewContext())
var Gl *Context

// ENGOGLTexture one day my become gl. from "engo.io/gl"
type ENGOGLTexture struct{ uint32 }

// ENGOGLBuffer ...
type ENGOGLBuffer struct{ uint32 }

// ENGOGLFrameBuffer ...
type ENGOGLFrameBuffer struct{ uint32 }

// ENGOGLRenderBuffer ...
type ENGOGLRenderBuffer struct{ uint32 }

// ENGOGLProgram ...
type ENGOGLProgram struct{ uint32 }

// ENGOGLUniformLocation ...
type ENGOGLUniformLocation struct{ int32 }

// ENGOGLShader ...
type ENGOGLShader struct{ uint32 }

// ENGOGLVertexArray ...
type ENGOGLVertexArray struct{ uint32 }

func (c *Context) DeleteProgram(program *ENGOGLProgram) {
	gl.DeleteProgram(program.uint32)
}

func (c *Context) GetProgramInfoLog(program *ENGOGLProgram) string {
	var maxLength int32
	gl.GetProgramiv(program.uint32, gl.INFO_LOG_LENGTH, &maxLength)

	errorLog := make([]byte, maxLength)
	gl.GetProgramInfoLog(program.uint32, maxLength, &maxLength, (*uint8)(gl.Ptr(errorLog)))

	return string(errorLog)
}

func (c *Context) ValidateProgram(program *ENGOGLProgram) {
	if program == nil {
		gl.ValidateProgram(0)
		return
	}
	gl.ValidateProgram(program.uint32)
}

func (c *Context) GetProgramParameterb(program *ENGOGLProgram, pname int) bool {
	var success int32 = gl.FALSE
	gl.GetProgramiv(program.uint32, uint32(pname), &success)
	return success == gl.TRUE
}

func (c *Context) GetProgramParameteri(program *ENGOGLProgram, pname int) int {
	var success int32 = gl.FALSE
	gl.GetProgramiv(program.uint32, uint32(pname), &success)
	return int(success)
}

func (c *Context) LinkProgram(program *ENGOGLProgram) {
	gl.LinkProgram(program.uint32)
}

func (c *Context) GetUniformLocation(program *ENGOGLProgram, name string) *ENGOGLUniformLocation {
	return &ENGOGLUniformLocation{gl.GetUniformLocation(program.uint32, gl.Str(name+"\x00"))}
}

func (c *Context) GetAttribLocation(program *ENGOGLProgram, name string) int {
	return int(gl.GetAttribLocation(program.uint32, gl.Str(name+"\x00")))
}

func (c *Context) UseProgram(program *ENGOGLProgram) {
	if program == nil {
		gl.UseProgram(0)
		return
	}
	gl.UseProgram(program.uint32)
}

func (c *Context) Uniform1f(location *ENGOGLUniformLocation, x float32) {
	gl.Uniform1f(location.int32, x)
}

// Assigns a integer value to a uniform variable for the current program object.
func (c *Context) Uniform1i(location *ENGOGLUniformLocation, x int) {
	gl.Uniform1i(location.int32, int32(x))
}

func (c *Context) Uniform2f(location *ENGOGLUniformLocation, x, y float32) {
	gl.Uniform2f(location.int32, x, y)
}

func (c *Context) Uniform3f(location *ENGOGLUniformLocation, x, y, z float32) {
	gl.Uniform3f(location.int32, x, y, z)
}

func (c *Context) Uniform4f(location *ENGOGLUniformLocation, x, y, z, w float32) {
	gl.Uniform4f(location.int32, x, y, z, w)
}

func (c *Context) UniformMatrix3fv(location *ENGOGLUniformLocation, transpose bool, value []float32) {
	// TODO: count value of 1 is currently hardcoded.
	//       Perhaps it should be len(value) / 16 or something else?
	//       In OpenGL 2.1 it is a manually supplied parameter, but WebGL does not have it.
	//       Not sure if WebGL automatically deduces it and supports count values greater than 1, or if 1 is always assumed.
	gl.UniformMatrix3fv(location.int32, 1, transpose, &value[0])
}

func (c *Context) UniformMatrix4fv(location *ENGOGLUniformLocation, transpose bool, value []float32) {
	// TODO: count value of 1 is currently hardcoded.
	//       Perhaps it should be len(value) / 16 or something else?
	//       In OpenGL 2.1 it is a manually supplied parameter, but WebGL does not have it.
	//       Not sure if WebGL automatically deduces it and supports count values greater than 1, or if 1 is always assumed.
	gl.UniformMatrix4fv(location.int32, 1, transpose, &value[0])
}

func (c *Context) CreateProgram() *ENGOGLProgram {
	return &ENGOGLProgram{gl.CreateProgram()}
}

func (c *Context) AttachShader(program *ENGOGLProgram, shader *ENGOGLShader) {
	gl.AttachShader(program.uint32, shader.uint32)
}

func (c *Context) CreateShader(typ int) *ENGOGLShader {
	shader := &ENGOGLShader{gl.CreateShader(uint32(typ))}
	return shader
}

func (c *Context) ShaderSource(shader *ENGOGLShader, source string) {
	glsource, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader.uint32, 1, glsource, nil)
	free()
}

func (c *Context) CompileShader(shader *ENGOGLShader) {
	gl.CompileShader(shader.uint32)
}

func (c *Context) GetShaderiv(shader *ENGOGLShader, pname uint32) bool {
	var success int32
	gl.GetShaderiv(shader.uint32, pname, &success)
	return success == int32(gl.TRUE)
}

func (c *Context) GetShaderInfoLog(shader *ENGOGLShader) string {
	var maxLength int32
	gl.GetShaderiv(shader.uint32, gl.INFO_LOG_LENGTH, &maxLength)

	errorLog := make([]byte, maxLength)
	gl.GetShaderInfoLog(shader.uint32, maxLength, &maxLength, (*uint8)(gl.Ptr(errorLog)))

	return string(errorLog)
}

func (c *Context) BindAttribLocation(program *ENGOGLProgram, index int, name string) {
	gl.BindAttribLocation(program.uint32, uint32(index), gl.Str(name+"\x00"))
}

func (c *Context) DeleteShader(shader *ENGOGLShader) {
	gl.DeleteShader(shader.uint32)
}

func (c *Context) DeleteTexture(texture *ENGOGLTexture) {
	gl.DeleteTextures(1, &[]uint32{texture.uint32}[0])
}

func (c *Context) DeleteBuffer(buffer *ENGOGLBuffer) {
	gl.DeleteBuffers(1, &[]uint32{buffer.uint32}[0])
}
