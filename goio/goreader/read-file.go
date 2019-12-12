package main

import (
	"io"
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

var textFile string = "../test/lorem-ipsum.txt"

func main() {

	// Method1: Open a file and read all content as a byte array.
	Method1()

	// Method2: Read the file by ioutil.ReadFile directly.
	Method2()

	// Method3: Using bufio.Reader ReadBytes method
	Method3()

	// Method4: Using bufio.Reader ReadString method
	Method4()

	// Method5: Using bufio.Scanner Scan method
	Method5()
}

// Method1: Open a file and read all content as a byte array.
func Method1() {

	// Use os.Open to open a file, the first return value is os.File.
	// os.File implements io.Reader interface, so we can pass it to any
	// functions which accept io.Reader as a argument.
	f, err := os.Open(textFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// This method is not good for a big file, because it will consume
	// too many memory. Or even can't be fit in the whole memory.
	// But its the simple and fast way for a simple small file.
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Convert the bytes into string and fomrat print.
	// Notes that in GO, string is not thing but a byte array. so we can
	// simply convert them back and forth.
	fmt.Println("ioutil.ReadAll")
	fmt.Println(string(bytes))
}

// Method2: Read the file by ioutil.ReadFile directly.
func Method2() {

	bytes, err := ioutil.ReadFile(textFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("ioutil.ReadFile")
	fmt.Println(string(bytes))
}

// Method3: Using bufio.Reader ReadBytes method
func Method3() {

	f, err := os.Open(textFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a bufio.Reader from a io.Reader interface.
	reader := bufio.NewReader(f)
	
	// Keep reading until the character newline or any error occurs.
	// Must notice that this method will have byte array contains newline
	// character too. If you dont want the tail newline character. Please
	// use string slice like str[:len(str)-1].
	fmt.Println("bufio.Reader ReadBytes")
	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(string(bytes))
	}
}

// Method4: Using bufio.Reader ReadString method
func Method4() {

	f, err := os.Open(textFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a bufio.Reader from a io.Reader interface.
	reader := bufio.NewReader(f)
	
	// Keep reading until the character newline or any error occurs.
	// Must notice that this method will have byte array contains newline
	// character too. If you dont want the tail newline character. Please
	// use string slice like str[:len(str)-1].
	fmt.Println("bufio.Reader ReadBytes")
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(str)
	}
}

// Method5: Using bufio.Scanner Scan method
func Method5() {

	f, err := os.Open(textFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a bufio.Scanner from underlying io.Reader interface
	scanner := bufio.NewScanner(f)

	// scanner Scan idiom
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		text := scanner.Text()
		fmt.Println(text)
	}
}