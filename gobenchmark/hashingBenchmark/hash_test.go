package hash

import (
	"path/filepath"
	"testing"
)

func BenchmarkNaiveStringHash(b *testing.B) {

	str1 := "hello"
	str2 := "world"

	for i := 0 ; i <= b.N ; i++ {
		naiveStringHash(str1, str2)
	}
}

func BenchmarkXXhashStringHash(b *testing.B) {

	str1 := "hello"
	str2 := "world"
	str := filepath.Join(str1, str2)

	for i := 0 ; i <= b.N ; i++ {
		XXhashStringHash(str)
	}
}