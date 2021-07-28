package objc

import (
	"unsafe"

	"github.com/progrium/macdriver/misc/variadic"
)

/*
#cgo LDFLAGS: -lobjc
#import <objc/message.h>

void * addr_msgSend = &objc_msgSend;
void * addr_msgSendSuper = &objc_msgSendSuper;
void * addr_msgSend_stret = &objc_msgSend_stret;
void * addr_msgSendSuper_stret = &objc_msgSendSuper_stret;
*/
import "C"

type function int

const (
	msgSend function = iota
	msgSendSuper
	msgSend_stret
	msgSendSuper_stret
)

func (f function) String() string {
	switch f {
	case msgSend:
		return "objc_msgSend"
	case msgSendSuper:
		return "objc_msgSendSuper"
	case msgSend_stret:
		return "objc_msgSend_stret"
	case msgSendSuper_stret:
		return "objc_msgSendSuper_stret"
	}
	panic("unknown function")
}

func (f function) IsSuper() bool {
	switch f {
	case msgSendSuper, msgSendSuper_stret:
		return true
	default:
		return false
	}
}

func (f function) IsStret() bool {
	switch f {
	case msgSend_stret, msgSendSuper_stret:
		return true
	default:
		return false
	}
}

func (f function) StructReturn() function {
	switch f {
	case msgSend:
		return msgSend_stret
	case msgSendSuper:
		return msgSendSuper_stret
	default:
		return f
	}
}

func (f function) Addr() unsafe.Pointer {
	switch f {
	case msgSend:
		return C.addr_msgSend
	case msgSendSuper:
		return C.addr_msgSendSuper
	case msgSend_stret:
		return C.addr_msgSend_stret
	case msgSendSuper_stret:
		return C.addr_msgSendSuper_stret
	}
	panic("unknown function")
}

func (f function) NewCall() *variadic.FunctionCall {
	return variadic.NewFunctionCallAddr(f.Addr())
}
