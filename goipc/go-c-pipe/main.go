package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"syscall"
)

var pipeFile = "go-c-pipe.pipe"

func main() {
	os.Remove(pipeFile)
	err := syscall.Mkfifo(pipeFile, 0666)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(pipeFile, os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Printf("%s", string(line))
		}
	}
}