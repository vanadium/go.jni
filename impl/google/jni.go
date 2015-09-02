// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build java android

package google

import (
	jchannel "v.io/x/jni/impl/google/channel"
	jns "v.io/x/jni/impl/google/namespace"
	jrpc "v.io/x/jni/impl/google/rpc"
	jrt "v.io/x/jni/impl/google/rt"
	jsyncbased "v.io/x/jni/impl/google/services/syncbase/syncbased"
	jutil "v.io/x/jni/util"
)

// #include "jni.h"
import "C"

// Init initializes the JNI code with the given Java environment.  This method
// must be invoked before any other method in this package and must be called
// from the main Java thread (e.g., On_Load()).
// interface and then cast into the package-local environment type.
func Init(env jutil.Env) error {
	if err := jrpc.Init(env); err != nil {
		return err
	}
	if err := jrt.Init(env); err != nil {
		return err
	}
	if err := jchannel.Init(env); err != nil {
		return err
	}
	if err := jns.Init(env); err != nil {
		return err
	}
	if err := jsyncbased.Init(env); err != nil {
		return err
	}
	return nil
}
