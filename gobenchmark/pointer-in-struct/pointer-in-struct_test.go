package benchmark

import (
	"testing"
)

func BenchmarkPointerRecvStruct(b *testing.B) {
	s := ptrRecvStruct{}
	for i := 0 ; i < b.N ; i++ {
		s.changeValue(1,2,3)
	}
}

func BenchmarkPointerInStruct(b *testing.B) {
	s := ptrInStruct{}
	x, y, z := 1, 2, 3
	for i := 0 ; i < b.N ; i++ {
		s.changeValue(&x,&y,&z)
	}
}