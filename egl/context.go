// Copyright 2018 The Gooid Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package egl

import (
	"log"
	"unsafe"
)

type Window unsafe.Pointer
type EGLContext struct {
	surface       Surface
	display       Display
	context       Context
	config        Config
	window        NativeWindow
	esMajor       EGLint
	esMinor       EGLint
	depthSize     EGLint
	bAntialiasing bool
}

/* depthSize : 16, 24
 * esVersion : 2 (es2), 3 (es3)
 */
func NewContext(window NativeWindow, ndisplay NativeDisplay) *EGLContext {
	return NewContextEx(window, ndisplay, 16, 2, 0)
}

func NewContextEx(window NativeWindow, ndisplay NativeDisplay, depthSize, esVersion, minor int) *EGLContext {
	if !(depthSize == 32 || depthSize == 24) {
		depthSize = 16
	}
	if esVersion != 3 {
		esVersion = 2
	}
	//display := Display(nil)
	display := GetDisplay(ndisplay)
	if display == nil {
		log.Println("EGL GetDisplay failed")
	}
	return &EGLContext{window: window,
		display:       display,
		surface:       NO_SURFACE,
		depthSize:     EGLint(depthSize),
		esMajor:       EGLint(esVersion),
		esMinor:       EGLint(minor),
		bAntialiasing: true}
}

func (ctx *EGLContext) GetFormat() int {
	format, _ := GetConfigAttrib(ctx.display, ctx.config, NATIVE_VISUAL_ID)
	return format
}

func (ctx *EGLContext) InitEGLSurfaceX() bool {
	if !Initialize(ctx.display) {
		log.Println("EGL initialize failed")
		return false
	}

	/*
	 * Here specify the attributes of the desired configuration.
	 * Below, we select an EGLConfig with at least 8 bits per color
	 * component compatible with on-screen windows
	 */
	attribs := []EGLint{RENDERABLE_TYPE, OPENGL_ES2_BIT, //Request opengl ES2.0
		SURFACE_TYPE, WINDOW_BIT,
		BLUE_SIZE, 5, GREEN_SIZE, 6,
		RED_SIZE, 5, BUFFER_SIZE, 16,
		SAMPLE_BUFFERS, 1, SAMPLES, 2, // Antialiasing
		RENDERABLE_TYPE, 4, DEPTH_SIZE, 16,
		NONE}

	var confs [1]Config
	if ctx.esMajor == 3 {
		attribs[1] = OPENGL_ES3_BIT
	}
	attribs[11] = ctx.depthSize
	num_configs := ChooseConfig(ctx.display, attribs, confs[:])

	if 0 == num_configs && ctx.depthSize > 16 {
		//Fall back to 16bit depth buffer
		attribs[11] = 16
		num_configs = ChooseConfig(ctx.display, attribs, confs[:])
	}

	if 0 == num_configs {
		return false
	}
	ctx.config = confs[0]
	ctx.surface = CreateWindowSurface(ctx.display, ctx.config, ctx.window, nil)
	if ctx.surface == nil {
		return false
	}

	//screen_width, _ := QuerySurface(ctx.display, ctx.surface, WIDTH)
	//screen_height, _ := QuerySurface(ctx.display, ctx.surface, HEIGHT)

	/* EGL_NATIVE_VISUAL_ID is an attribute of the EGLConfig that is
	 * guaranteed to be accepted by ANativeWindow_setBuffersGeometry().
	 * As soon as we picked a EGLConfig, we can safely reconfigure the
	 * ANativeWindow buffers to match, using EGL_NATIVE_VISUAL_ID. */
	//EGLint format;
	//eglGetConfigAttrib( display_, config_, EGL_NATIVE_VISUAL_ID, &format );
	//ANativeWindow_setBuffersGeometry( window_, 0, 0, format );

	return true
}

func (ctx *EGLContext) InitEGLSurface() bool {
	if !Initialize(ctx.display) {
		log.Println("EGL initialize failed")
		return false
	}

	var depths []int
	if ctx.depthSize > 16 {
		depths = []int{32, 16}
	} else {
		depths = []int{16, 32}
	}

	for _, depth := range depths {
		var attribs []EGLint
		if depth > 16 {
			attribs = []EGLint{RENDERABLE_TYPE, OPENGL_ES2_BIT, //Request opengl ES2.0
				BLUE_SIZE, 8, GREEN_SIZE, 8, RED_SIZE, 8,
				NONE}
		} else {
			attribs = []EGLint{RENDERABLE_TYPE, OPENGL_ES2_BIT, //Request opengl ES2.0
				BLUE_SIZE, 5, GREEN_SIZE, 6, RED_SIZE, 5,
				NONE}
		}

		var confs [16]Config
		if ctx.esMajor == 3 {
			attribs[1] = OPENGL_ES3_BIT
		}
		num_configs := ChooseConfig(ctx.display, attribs, confs[:])
		if 0 == num_configs {
			continue
		}

		// find match config
		for _, conf := range confs[:num_configs] {
			r, _ := GetConfigAttrib(ctx.display, conf, RED_SIZE)
			g, _ := GetConfigAttrib(ctx.display, conf, GREEN_SIZE)
			b, _ := GetConfigAttrib(ctx.display, conf, BLUE_SIZE)
			s, _ := GetConfigAttrib(ctx.display, conf, SAMPLE_BUFFERS)
			if (ctx.depthSize > 16 && r == 8 && g == 8 && b == 8) ||
				(ctx.depthSize == 16 && r == 5 && g == 6 && b == 5) {
				ctx.config = conf
				if ctx.bAntialiasing {
					if s > 0 {
						break
					}
				} else {
					break
				}
			}
		}

		if ctx.config == nil {
			ctx.config = confs[0]
		}
	}

	if ctx.config == nil {
		return false
	}

	ctx.surface = CreateWindowSurface(ctx.display, ctx.config, ctx.window, nil)
	if ctx.surface == nil {
		return false
	}

	return true
}

// auto match
func (ctx *EGLContext) InitEGLSurface_XX() bool {
	if !Initialize(ctx.display) {
		log.Println("EGL initialize failed")
		return false
	}

	var depths []int
	if ctx.depthSize > 16 {
		depths = []int{32, 16}
	} else {
		depths = []int{16, 32}
	}

	for _, depth := range depths {
		var attribs []EGLint
		if depth > 16 {
			attribs = []EGLint{RENDERABLE_TYPE, OPENGL_ES2_BIT, //Request opengl ES2.0
				BLUE_SIZE, 8, GREEN_SIZE, 8, RED_SIZE, 8,
				NONE}
		} else {
			attribs = []EGLint{RENDERABLE_TYPE, OPENGL_ES2_BIT, //Request opengl ES2.0
				BLUE_SIZE, 5, GREEN_SIZE, 6, RED_SIZE, 5,
				NONE}
		}

		var confs [16]Config
		if ctx.esMajor == 3 {
			attribs[1] = OPENGL_ES3_BIT
		}
		num_configs := ChooseConfig(ctx.display, attribs, confs[:])
		if 0 == num_configs {
			continue
		}

		// find match config
		for _, conf := range confs[:num_configs] {
			ctx.surface = CreateWindowSurface(ctx.display, conf, ctx.window, nil)
			if ctx.surface != nil {
				ctx.config = conf
				return true
			}
		}
	}
	return false
}

func (ctx *EGLContext) InitEGLContext() bool {
	context_attribs := []EGLint{
		CONTEXT_CLIENT_VERSION, 2, //Request opengl ES2.0
		CONTEXT_MAJOR_VERSION, 2,
		CONTEXT_MINOR_VERSION, 0,
		NONE, NONE}
	if ctx.esMajor >= 3 {
		context_attribs[1] = ctx.esMajor
		context_attribs[3] = ctx.esMajor
		context_attribs[5] = ctx.esMinor
	} else {
		context_attribs[2] = NONE
	}
	ctx.context = CreateContext(ctx.display, ctx.config, nil, context_attribs)

	if !MakeCurrent(ctx.display, ctx.surface, ctx.surface, ctx.context) {
		log.Println("Unable to eglMakeCurrent")
		return false
	}
	return true
}

func (ctx *EGLContext) ReinitEGLContext() bool {
	if ctx.display != NO_DISPLAY {
		MakeCurrent(ctx.display, NO_SURFACE, NO_SURFACE, NO_CONTEXT)
		if ctx.context != NO_CONTEXT {
			DestroyContext(ctx.display, ctx.context)
		}
	}
	return ctx.InitEGLContext()
}

func (ctx *EGLContext) SwapBuffers() bool {
	b := SwapBuffers(ctx.display, ctx.surface)
	if !b {
		err := GetError()
		if err == BAD_SURFACE {
			//Recreate surface
			ctx.InitEGLSurface()
			return true //Still consider glContext is valid
		} else if err == CONTEXT_LOST || err == BAD_CONTEXT {
			//Context has been lost!!
			ctx.Terminate()
			ctx.InitEGLContext()
		}
		return false
	}
	return true
}

func (ctx *EGLContext) Terminate() {
	if ctx.display != NO_DISPLAY {
		MakeCurrent(ctx.display, NO_SURFACE, NO_SURFACE, NO_CONTEXT)
		if ctx.context != NO_CONTEXT {
			DestroyContext(ctx.display, ctx.context)
		}

		if ctx.surface != NO_SURFACE {
			DestroySurface(ctx.display, ctx.surface)
		}
		Terminate(ctx.display)
	}

	ctx.display = NO_DISPLAY
	ctx.context = NO_CONTEXT
	ctx.surface = NO_SURFACE
}

func (ctx *EGLContext) Resume() bool {
	//Create surface
	log.Println("EGLContext.Resume...")
	ctx.surface = CreateWindowSurface(ctx.display, ctx.config, ctx.window, nil)
	//screen_width := QuerySurface(display_, surface_, WIDTH, &screen_width_)
	//screen_height := QuerySurface(display_, surface_, HEIGHT, &screen_height_)

	if MakeCurrent(ctx.display, ctx.surface, ctx.surface, ctx.context) {
		return true
	}

	err := GetError()
	log.Printf("Unable to eglMakeCurrent %d\n", err)

	if err == CONTEXT_LOST {
		//Recreate context
		log.Println("Re-creating egl context")
		ctx.InitEGLContext()
	} else {
		//Recreate surface
		ctx.Terminate()
		ctx.InitEGLSurface()
		ctx.InitEGLContext()
	}
	return true
}

func (ctx *EGLContext) Suspend() {
	if ctx.surface != NO_SURFACE {
		log.Println("EGLContext.Suspend...")
		DestroySurface(ctx.display, ctx.surface)
		ctx.surface = NO_SURFACE
	}
}

func (ctx *EGLContext) IsReady() bool {
	return NO_SURFACE != ctx.surface
}
