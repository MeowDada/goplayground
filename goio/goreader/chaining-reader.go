package main

import (
	"io/ioutil"
	"io"
	"os"
	"hash"
	"fmt"
	"crypto/sha256"
)

var textFile string = "../test/lorem-ipsum.txt"

type HashReader struct {
	rawReader io.Reader
	hasher    hash.Hash
}

func NewHashReader(reader io.Reader, hasher hash.Hash) HashReader{
	return HashReader{
		rawReader: reader,
		hasher:    hasher,
	}
}

func (hr *HashReader) Read(p []byte) (n int, e error) {
	
	reader := hr.rawReader

	nr, re := reader.Read(p)
	if re != nil {
		return nr, re
	}

	nw, we := hr.hasher.Write(p[:nr])
	if we != nil {
		return nw, we
	}

	return n, e
}

func main() {

	f, err := os.Open(textFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a SHA256 hasher
	hasher := sha256.New()

	// Create out customized reader, which chaining an io.Reader.
	reader := NewHashReader(f, hasher)

	// Read the whole file until the end.
	_, err = ioutil.ReadAll(&reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%-20s: %x\n", "HashReader's hash", reader.hasher.Sum(nil))

	// Reset the file stream pointer to the starting point.
	if _, err := f.Seek(0, os.SEEK_SET); err != nil {
		panic(err)
	}

	// Reset the hasher too, or it will compute from old value.
	hasher.Reset()

	if _, err := io.Copy(hasher, f); err != nil {
		panic(err)
	}
	fmt.Printf("%-20s: %x\n", "Directly SHA256 hash", hasher.Sum(nil))
}