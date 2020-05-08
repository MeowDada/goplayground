package composite

import (
	"testing"
)

func BenchmarkCopyStruct(b *testing.B) {
	p := CopyStruct{}
	for i := 0 ; i < b.N ; i++ {
		p = p.SetVal(1).SetString("123")
	}
}

func BenchmarkPtrStruct(b *testing.B) {
	p := &PtrStruct{}
	for i := 0 ; i < b.N ; i++ {
		p = p.SetVal(5).SetString("456")
	}
}