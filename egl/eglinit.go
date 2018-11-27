// Copyright 2018 The Gooid Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package egl

import (
	"log"
	"unsafe"
)

type NativeObj interface {
	NativeDisplay() unsafe.Pointer
	NativeWindow() unsafe.Pointer
	SetBuffersGeometry(int) int
	WindowSize() (w, h int)
}

func CreateEGLContext(native NativeObj) *EGLContext {
	return CreateEGLContextEx(native, 16, 2)
}

func CreateEGLContextEx(native NativeObj, depthSize, esVersion int) *EGLContext {
	eglctx := NewContextEx(NativeWindow(native.NativeWindow()),
		NativeDisplay(native.NativeDisplay()), depthSize, esVersion, 0)
	log.Println("EGL InitEGLSurface...")
	if !eglctx.InitEGLSurface() {
		log.Println("Init EGL Surface failed.", GetError())
		return nil
	}

	format := eglctx.GetFormat()
	if native.SetBuffersGeometry(format) != 0 {
		log.Println("EGL set buffers geometry failed.", GetError())
		eglctx.Terminate()
		return nil
	}

	log.Println("EGL InitEGLContext...")
	if !eglctx.InitEGLContext() {
		log.Println("Init EGL Context failed.", GetError())
		eglctx.Terminate()
		return nil
	}

	return eglctx
}
