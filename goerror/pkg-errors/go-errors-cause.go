package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type foo struct{}
type bar struct{}
type baz struct{}

type fooError struct {}

func (e fooError) Error() string {
	return "fooError!"
}

type barError struct {}

func (e barError) Error() string {
	return "barError!"
}

func (e barError) Cause() error {
	return &fooError{}
}

type bazError struct {}

func (e bazError) Error() string {
	return "bazError!"
}

func (e bazError) Cause() error {
	return &barError{}
}

func (f foo) FooDoSomething() error {
	return errors.New("FooDoSomething wrong!")
}

func (b bar) BarDoSomething() error {
	foo := foo{}
	if err := foo.FooDoSomething(); err != nil {
		return errors.Wrap(err, "BarDoSomething wrong!")
	}
	return nil
}

func (b baz) BazDoSomething() error {
	bar := bar{}
	if err := bar.BarDoSomething(); err != nil {
		return errors.Wrap(err, "BazDoSomething wrong!")
	}
	return nil
}

func main() {

	err := baz{}.BazDoSomething()
	if err != nil {
		fmt.Println("Print out error of baz{}.BazDoSomething() directly")
		fmt.Println(err)
		fmt.Println("print out error of baz{}.BazDoSomething() with errors.Cause(err)")
		rootErr := errors.Cause(err)
		fmt.Println(rootErr)
	}
}