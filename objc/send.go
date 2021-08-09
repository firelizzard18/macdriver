package objc

// Code generated by 'go run ../internal/cmd/gensend'. DO NOT EDIT.

import "unsafe"

// #cgo LDFLAGS: -lobjc
// #include <objc/message.h>
//
// typedef id (*_X)(id self, SEL cmd);
// typedef id (*_XX)(id self, SEL cmd, id a0);
// typedef unsigned long long (*_Q)(id self, SEL cmd);
// typedef char (*_c)(id self, SEL cmd);
// typedef char (*_cX)(id self, SEL cmd, id a0);
// typedef long long (*_q)(id self, SEL cmd);
// typedef void (*_v)(id self, SEL cmd);
// typedef void (*_vX)(id self, SEL cmd, id a0);
// typedef void (*_vQ)(id self, SEL cmd, unsigned long long a0);
// typedef void (*_vc)(id self, SEL cmd, char a0);
//
// id sendX(id self, SEL cmd) { return ((_X)&objc_msgSend)(self, cmd); }
// id sendXX(id self, SEL cmd, id a0) { return ((_XX)&objc_msgSend)(self, cmd, a0); }
// unsigned long long sendQ(id self, SEL cmd) { return ((_Q)&objc_msgSend)(self, cmd); }
// char sendc(id self, SEL cmd) { return ((_c)&objc_msgSend)(self, cmd); }
// char sendcX(id self, SEL cmd, id a0) { return ((_cX)&objc_msgSend)(self, cmd, a0); }
// long long sendq(id self, SEL cmd) { return ((_q)&objc_msgSend)(self, cmd); }
// void sendv(id self, SEL cmd) { ((_v)&objc_msgSend)(self, cmd); }
// void sendvX(id self, SEL cmd, id a0) { ((_vX)&objc_msgSend)(self, cmd, a0); }
// void sendvQ(id self, SEL cmd, unsigned long long a0) { ((_vQ)&objc_msgSend)(self, cmd, a0); }
// void sendvc(id self, SEL cmd, char a0) { ((_vc)&objc_msgSend)(self, cmd, a0); }
//
import "C"

func objectGo2C(v Object) C.id      { return C.id(unsafe.Pointer(v.Pointer())) }
func classGo2C(v Class) C.Class     { return C.Class(unsafe.Pointer(v.Pointer())) }
func selectorGo2C(v Selector) C.SEL { return C.SEL(v.Pointer()) }

func objectC2Go(v C.id) Object      { return ObjectPtr(uintptr(unsafe.Pointer(v))) }
func classC2Go(v C.Class) Class     { return ClassPtr(uintptr(unsafe.Pointer(v))) }
func selectorC2Go(v C.SEL) Selector { return SelectorPtr(unsafe.Pointer(v)) }

// SendX calls objc_msgSend with type encoding @@:.
// WARNING! This may crash if the method's signature does not match.
func SendX(self Object, cmd Selector) Object {
	r := C.sendX(objectGo2C(self), selectorGo2C(cmd))
	return objectC2Go(r)
}

// SendXX calls objc_msgSend with type encoding @@:@.
// WARNING! This may crash if the method's signature does not match.
func SendXX(self Object, cmd Selector, a0 Object) Object {
	r := C.sendXX(objectGo2C(self), selectorGo2C(cmd), objectGo2C(a0))
	return objectC2Go(r)
}

// SendQ calls objc_msgSend with type encoding Q@:.
// WARNING! This may crash if the method's signature does not match.
func SendQ(self Object, cmd Selector) ULongLong {
	r := C.sendQ(objectGo2C(self), selectorGo2C(cmd))
	return ULongLong(r)
}

// Sendc calls objc_msgSend with type encoding c@:.
// WARNING! This may crash if the method's signature does not match.
func Sendc(self Object, cmd Selector) Char {
	r := C.sendc(objectGo2C(self), selectorGo2C(cmd))
	return Char(r)
}

// SendcX calls objc_msgSend with type encoding c@:@.
// WARNING! This may crash if the method's signature does not match.
func SendcX(self Object, cmd Selector, a0 Object) Char {
	r := C.sendcX(objectGo2C(self), selectorGo2C(cmd), objectGo2C(a0))
	return Char(r)
}

// Sendq calls objc_msgSend with type encoding q@:.
// WARNING! This may crash if the method's signature does not match.
func Sendq(self Object, cmd Selector) LongLong {
	r := C.sendq(objectGo2C(self), selectorGo2C(cmd))
	return LongLong(r)
}

// Sendv calls objc_msgSend with type encoding v@:.
// WARNING! This may crash if the method's signature does not match.
func Sendv(self Object, cmd Selector) {
	C.sendv(objectGo2C(self), selectorGo2C(cmd))
}

// SendvX calls objc_msgSend with type encoding v@:@.
// WARNING! This may crash if the method's signature does not match.
func SendvX(self Object, cmd Selector, a0 Object) {
	C.sendvX(objectGo2C(self), selectorGo2C(cmd), objectGo2C(a0))
}

// SendvQ calls objc_msgSend with type encoding v@:Q.
// WARNING! This may crash if the method's signature does not match.
func SendvQ(self Object, cmd Selector, a0 ULongLong) {
	C.sendvQ(objectGo2C(self), selectorGo2C(cmd), C.ulonglong(a0))
}

// Sendvc calls objc_msgSend with type encoding v@:c.
// WARNING! This may crash if the method's signature does not match.
func Sendvc(self Object, cmd Selector, a0 Char) {
	C.sendvc(objectGo2C(self), selectorGo2C(cmd), C.char(a0))
}

func sendMsgDirect(self Object, cmd Selector, encoding string, args []interface{}) (Object, bool) {
	switch encoding {
	case "@@:":
		r := SendX(self, cmd)
		return r, true
	case "@@:@":
		r := SendXX(self, cmd, asObject(args[0]))
		return r, true
	case "Q@:":
		r := SendQ(self, cmd)
		return notAnObject{value: r}, true
	case "c@:":
		r := Sendc(self, cmd)
		return notAnObject{value: r}, true
	case "c@:@":
		r := SendcX(self, cmd, asObject(args[0]))
		return notAnObject{value: r}, true
	case "q@:":
		r := Sendq(self, cmd)
		return notAnObject{value: r}, true
	case "v@:":
		Sendv(self, cmd)
		return notAnObject{value: struct{}{}}, true
	case "v@:@":
		SendvX(self, cmd, asObject(args[0]))
		return notAnObject{value: struct{}{}}, true
	case "v@:Q":
		SendvQ(self, cmd, ULongLong(asUint(args[0])))
		return notAnObject{value: struct{}{}}, true
	case "v@:c":
		Sendvc(self, cmd, Char(asInt(args[0])))
		return notAnObject{value: struct{}{}}, true
	}
	return nil, false
}
