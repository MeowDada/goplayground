package mybenchmark

import (
	"testing"
)

func BenchmarkMarshalFoobar(b *testing.B) {

	fb := Foobar{"foobar",100}

	for n := 0 ; n < b.N ; n++ {
		MarshalFoobar(fb)
	}
}

func BenchmarkEncodeFoobar(b *testing.B) {
	
	stream := `{"Foo":"foobar","bar":100}`

	for n := 0 ; n < b.N ; n++ {
		EncodeFoobar(stream)
	}
}

func BenchmarkUnmarshalFoobar(b *testing.B) {

	stream := `{"Foo":"foobar","bar":100}`

	for n := 0 ; n < b.N ; n++ {
		UnmarshalFoobar(stream)
	}
}

func BenchmarkDecodeFoobar(b *testing.B) {

	stream := `[{"Foo":"foobar","bar":100}]`

	for n := 0 ; n < b.N ; n++ {
		DecodeFoobar(stream)
	}
}