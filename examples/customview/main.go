package main

import (
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type MyCustomView struct {
	objc.Object `objc:"MyCustomView : NSView"`
	counter     int
}

func (v *MyCustomView) awakeFromNib() {
	v.SendSuper("awakeFromNib")

	v.counter = rand.Int()
}

func (v *MyCustomView) drawRect(*core.NSRect) {
	println("draw", v.counter)
	v.counter++
	label := core.NSArray{Object: v.Send("subviews")}.ObjectAtIndex(0)
	label.Send("setStringValue:", core.NSString_FromString(fmt.Sprint(v.counter)))
}

func init() {
	rand.Seed(time.Now().UnixNano())

	c := objc.NewClassFromStruct(MyCustomView{})
	c.AddMethod("awakeFromNib", (*MyCustomView).awakeFromNib)
	c.AddMethod("drawRect:", (*MyCustomView).drawRect)
	objc.RegisterClass(c)
}

//go:generate ibtool --compile app.nib app.xib

//go:embed app.nib
var nibFile []byte

func main() {
	var contentView cocoa.NSView

	var app cocoa.NSApplication
	app = cocoa.NSApp_WithDidLaunch(func(objc.Object) {
		data := core.NSData_WithBytes(nibFile, uint64(len(nibFile)))
		nib := cocoa.NSNib{Object: objc.Get("NSNib").Alloc().Send("initWithNibData:bundle:", data, cocoa.NSBundle_Main())}

		var ptr uintptr
		ok := nib.Send("instantiateWithOwner:topLevelObjects:", app, &ptr).Bool()
		if !ok {
			log.Fatal("Failed to load NIB")
		}

		objs := core.NSArray{Object: objc.ObjectPtr(ptr)}
		for i, n := uint64(0), objs.Count(); i < n; i++ {
			obj := objs.ObjectAtIndex(i)
			if obj.Class().String() == "NSWindow" {
				contentView = cocoa.NSWindow{Object: obj}.ContentView()
			}
		}

		if contentView.Pointer() == 0 {
			log.Fatal("Could not locate window")
		}

		if contentView.Class().String() != "MyCustomView" {
			log.Fatal("Content view is the wrong type")
		}

		app.ActivateIgnoringOtherApps(true)
	})

	go func() {
		for range time.Tick(time.Second / 2) {
			contentView.Send("setNeedsDisplay:", true)
		}
	}()

	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)

	app.Run()
}
