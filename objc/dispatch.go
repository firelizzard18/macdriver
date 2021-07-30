package objc

import (
	"math"
	"sync"
	"unsafe"

	"github.com/progrium/macdriver/misc/variadic"
)

type dispatcher func(fn variadic.Function, self, sel interface{}, args []interface{}, stret interface{}) uintptr

var typeInfoDispatcher = struct {
	sync.RWMutex
	e map[string]dispatcher
}{
	e: map[string]dispatcher{},
}

func buildDispatcher(typeInfo string) dispatcher {
	typeInfoDispatcher.RLock()
	fn, ok := typeInfoDispatcher.e[typeInfo]
	typeInfoDispatcher.RUnlock()
	if ok {
		return fn
	}

	typeInfoDispatcher.Lock()
	defer typeInfoDispatcher.Unlock()

	fn, ok = typeInfoDispatcher.e[typeInfo]
	if ok {
		return fn
	}

	var pack []func(*callArgs, interface{})

	it := typeInfoIter{typeInfo, 0}
	ret, ok := it.Next()
	if !ok {
		panic("no return value")
	}

	t, ok := it.Next()
	if !ok || t != encId {
		panic("first argument must be an object")
	}

	t, ok = it.Next()
	if !ok || t != encSelector {
		panic("second argument must be a selector")
	}

	for {
		t, ok = it.Next()
		if !ok {
			break
		}

		switch t {
		case encId, encClass, encPtr:
			pack = append(pack, (*callArgs).pushPointerVal)
		case encSelector:
			pack = append(pack, (*callArgs).pushSelectorVal)
		case encBool:
			pack = append(pack, (*callArgs).pushBoolVal)
		case encChar, encShort, encInt, encLong, encLongLong,
			encUChar, encUShort, encUInt, encULong, encULongLong:
			pack = append(pack, (*callArgs).pushIntVal)
		case encFloat:
			pack = append(pack, (*callArgs).pushFloat32Val)
		case encDouble:
			pack = append(pack, (*callArgs).pushFloat64Val)
		case encStructBegin:
			pack = append(pack, (*callArgs).pushStructVal)
		}
	}

	fn = func(fn variadic.Function, obj, cmd interface{}, args []interface{}, stret interface{}) uintptr {
		if ret == encStructBegin {
			if stret == nil {
				panic("struct return pointer missing")
			}
			fn = fn.StRet()
		}

		fc := fn.NewCall()
		a := callArgs{Registers: fc.Words[:]}

		if ret == encStructBegin {
			a.pushPointerVal(stret)
		}

		a.pushPointerVal(obj)
		a.pushSelectorVal(cmd)
		for i, pack := range pack {
			pack(&a, args[i])
		}

		fc.NumFloat = int64(a.NumFloats)
		if len(a.Stack) > 0 {
			fc.NumMemory = int64(len(a.Stack))
			fc.Memory = uintptr(unsafe.Pointer(&a.Stack[0]))
		}

		switch ret {
		case encFloat:
			return uintptr(math.Float32bits(fc.CallFloat32()))
		case encDouble:
			return uintptr(math.Float64bits(fc.CallFloat64()))
		case encStructBegin:
			fc.Call()
			return 0
		default:
			return fc.Call()
		}
	}

	typeInfoDispatcher.e[typeInfo] = fn
	return fn
}
