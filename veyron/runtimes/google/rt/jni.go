// +build android

package rt

import (
	jutil "v.io/jni/util"
	jipc "v.io/jni/veyron/runtimes/google/ipc"
	jnaming "v.io/jni/veyron/runtimes/google/naming"
	jcontext "v.io/jni/veyron2/context"
	jsecurity "v.io/jni/veyron2/security"

	_ "v.io/core/veyron/profiles/roaming"
	"v.io/core/veyron2"
)

// #cgo LDFLAGS: -ljniwrapper
// #include "jni_wrapper.h"
import "C"

// Init initializes the JNI code with the given Java environment.  This method
// must be invoked before any other method in this package and must be called
// from the main Java thread (e.g., On_Load()).
// NOTE: Because CGO creates package-local types and because this method may be
// invoked from a different package, Java environment is passed in an empty
// interface and then cast into the package-local environment type.
func Init(jEnv interface{}) {}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeInit
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeInit(env *C.JNIEnv, jRuntime C.jclass) C.jobject {
	ctx, _ := veyron2.Init()
	// Get the original spec, which is guaranteed to be a roaming spec (as we
	// import the roaming profile).
	roamingSpec := veyron2.GetListenSpec(ctx)
	jipc.SetRoamingSpec(roamingSpec)
	jCtx, err := jcontext.JavaContext(env, ctx, nil)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jCtx)
}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeSetNewClient
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeSetNewClient(env *C.JNIEnv, jRuntime C.jclass, jContext C.jobject, jOptions C.jobject) C.jobject {
	// TODO(spetrovic): Have Java context support nativePtr()?
	ctx, err := jcontext.GoContext(env, jContext)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	// No options supported yet.
	newCtx, _, err := veyron2.SetNewClient(ctx)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	jNewCtx, err := jcontext.JavaContext(env, newCtx, nil)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jNewCtx)
}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeGetClient
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeGetClient(env *C.JNIEnv, jRuntime C.jclass, jContext C.jobject) C.jobject {
	// TODO(spetrovic): Have Java context support nativePtr()?
	ctx, err := jcontext.GoContext(env, jContext)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	client := veyron2.GetClient(ctx)
	jClient, err := jipc.JavaClient(env, client)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jClient)
}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeNewServer
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeNewServer(env *C.JNIEnv, jRuntime C.jclass, jContext C.jobject, jOptions C.jobject) C.jobject {
	// TODO(spetrovic): Have Java context support nativePtr()?
	ctx, err := jcontext.GoContext(env, jContext)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	server, err := veyron2.NewServer(ctx)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	jServer, err := jipc.JavaServer(env, server, jOptions)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jServer)
}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeSetPrincipal
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeSetPrincipal(env *C.JNIEnv, jRuntime C.jclass, jContext C.jobject, jPrincipal C.jobject) C.jobject {
	// TODO(spetrovic): Have Java context support nativePtr()?
	ctx, err := jcontext.GoContext(env, jContext)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	principal, err := jsecurity.GoPrincipal(env, jPrincipal)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	newCtx, err := veyron2.SetPrincipal(ctx, principal)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	jNewCtx, err := jcontext.JavaContext(env, newCtx, nil)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jNewCtx)
}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeGetPrincipal
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeGetPrincipal(env *C.JNIEnv, jRuntime C.jclass, jContext C.jobject) C.jobject {
	// TODO(spetrovic): Have Java context support nativePtr()?
	ctx, err := jcontext.GoContext(env, jContext)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	principal := veyron2.GetPrincipal(ctx)
	jPrincipal, err := jsecurity.JavaPrincipal(env, principal)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jPrincipal)
}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeSetNamespace
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeSetNamespace(env *C.JNIEnv, jRuntime C.jclass, jContext C.jobject, jRoots C.jobjectArray) C.jobject {
	// TODO(spetrovic): Have Java context support nativePtr()?
	ctx, err := jcontext.GoContext(env, jContext)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	roots := jutil.GoStringArray(env, jRoots)
	newCtx, _, err := veyron2.SetNewNamespace(ctx, roots...)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	jNewCtx, err := jcontext.JavaContext(env, newCtx, nil)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jNewCtx)
}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeGetNamespace
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeGetNamespace(env *C.JNIEnv, jRuntime C.jclass, jContext C.jobject) C.jobject {
	// TODO(spetrovic): Have Java context support nativePtr()?
	ctx, err := jcontext.GoContext(env, jContext)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	namespace := veyron2.GetNamespace(ctx)
	jNamespace, err := jnaming.JavaNamespace(env, namespace)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jNamespace)
}

//export Java_io_v_core_veyron_runtimes_google_VRuntime_nativeGetListenSpec
func Java_io_v_core_veyron_runtimes_google_VRuntime_nativeGetListenSpec(env *C.JNIEnv, jRuntime C.jclass, jContext C.jobject) C.jobject {
	ctx, err := jcontext.GoContext(env, jContext)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	spec := veyron2.GetListenSpec(ctx)
	jSpec, err := jipc.JavaListenSpec(env, spec)
	if err != nil {
		jutil.JThrowV(env, err)
		return nil
	}
	return C.jobject(jSpec)
}
