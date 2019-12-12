package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	var (
		testFile = "test-file"
	)

	os.Create(testFile)

	if err := flock(testFile); err != nil {
		fmt.Println(err)
	}

	if err := flockRemove(testFile); err != nil {
		fmt.Println(err)
	}

	if err := funlock(testFile); err != nil {
		fmt.Println(err)
	}

	if err := flockRemove(testFile); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Remove funlocked file successfully")
}

func flockRemove(path string) error {
	err := flock(path)
	if err == nil {
		os.Remove(path)
	}

	return err
}

func flock(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX | syscall.LOCK_NB)
}

func funlock(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
}