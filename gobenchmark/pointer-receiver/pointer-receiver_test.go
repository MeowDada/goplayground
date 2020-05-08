package benchmark

import (
	"testing"
)

func BenchmarkCopyStruct(b *testing.B) {
	n := normalStruct{}
	for i := 0 ; i < b.N; i++ {
		n.Add()
	}
}

func BenchmarkPointerReceiver(b *testing.B) {
	p := ptrRecvStruct{}
	for i := 0 ; i < b.N; i++ {
		p.Add()
	}
}
