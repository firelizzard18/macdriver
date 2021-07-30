package objc

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

const firstInt = 0
const maxInt = 6
const firstFloat = maxInt
const maxFloat = 8

type callArgs struct {
	Registers []uintptr
	NumInts   int
	NumFloats int
	Stack     []uintptr
}

func (a *callArgs) addIntReg(v uintptr) {
	if a.NumInts >= maxInt {
		panic("too many integer arguments!")
	}
	a.Registers[firstInt+a.NumInts] = v
	a.NumInts++
}

func (a *callArgs) addFloatReg(v uintptr) {
	if a.NumFloats >= maxFloat {
		panic("too many float arguments!")
	}
	a.Registers[firstFloat+a.NumFloats] = v
	a.NumFloats++
}

func (a *callArgs) pushBoolVal(v interface{}) {
	switch v := v.(type) {
	case bool:
		if v {
			a.addIntReg(1)
		} else {
			a.addIntReg(0)
		}
		return
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			a.addIntReg(1)
		} else {
			a.addIntReg(0)
		}
		return
	}

	panic(fmt.Errorf("expected a bool, got %T", v))
}

func (a *callArgs) pushIntVal(v interface{}) {
	switch v := v.(type) {
	case bool:
		if v {
			a.addIntReg(1)
		} else {
			a.addIntReg(0)
		}
		return
	case int:
		a.addIntReg(uintptr(v))
		return
	case uint:
		a.addIntReg(uintptr(v))
		return
	case int8:
		a.addIntReg(uintptr(v))
		return
	case uint8:
		a.addIntReg(uintptr(v))
		return
	case int16:
		a.addIntReg(uintptr(v))
		return
	case uint16:
		a.addIntReg(uintptr(v))
		return
	case int32:
		a.addIntReg(uintptr(v))
		return
	case uint32:
		a.addIntReg(uintptr(v))
		return
	case int64:
		a.addIntReg(uintptr(v))
		return
	case uint64:
		a.addIntReg(uintptr(v))
		return
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a.addIntReg(uintptr(rv.Int()))
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		a.addIntReg(uintptr(rv.Uint()))
		return
	case reflect.Bool:
		if rv.Bool() {
			a.addIntReg(1)
		} else {
			a.addIntReg(0)
		}
		return
	}

	panic(fmt.Errorf("expected an int, got %T", v))
}

func (a *callArgs) pushFloat32Val(v interface{}) {
	switch v := v.(type) {
	case float32:
		a.addFloatReg(uintptr(math.Float32bits(v)))
		return
	case float64:
		a.addFloatReg(uintptr(math.Float32bits(float32(v))))
		return
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Float32, reflect.Float64:
		a.addFloatReg(uintptr(math.Float32bits(float32(rv.Float()))))
		return
	}

	panic(fmt.Errorf("expected a float, got %T", v))
}

func (a *callArgs) pushFloat64Val(v interface{}) {
	switch v := v.(type) {
	case float32:
		a.addFloatReg(uintptr(math.Float64bits(float64(v))))
		return
	case float64:
		a.addFloatReg(uintptr(math.Float64bits(v)))
		return
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Float32, reflect.Float64:
		a.addFloatReg(uintptr(math.Float64bits(rv.Float())))
		return
	}

	panic(fmt.Errorf("expected a float, got %T", v))
}

func (a *callArgs) pushPointerVal(v interface{}) {
	if v == nil {
		a.addIntReg(0)
		return
	}

	switch v := v.(type) {
	case object:
		// avoid an indirect call if possible
		a.addIntReg(v.ptr)
		return
	case Object:
		a.addIntReg(v.Pointer())
		return
	case Class:
		a.addIntReg(v.Pointer())
		return
	case uintptr:
		a.addIntReg(v)
		return
	case unsafe.Pointer:
		a.addIntReg(uintptr(v))
		return
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.UnsafePointer:
		a.addIntReg(rv.Pointer())
		return
	case reflect.Uintptr:
		a.addIntReg(uintptr(rv.Uint()))
		return
	}

	panic(fmt.Errorf("expected a pointer, got %T", v))
}

func (a *callArgs) pushSelectorVal(v interface{}) {
	if v == nil {
		a.addIntReg(0)
		return
	}

	switch v := v.(type) {
	case selector:
		// avoid an indirect call if possible
		a.addIntReg(uintptr(selectorWithName(string(v))))
		return
	case Selector:
		a.addIntReg(uintptr(selectorWithName(v.String())))
		return
	case string:
		a.addIntReg(uintptr(selectorWithName(v)))
		return
	case uintptr:
		a.addIntReg(v)
		return
	case unsafe.Pointer:
		a.addIntReg(uintptr(v))
		return
	}

	panic(fmt.Errorf("expected a selector, got %T", v))
}

func (a *callArgs) pushStructVal(v interface{}) {
	a.pushStructVal(reflect.ValueOf(v))
}

func (a *callArgs) pushStructRval(v reflect.Value) {
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("expected a struct, got %T", v))
	}

	for i := 0; i < v.NumField(); i++ {
		v := v.Field(i)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			a.Stack = append(a.Stack, uintptr(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			a.Stack = append(a.Stack, uintptr(v.Uint()))
		case reflect.Float32, reflect.Float64:
			a.Stack = append(a.Stack, uintptr(math.Float64bits(v.Float())))
		case reflect.Ptr:
			a.Stack = append(a.Stack, v.Pointer())
		case reflect.Struct:
			a.pushStructRval(v)
		}
	}
}
