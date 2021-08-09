package objc

import (
	"fmt"
	"unsafe"
)

type notAnObject struct {
	value interface{}
}

func (notAnObject) Pointer() unsafe.Pointer                               { panic("not an object") }
func (notAnObject) Send(selector string, args ...interface{}) Object      { panic("not an object") }
func (notAnObject) SendSuper(selector string, args ...interface{}) Object { panic("not an object") }
func (notAnObject) Class() Class                                          { panic("not an object") }
func (notAnObject) Alloc() Object                                         { panic("not an object") }
func (notAnObject) Init() Object                                          { panic("not an object") }
func (notAnObject) Retain() Object                                        { panic("not an object") }
func (notAnObject) Release() Object                                       { panic("not an object") }
func (notAnObject) Autorelease() Object                                   { panic("not an object") }
func (notAnObject) Copy() Object                                          { panic("not an object") }
func (notAnObject) Equals(o Object) bool                                  { panic("not an object") }
func (notAnObject) Set(setter string, args ...interface{})                { panic("not an object") }
func (notAnObject) Get(getter string) Object                              { panic("not an object") }
func (notAnObject) GetSt(getter string, ret interface{})                  { panic("not an object") }

func (o notAnObject) String() string  { return fmt.Sprint(o.value) }
func (o notAnObject) Bool() bool      { return o.value.(bool) }
func (o notAnObject) CString() string { return o.value.(string) }

func (o notAnObject) Uint() uint64 {
	switch v := o.value.(type) {
	case int:
		return uint64(v)
	case uint:
		return uint64(v)
	case int8:
		return uint64(v)
	case uint8:
		return uint64(v)
	case int16:
		return uint64(v)
	case uint16:
		return uint64(v)
	case int32:
		return uint64(v)
	case uint32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint64:
		return uint64(v)
	}
	panic("not an int")
}

func (o notAnObject) Int() int64 {
	switch v := o.value.(type) {
	case int:
		return int64(v)
	case uint:
		return int64(v)
	case int8:
		return int64(v)
	case uint8:
		return int64(v)
	case int16:
		return int64(v)
	case uint16:
		return int64(v)
	case int32:
		return int64(v)
	case uint32:
		return int64(v)
	case int64:
		return int64(v)
	case uint64:
		return int64(v)
	}
	panic("not an int")
}

func (o notAnObject) Float() float64 {
	switch v := o.value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	}
	panic("not a float")
}
