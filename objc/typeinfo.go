// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
)

const (
	encId          = '@'
	encClass       = '#'
	encSelector    = ':'
	encChar        = 'c'
	encUChar       = 'C'
	encShort       = 's'
	encUShort      = 'S'
	encInt         = 'i'
	encUInt        = 'I'
	encLong        = 'l'
	encULong       = 'L'
	encLongLong    = 'q'
	encULongLong   = 'Q'
	encFloat       = 'f'
	encDouble      = 'd'
	encDFLD        = 'b'
	encBool        = 'B'
	encVoid        = 'v'
	encUndef       = '?'
	encPtr         = '^'
	encCharPtr     = '*'
	encAtom        = '%'
	encArrayBegin  = '['
	encArrayEnd    = ']'
	encUnionBegin  = '('
	encUnionEnd    = ')'
	encStructBegin = '{'
	encStructEnd   = '}'
	encVector      = '!'
	encConst       = 'r'
)

var (
	objectInterfaceType   = reflect.TypeOf((*Object)(nil)).Elem()
	classInterfaceType    = reflect.TypeOf((*Class)(nil)).Elem()
	selectorInterfaceType = reflect.TypeOf((*Selector)(nil)).Elem()
)

// Returns the function's typeInfo
func funcTypeInfo(fn interface{}) string {
	typ := reflect.TypeOf(fn)
	kind := typ.Kind()
	if kind != reflect.Func {
		panic("not a func")
	}

	typeInfo := ""
	numOut := typ.NumOut()
	switch numOut {
	case 0:
		typeInfo += string(encVoid)
	case 1:
		typeInfo += typeInfoForType(typ.Out(0))
	default:
		panic("too many output parameters")
	}

	if typ.NumIn() == 0 {
		panic("funcTypeInfo: bad func")
	}

	typeInfo += typeInfoForType(typ.In(0))
	typeInfo += string(encSelector)

	for i := 1; i < typ.NumIn(); i++ {
		typeInfo += typeInfoForType(typ.In(i))
	}
	return typeInfo
}

type typeInfoIter struct {
	value string
	pos   int
}

func (i *typeInfoIter) Next() (byte, bool) {
	if i.pos >= len(i.value) {
		return 0, false
	}

	v := i.value[i.pos]
	switch v {
	case encArrayBegin, encUnionBegin, encStructBegin:
		i.close()
	default:
		i.pos++
	}
	return v, true
}

func (i *typeInfoIter) findNested() {
	for i.value[i.pos] != '=' {
		i.pos++
	}
	i.pos++
}

func (i *typeInfoIter) close() {
	open := i.value[i.pos]
	i.pos++

	var close byte
	switch open {
	case encArrayBegin:
		close = encArrayEnd
	case encUnionBegin:
		close = encUnionEnd
		i.findNested()
	case encStructBegin:
		close = encStructEnd
		i.findNested()
	}

	for i.value[i.pos] != close {
		switch i.value[i.pos] {
		case encArrayBegin, encUnionBegin, encStructBegin:
			i.close()
		default:
			i.pos++
		}
	}
}
