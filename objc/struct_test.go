// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"
)

type NSPoint struct {
	X float32
	Y float32
}

type NSSize struct {
	Width  float32
	Height float32
}

type NSRect struct {
	Origin NSPoint
	Size   NSSize
}

func TestStructPassing(t *testing.T) {
	rect := NSRect{
		NSPoint{
			X: 0,
			Y: 0,
		},
		NSSize{
			Width:  500,
			Height: 500,
		},
	}

	obj := GetClass("NSValue").Send("valueWithRect:", rect)
	if obj.Pointer() == nil {
		t.Fatalf("unable to create NSValue, got nil ptr")
	}
	want := "NSRect: {{0, 0}, {500, 500}}"
	got := obj.String()
	if got != want {
		t.Fatalf("bad NSValue was constructed: got %v, want %v", got, want)
	}
}
