package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#define OBJC2_UNAVAILABLE
#include <objc/objc.h>
#include <objc/runtime.h>
*/
import "C"

func main() {
	n := C.objc_getClassList(nil, 0)
	classes := make([]C.Class, n)
	C.objc_getClassList(&classes[0], n)

	for _, class := range classes[:100] {
		println(C.GoString(C.class_getName(class)))
	}
}
