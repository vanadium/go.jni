// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build java android

package channel

import (
	"fmt"

	jutil "v.io/x/jni/util"
)

// #include "jni.h"
import "C"

var (
	// Global reference for io.v.impl.google.channel.InputChannel class.
	jInputChannelImplClass C.jclass
	// Global reference for io.v.impl.google.channel.OutputChannel class.
	jOutputChannelImplClass C.jclass
	// Global reference for java.io.EOFException class.
	jEOFExceptionClass C.jclass
)

// Init initializes the JNI code with the given Java environment.  This method
// must be invoked before any other method in this package and must be called
// from the main Java thread (e.g., On_Load()).
// NOTE: Because CGO creates package-local types and because this method may be
// invoked from a different package, Java environment is passed in an empty
// interface and then cast into the package-local environment type.
func Init(jEnv interface{}) error {
	class, err := jutil.JFindClass(jEnv, "io/v/impl/google/channel/InputChannel")
	if err != nil {
		return err
	}
	jInputChannelImplClass = C.jclass(class)
	class, err = jutil.JFindClass(jEnv, "io/v/impl/google/channel/OutputChannel")
	if err != nil {
		return err
	}
	jOutputChannelImplClass = C.jclass(class)
	class, err = jutil.JFindClass(jEnv, "java/io/EOFException")
	if err != nil {
		return err
	}
	jEOFExceptionClass = C.jclass(class)
	return nil
}

//export Java_io_v_impl_google_channel_InputChannel_nativeAvailable
func Java_io_v_impl_google_channel_InputChannel_nativeAvailable(env *C.JNIEnv, jInputChannel C.jobject, goChanPtr C.jlong) C.jboolean {
	ch := *(*chan C.jobject)(jutil.Ptr(goChanPtr))
	if len(ch) > 0 {
		return C.JNI_TRUE
	}
	return C.JNI_FALSE
}

//export Java_io_v_impl_google_channel_InputChannel_nativeReadValue
func Java_io_v_impl_google_channel_InputChannel_nativeReadValue(env *C.JNIEnv, jInputChannel C.jobject, goChanPtr C.jlong) C.jobject {
	ch := *(*chan C.jobject)(jutil.Ptr(goChanPtr))
	jObj, ok := <-ch
	if !ok {
		jutil.JThrow(env, jEOFExceptionClass, "Channel closed.")
		return nil
	}
	jObjLocal := C.jobject(jutil.NewLocalRef(env, jObj))
	jutil.DeleteGlobalRef(env, jObj)
	return jObjLocal
}

//export Java_io_v_impl_google_channel_InputChannel_nativeFinalize
func Java_io_v_impl_google_channel_InputChannel_nativeFinalize(env *C.JNIEnv, jInputChannel C.jobject, goChanPtr C.jlong, goSourceChanPtr C.jlong) {
	jutil.GoUnref(jutil.Ptr(goChanPtr))
	jutil.GoUnref(jutil.Ptr(goSourceChanPtr))
}

//export Java_io_v_impl_google_channel_OutputChannel_nativeWriteValue
func Java_io_v_impl_google_channel_OutputChannel_nativeWriteValue(env *C.JNIEnv, jOutputChannelClass C.jclass, goChanPtr C.jlong, jObject C.jobject) {
	outCh := *(*outputChannel)(jutil.Ptr(goChanPtr))
	jGlobalObject := C.jobject(jutil.NewGlobalRef(env, jObject))
	if err := outCh.ReadFunc(jGlobalObject); err != nil {
		jutil.JThrowV(env, fmt.Errorf("Exception while writing to OutputChannel: %+v", err))
	}
}

//export Java_io_v_impl_google_channel_OutputChannel_nativeClose
func Java_io_v_impl_google_channel_OutputChannel_nativeClose(env *C.JNIEnv, jOutputChannelClass C.jclass, goChanPtr C.jlong) {
	outCh := *(*outputChannel)(jutil.Ptr(goChanPtr))

	if err := outCh.CloseFunc(); err != nil {
		jutil.JThrowV(env, fmt.Errorf("Exception while closing OutputChannel: %+v", err))
	}
}

//export Java_io_v_impl_google_channel_OutputChannel_nativeFinalize
func Java_io_v_impl_google_channel_OutputChannel_nativeFinalize(env *C.JNIEnv, jOutputChannelClass C.jclass, goChanPtr C.jlong) {
	jutil.GoUnref(jutil.Ptr(goChanPtr))
}
