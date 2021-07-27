package main

// #cgo CFLAGS: -x objective-c -DGL_SILENCE_DEPRECATION
// #cgo LDFLAGS: -framework OpenGL -framework Foundation
// #include <OpenGL/OpenGL.h>
// #include <AppKit/AppKit.h>
import "C"
import (
	"log"
	"sync"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

var renderFuncs sync.Map

type myOpenGLView struct {
	objc.Object `objc:"MyOpenGLView : NSOpenGLView"`
}

func init() {
	c := objc.NewClassFromStruct(myOpenGLView{})
	c.AddMethod("awakeFromNib", (*myOpenGLView).awakeFromNib)
	c.AddMethod("drawRect:", (*myOpenGLView).drawRect)
	c.AddMethod("dealloc", func(v *myOpenGLView) {
		renderFuncs.Delete(v.Pointer())
		v.SendSuper("dealloc")
	})
	objc.RegisterClass(c)
}

func (v *myOpenGLView) awakeFromNib() {
	v.SendSuper("awakeFromNib")

	pixFmtAttrs := []C.NSOpenGLPixelFormatAttribute{
		C.NSOpenGLPFAOpenGLProfile, C.NSOpenGLProfileVersion4_1Core,
		C.NSOpenGLPFAColorSize, 24,
		C.NSOpenGLPFAAlphaSize, 8,
		C.NSOpenGLPFADoubleBuffer,
		C.NSOpenGLPFAAccelerated,
		0,
	}
	pixFmt := objc.Get("NSOpenGLPixelFormat").Alloc().Send("initWithAttributes:", &pixFmtAttrs[0])
	if pixFmt.Pointer() == 0 {
		log.Fatal("Failed to create pixel format")
	}

	v.SendSuper("setPixelFormat:", pixFmt)
	pixFmt.Release()
}

func (v *myOpenGLView) drawRect(r *core.NSRect) {
	v.SendSuper("drawRect:", r)

	fn, ok := renderFuncs.Load(v.Pointer())
	if !ok {
		return
	}

	ok = fn.(glRenderFunc)(cocoa.NSView{Object: v.Object})
	if ok {
		v.Send("openGLContext").Send("flushBuffer")
	}
}

type glRenderFunc func(cocoa.NSView) bool

func setGLRenderFunc(view cocoa.NSView, fn glRenderFunc) {
	if view.Class().String() != "MyOpenGLView" {
		panic("view is not a MyOpenGLView")
	}
	renderFuncs.Store(view.Pointer(), fn)
}
