package main

func main() {

	// Below code section will fail because if the channel is created as a non-buffered channel.
	// The sender will be blocked forever if no reader is receiving from the channel.
	/*
	ch := make(chan bool)
	ch <- true
	<- ch
	*/

	// This will works because the channel is able to buffer a bool variable with length 1.
	// So the sender channel will not be blocked.
	ch := make(chan bool, 1)
	ch <- true
	<- ch
}