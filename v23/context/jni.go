// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build java android

package context

import (
	"errors"
	"unsafe"

	"v.io/v23/context"
	jutil "v.io/x/jni/util"
)

// #include "jni.h"
// #include <stdlib.h>
import "C"

var (
	classSign = jutil.ClassSign("java.lang.Class")
	// Global reference for io.v.v23.context.VContext class.
	jVContextClass jutil.Class
)

// Init initializes the JNI code with the given Java environment. This method
// must be called from the main Java thread.
func Init(env jutil.Env) error {
	// Cache global references to all Java classes used by the package.  This is
	// necessary because JNI gets access to the class loader only in the system
	// thread, so we aren't able to invoke FindClass in other threads.
	var err error
	jVContextClass, err = jutil.JFindClass(env, "io/v/v23/context/VContext")
	if err != nil {
		return err
	}
	return nil
}

//export Java_io_v_v23_context_VContext_nativeCreate
func Java_io_v_v23_context_VContext_nativeCreate(jenv *C.JNIEnv, jVContext C.jclass) C.jobject {
	env := jutil.Env(uintptr(unsafe.Pointer(jenv)))
	ctx, _ := context.RootContext()
	jContext, err := JavaContext(env, ctx, nil)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(unsafe.Pointer(jContext))
}

//export Java_io_v_v23_context_VContext_nativeCancel
func Java_io_v_v23_context_VContext_nativeCancel(jenv *C.JNIEnv, jVContext C.jobject, goCancelPtr C.jlong) {
	(*(*context.CancelFunc)(jutil.NativePtr(goCancelPtr)))()
}

//export Java_io_v_v23_context_VContext_nativeIsCanceled
func Java_io_v_v23_context_VContext_nativeIsCanceled(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong) C.jboolean {
	if (*(*context.T)(jutil.NativePtr(goPtr))).Err() == nil {
		return C.JNI_FALSE
	}
	return C.JNI_TRUE
}

//export Java_io_v_v23_context_VContext_nativeDeadline
func Java_io_v_v23_context_VContext_nativeDeadline(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong) C.jobject {
	env := jutil.Env(uintptr(unsafe.Pointer(jenv)))
	d, ok := (*(*context.T)(jutil.NativePtr(goPtr))).Deadline()
	if !ok {
		return nil
	}
	jDeadline, err := jutil.JTime(env, d)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(unsafe.Pointer(jDeadline))
}

//export Java_io_v_v23_context_VContext_nativeDone
func Java_io_v_v23_context_VContext_nativeDone(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong, jCallbackObj C.jobject) {
	env := jutil.Env(uintptr(unsafe.Pointer(jenv)))
	jCallback := jutil.Object(uintptr(unsafe.Pointer(jCallbackObj)))
	c := (*(*context.T)(jutil.NativePtr(goPtr))).Done()
	if c == nil {
		jutil.CallbackOnFailure(env, jCallback, errors.New("Context isn't cancelable"))
		return
	}
	jutil.DoAsyncCall(env, jCallback, func() (jutil.Object, error) {
		<-c
		return jutil.NullObject, nil
	})
}

//export Java_io_v_v23_context_VContext_nativeValue
func Java_io_v_v23_context_VContext_nativeValue(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong, jKey C.jobject) C.jobject {
	env := jutil.Env(uintptr(unsafe.Pointer(jenv)))
	key, err := GoContextKey(env, jutil.Object(uintptr(unsafe.Pointer(jKey))))
	value := (*(*context.T)(jutil.NativePtr(goPtr))).Value(key)
	jValue, err := JavaContextValue(env, value)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(unsafe.Pointer(jValue))
}

//export Java_io_v_v23_context_VContext_nativeWithCancel
func Java_io_v_v23_context_VContext_nativeWithCancel(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong) C.jobject {
	env := jutil.Env(uintptr(unsafe.Pointer(jenv)))
	ctx, cancelFunc := context.WithCancel((*context.T)(jutil.NativePtr(goPtr)))
	jCtx, err := JavaContext(env, ctx, cancelFunc)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(unsafe.Pointer(jCtx))
}

//export Java_io_v_v23_context_VContext_nativeWithDeadline
func Java_io_v_v23_context_VContext_nativeWithDeadline(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong, jDeadline C.jobject) C.jobject {
	env := jutil.Env(uintptr(unsafe.Pointer(jenv)))
	deadline, err := jutil.GoTime(env, jutil.Object(uintptr(unsafe.Pointer(jDeadline))))
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	ctx, cancelFunc := context.WithDeadline((*context.T)(jutil.NativePtr(goPtr)), deadline)
	jCtx, err := JavaContext(env, ctx, cancelFunc)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(unsafe.Pointer(jCtx))
}

//export Java_io_v_v23_context_VContext_nativeWithTimeout
func Java_io_v_v23_context_VContext_nativeWithTimeout(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong, jTimeout C.jobject) C.jobject {
	env := jutil.Env(uintptr(unsafe.Pointer(jenv)))
	timeout, err := jutil.GoDuration(env, jutil.Object(uintptr(unsafe.Pointer(jTimeout))))
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	ctx, cancelFunc := context.WithTimeout((*context.T)(jutil.NativePtr(goPtr)), timeout)
	jCtx, err := JavaContext(env, ctx, cancelFunc)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(unsafe.Pointer(jCtx))
}

//export Java_io_v_v23_context_VContext_nativeWithValue
func Java_io_v_v23_context_VContext_nativeWithValue(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong, goCancelPtr C.jlong, jKey C.jobject, jValue C.jobject) C.jobject {
	env := jutil.Env(uintptr(unsafe.Pointer(jenv)))
	key, err := GoContextKey(env, jutil.Object(uintptr(unsafe.Pointer(jKey))))
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	value, err := GoContextValue(env, jutil.Object(uintptr(unsafe.Pointer(jValue))))
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	ctx := context.WithValue((*context.T)(jutil.NativePtr(goPtr)), key, value)
	var cancel context.CancelFunc
	if goCancelPtr != 0 {
		cancel = (*(*context.CancelFunc)(jutil.NativePtr(goCancelPtr)))
	}
	jCtx, err := JavaContext(env, ctx, cancel)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(unsafe.Pointer(jCtx))
}

//export Java_io_v_v23_context_VContext_nativeFinalize
func Java_io_v_v23_context_VContext_nativeFinalize(jenv *C.JNIEnv, jVContext C.jobject, goPtr C.jlong, goCancelPtr C.jlong) {
	jutil.GoUnref(jutil.NativePtr(goPtr))
	if goCancelPtr != 0 {
		jutil.GoUnref(jutil.NativePtr(goCancelPtr))
	}

}
