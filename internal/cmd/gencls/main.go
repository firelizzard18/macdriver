package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Accelerate
#cgo LDFLAGS: -framework Accounts
#cgo LDFLAGS: -framework AddressBook
#cgo LDFLAGS: -framework Contacts
#cgo LDFLAGS: -framework AGL
#cgo LDFLAGS: -framework AppKit
#cgo LDFLAGS: -framework Cocoa
#cgo LDFLAGS: -framework AppleScriptKit
#cgo LDFLAGS: -framework AppleScriptObjC
#cgo LDFLAGS: -framework ApplicationServices
#cgo LDFLAGS: -framework AudioToolbox
#cgo LDFLAGS: -framework AudioUnit
#cgo LDFLAGS: -framework AudioVideoBridging
#cgo LDFLAGS: -framework Automator
#cgo LDFLAGS: -framework AVFoundation
#cgo LDFLAGS: -framework AVKit
#cgo LDFLAGS: -framework CalendarStore
#cgo LDFLAGS: -framework Carbon
#cgo LDFLAGS: -framework CFNetwork
#cgo LDFLAGS: -framework CloudKit
#cgo LDFLAGS: -framework Foundation
#cgo LDFLAGS: -framework CoreData
#cgo LDFLAGS: -framework Collaboration
#cgo LDFLAGS: -framework CoreAudio
#cgo LDFLAGS: -framework CoreBluetooth
#cgo LDFLAGS: -framework CoreAudioKit
#cgo LDFLAGS: -framework CoreFoundation
#cgo LDFLAGS: -framework CoreGraphics
#cgo LDFLAGS: -framework CoreImage
#cgo LDFLAGS: -framework CoreLocation
#cgo LDFLAGS: -framework CoreMedia
#cgo LDFLAGS: -framework CoreMediaIO
#cgo LDFLAGS: -framework CoreMIDI
#cgo LDFLAGS: -framework CoreMIDIServer
#cgo LDFLAGS: -framework CoreServices
#cgo LDFLAGS: -framework CoreText
#cgo LDFLAGS: -framework CoreVideo
#cgo LDFLAGS: -framework CoreWLAN
#cgo LDFLAGS: -framework CryptoTokenKit
#cgo LDFLAGS: -framework DirectoryService
#cgo LDFLAGS: -framework OpenDirectory
#cgo LDFLAGS: -framework DiscRecording
#cgo LDFLAGS: -framework DiscRecordingUI
#cgo LDFLAGS: -framework DiskArbitration
#cgo LDFLAGS: -framework DVDPlayback
#cgo LDFLAGS: -framework EventKit
#cgo LDFLAGS: -framework ExceptionHandling
#cgo LDFLAGS: -framework FinderSync
#cgo LDFLAGS: -framework ForceFeedback
#cgo LDFLAGS: -framework FWAUserLib
#cgo LDFLAGS: -framework GameController
#cgo LDFLAGS: -framework GameKit
#cgo LDFLAGS: -framework GLKit
#cgo LDFLAGS: -framework GLUT
#cgo LDFLAGS: -framework GSS
#cgo LDFLAGS: -framework Hypervisor
#cgo LDFLAGS: -framework ICADevices
#cgo LDFLAGS: -framework ImageCaptureCore
#cgo LDFLAGS: -framework ImageIO
#cgo LDFLAGS: -framework IMServicePlugIn
#cgo LDFLAGS: -framework InputMethodKit
#cgo LDFLAGS: -framework InstallerPlugins
#cgo LDFLAGS: -framework InstantMessage
#cgo LDFLAGS: -framework IOBluetooth
#cgo LDFLAGS: -framework IOBluetoothUI
#cgo LDFLAGS: -framework IOKit
#cgo LDFLAGS: -framework IOSurface
#cgo LDFLAGS: -framework JavaScriptCore
#cgo LDFLAGS: -framework WebKit
#cgo LDFLAGS: -framework Kerberos
#cgo LDFLAGS: -framework LatentSemanticMapping
#cgo LDFLAGS: -framework LDAP
#cgo LDFLAGS: -framework LocalAuthentication
#cgo LDFLAGS: -framework MapKit
#cgo LDFLAGS: -framework MediaAccessibility
#cgo LDFLAGS: -framework MediaLibrary
#cgo LDFLAGS: -framework Metal
#cgo LDFLAGS: -framework MetalKit
#cgo LDFLAGS: -framework ModelIO
#cgo LDFLAGS: -framework MultipeerConnectivity
#cgo LDFLAGS: -framework NetFS
#cgo LDFLAGS: -framework NetworkExtension
#cgo LDFLAGS: -framework NotificationCenter
#cgo LDFLAGS: -framework OpenAL
#cgo LDFLAGS: -framework OpenCL
#cgo LDFLAGS: -framework OpenGL
#cgo LDFLAGS: -framework OSAKit
#cgo LDFLAGS: -framework PCSC
#cgo LDFLAGS: -framework PreferencePanes
#cgo LDFLAGS: -framework QTKit
#cgo LDFLAGS: -framework Quartz
#cgo LDFLAGS: -framework QuartzCore
#cgo LDFLAGS: -framework QuickLook
#cgo LDFLAGS: -framework Ruby
#cgo LDFLAGS: -framework SceneKit
#cgo LDFLAGS: -framework ScreenSaver
#cgo LDFLAGS: -framework ScriptingBridge
#cgo LDFLAGS: -framework Security
#cgo LDFLAGS: -framework SecurityFoundation
#cgo LDFLAGS: -framework SecurityInterface
#cgo LDFLAGS: -framework ServiceManagement
#cgo LDFLAGS: -framework Social
#cgo LDFLAGS: -framework SpriteKit
#cgo LDFLAGS: -framework StoreKit
#cgo LDFLAGS: -framework SyncServices
#cgo LDFLAGS: -framework System
#cgo LDFLAGS: -framework SystemConfiguration
#cgo LDFLAGS: -framework Tcl
#cgo LDFLAGS: -framework Tk
#cgo LDFLAGS: -framework TWAIN
#cgo LDFLAGS: -framework VideoDecodeAcceleration
#cgo LDFLAGS: -framework VideoToolbox
#define OBJC2_UNAVAILABLE
#include <objc/objc.h>
#include <objc/runtime.h>
#include <Foundation/Foundation.h>

bool isaNSObject(Class cls) {
	if (!class_getClassMethod(cls, @selector(isSubclassOfClass:))) {
		return NO;
	}
	return [cls isSubclassOfClass:[NSObject class]];
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unsafe"
)

const nameNSObject = "NSObject\x00"

func main() {
	n := C.objc_getClassList(nil, 0)
	classes := make([]C.Class, n)
	C.objc_getClassList(&classes[0], n)

	count := map[string]int{}
	for _, class := range classes {
		cname := C.GoString(C.class_getName(class))
		if strings.HasPrefix(cname, "_") {
			continue
		}

		// if !C.isaNSObject(class) {
		// 	continue
		// }

		for _, m := range getMethods(class) {
			mname := C.GoString(C.sel_getName(C.method_getName(m)))
			if strings.HasPrefix(mname, "_") {
				continue
			}

			enc := C.GoString(C.method_getTypeEncoding(m))
			enc = simplifyTypeInfo(enc)
			count[enc]++

			if enc == "@?@:" {
				fmt.Println(cname, mname)
				if count[enc] == 100 {
					return
				}
			}
		}
	}

	type Count struct {
		Encoding string
		Count    int
	}
	var foo []Count
	for enc, c := range count {
		foo = append(foo, Count{enc, c})
	}

	sort.Slice(foo, func(i, j int) bool { return foo[i].Count > foo[j].Count })

	for _, c := range foo[:25] {
		fmt.Println(c.Count, c.Encoding)
	}
}

func getMethods(cls C.Class) []C.Method {
	var count C.uint
	ptr := C.class_copyMethodList(cls, &count)
	defer C.free(unsafe.Pointer(ptr))

	var v []C.Method
	vh := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	vh.Data = uintptr(unsafe.Pointer(ptr))
	vh.Len = int(count)
	vh.Cap = int(count)

	u := make([]C.Method, count)
	copy(u, v)
	return u
}

// simplifyTypeInfo returns a simplified typeInfo representation
// with C specifiers and stack information stripped out.
func simplifyTypeInfo(typeInfo string) string {
	ti := typeInfo
	sti := []rune{}
	for _, r := range ti {
		if r >= '0' && r <= '9' {
			continue
		}
		if r == 'r' {
			continue
		}
		// fixme(mkrautz): What is V? The NSObject release method uses V.
		if r == 'V' {
			continue
		}
		sti = append(sti, r)
	}
	return string(sti)
}
