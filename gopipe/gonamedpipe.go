package main

import (
	"os"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"syscall"
	"os/exec"
	"bytes"
	"path/filepath"
)

func main() {

	var execerPath string

	// Defines a string flag with specified name, default value, and usage string
	flag.StringVar(&execerPath, "execer", "./execer", "path to execer")

	// Parse parses flag definitions from the argument list
	flag.Parse();

	// Create a temp directory called "named-pipes" under current directory
	tmpDir, err := ioutil.TempDir("", "named-pipes")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Remove the temp directory we created before the proecess goes down with error
	defer os.RemoveAll(tmpDir)

	// Create a named piped file as ./named-pipes/stdout with permission 0600 (rw-)
	namedPipe := filepath.Join(tmpDir, "stdout")
	syscall.Mkfifo(namedPipe, 0600)

	// A go rotuine for executing another program "execer"
	go func() {

		// exec.Command returns the Cmd struct to execute the named program with the given arguments
		cmd := exec.Command(execerPath, namedPipe)

		// Forward the standard output and execute the command
		cmd.Stdout = os.Stdout
		cmd.Run()
	}()

	// Open named pipe for reading
	fmt.Println("Opening named pipe for reading")
	stdout, err := os.OpenFile(namedPipe, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Reading...")

	var buff bytes.Buffer
	fmt.Println("Waiting for someone to write something")
	io.Copy(&buff, stdout)
	stdout.Close()
	fmt.Printf("Data: %s\n", buff.String())
}