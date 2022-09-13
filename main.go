package main

/*
#cgo CFLAGS: -I/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.312.b07-1.el7_9.x86_64/include
#cgo CFLAGS: -I/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.312.b07-1.el7_9.x86_64/include/linux

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

//  linux下的地址 /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.312.b07-1.el7_9.x86_64/include
// mac下的地址  /Library/Java/JavaVirtualMachines/jdk-18.0.2.1.jdk/Contents/Home/include

//export Java_Hello_add
func Java_Hello_add(env *C.JNIEnv, clazz C.jclass, x C.jlong, y C.jlong) C.jstring {
	fmt.Println("Hello from Go")

	s := fmt.Sprintf("res:%d + %d", int64(x), int64(y))

	r := C.CString(s)

	defer C.free(unsafe.Pointer(r))

	return C.createJString(env, r)
}

func main() {}
