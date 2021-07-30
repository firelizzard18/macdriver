// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

/*
//#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework Foundation
#define OBJC_OLD_DISPATCH_PROTOTYPES 1

// #cgo LDFLAGS: -lobjc -framework Foundation
#define __OBJC2__ 1
#include <objc/runtime.h>
#include <objc/message.h>
#include <stdlib.h>

void *GoObjc_GetObjectSuperClassStruct(void *obj) {
	struct objc_super *s = malloc(sizeof(struct objc_super));
	s->receiver = obj;
	s->super_class = class_getSuperclass(object_getClass(obj));
	return s;
}

void GoObjc_MsgSend_Stret0(void *stretAddr, void *self, void *op) {
	objc_msgSend_stret(stretAddr, self, op);
}

void GoObjc_MsgSend_Stret1(void *stretAddr, void *self, void *op, void *arg) {
	objc_msgSend_stret(stretAddr, self, op, arg);
}

*/
import "C"
import (
	"math"
	"reflect"

	"github.com/progrium/macdriver/misc/variadic"
)

func unpackStruct(val reflect.Value) []uintptr {
	memArgs := []uintptr{}
	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		kind := v.Kind()
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			memArgs = append(memArgs, uintptr(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			memArgs = append(memArgs, uintptr(v.Uint()))
		case reflect.Float32, reflect.Float64:
			memArgs = append(memArgs, uintptr(math.Float64bits(v.Float())))
		case reflect.Ptr:
			memArgs = append(memArgs, val.Pointer())
		case reflect.Struct:
			args := unpackStruct(v)
			memArgs = append(memArgs, args...)
		}
	}
	return memArgs
}

func sendMsg(obj Object, sendFunc variadic.Function, selector string, args ...interface{}) Object {
	// Keep ObjC semantics: messages can be sent to nil objects,
	// but the response is nil.
	if obj.Pointer() == 0 {
		return obj
	}

	sel := selectorWithName(selector)
	if sel == nil {
		return nil
	}

	typeInfo := simpleTypeInfoForMethod(obj, selector)
	if len(typeInfo) == 0 {
		panic("invalid type encoding")
	}

	fn := buildDispatcher(typeInfo)

	var stret interface{}
	if typeInfo[0] == encStructBegin {
		sendFunc = sendFunc.StRet()
		if len(args) > 0 {
			stret = args[len(args)-1]
		}

		if stret == nil || reflect.ValueOf(stret).Kind() != reflect.Ptr {
			panic("method returns a struct but pointer argument is missing")
		}
	}

	ptr := fn(sendFunc, obj, sel, args, stret)
	return ObjectPtr(ptr)
}

func (obj object) Send(selector string, args ...interface{}) Object {
	return sendMsg(obj, variadic.F_msgSend, selector, args...)
}

// func (obj object) SendMsgStret(ret uintptr, selector string, args ...interface{}) {
// 	sel := selectorWithName(selector)
// 	C.Debug_MsgSend_Stret(unsafe.Pointer(ret), unsafe.Pointer(obj.Pointer()), sel)
// }

func (obj object) SendSuper(selector string, args ...interface{}) Object {
	return sendMsg(obj, variadic.F_msgSendSuper, selector, args...)
}
