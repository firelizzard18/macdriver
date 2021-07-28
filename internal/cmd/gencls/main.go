package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#define OBJC2_UNAVAILABLE
#include <objc/objc.h>
#include <objc/runtime.h>
#include <Foundation/Foundation.h>

bool isaNSObject(Class cls) {
	if (!class_getClassMethod(cls, @selector(isSubclassOfClass:))) {
		return NO;
	}
	return [cls isSubclassOfClass:[NSObject class]];
}
*/
import "C"

const nameNSObject = "NSObject\x00"

func main() {
	n := C.objc_getClassList(nil, 0)
	classes := make([]C.Class, n)
	C.objc_getClassList(&classes[0], n)

	for _, class := range classes[:100] {
		name := C.GoString(C.class_getName(class))
		if !C.isaNSObject(class) {
			continue
		}

		println(name)
	}
}
