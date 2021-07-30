// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"fmt"
	"reflect"
)

func typeInfoForType(typ reflect.Type) string {
	if typ.Implements(classInterfaceType) {
		return string(encClass)
	} else if typ.Implements(objectInterfaceType) {
		return string(encId)
	} else if typ.Implements(selectorInterfaceType) {
		return string(encSelector)
	}

	kind := typ.Kind()
	switch kind {
	case reflect.Bool:
		return string(encBool)
	case reflect.Int:
		return string(encInt)
	case reflect.Int8:
		return string(encChar)
	case reflect.Int16:
		return string(encShort)
	case reflect.Int32:
		return string(encInt)
	case reflect.Int64:
		return string(encULong)
	case reflect.Uint:
		return string(encUInt)
	case reflect.Uint8:
		return string(encUChar)
	case reflect.Uint16:
		return string(encUShort)
	case reflect.Uint32:
		return string(encUInt)
	case reflect.Uint64:
		return string(encULong)
	case reflect.Uintptr:
		return string(encPtr)
	case reflect.Float32:
		return string(encFloat)
	case reflect.Float64:
		return string(encDouble)
	case reflect.Ptr:
		return string(encPtr)
	}

	panic("typeinfo: unhandled/invalid kind " + fmt.Sprintf("%v", kind) + " " + fmt.Sprintf("%v", typ))
}
