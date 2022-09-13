package main

/*
#cgo CFLAGS: -I/Library/Java/JavaVirtualMachines/jdk-18.0.2.1.jdk/Contents/Home/include
#cgo CFLAGS: -I/Library/Java/JavaVirtualMachines/jdk-18.0.2.1.jdk/Contents/Home/include/darwin

#include <stdlib.h>
#include <stdio.h>
#include <jni.h>
#include "jwrapper.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export Java_Hello_add
func Java_Hello_add(env *C.JNIEnv, clazz C.jclass, x C.jlong, y C.jlong) C.jstring {
	fmt.Println("Hello from Go")

	s := fmt.Sprintf("res:%d + %d", int64(x), int64(y))

	r := C.CString(s)

	defer C.free(unsafe.Pointer(r))

	return C.createJString(env, r)
}

func main() {}
