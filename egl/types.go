// Copyright 2014 The Gooid Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//http://mortdeus.mit-license.org

package egl

/*
#include <EGL/egl.h>
*/
import "C"
import "unsafe"

type (
	Config        unsafe.Pointer
	Context       unsafe.Pointer
	Display       unsafe.Pointer
	Surface       unsafe.Pointer
	ClientBuffer  unsafe.Pointer
	NativeDisplay unsafe.Pointer
	NativeWindow  unsafe.Pointer
	NativePixmap  unsafe.Pointer
	EGLint        C.EGLint
)

var (
	DEFAULT_DISPLAY NativeDisplay
	NO_CONTEXT      Context
	NO_DISPLAY      Display
	NO_SURFACE      Surface
)

var Version struct{ Maj, Min EGLint }

func goBool(n C.EGLBoolean) bool {
	return n == 1
}
func eglBool(n bool) C.EGLBoolean {
	var b int
	if n == true {
		b = 1
	}
	return C.EGLBoolean(b)
}

/*
func ProcAdress(proc string) uintptr {

}
*/
