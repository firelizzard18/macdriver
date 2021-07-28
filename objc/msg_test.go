package objc

import "testing"

func BenchmarkSendMsg(b *testing.B) {
	obj := Get("NSObject").Send("new")
	defer obj.Release()

	for i := 0; i < b.N; i++ {
		sendMsg(obj, "objc_msgSend", "description")
	}
}
