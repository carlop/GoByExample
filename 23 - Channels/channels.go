// Channels are the pipes thath connect concurrent goroutines. You can send
// values into channels from one goroutine and receive those values into another
// goroutine

package main

import "fmt"

func main() {
	// Create a new channel with make(chan val-type). Channels are typed by the
	// values they convey.
	messages := make(chan string)

	// SEnd a value into a channel usint the channel -> syntax. Here we send a
	// "ping" to the messages channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value fro the channel. Here we'll receive
	// the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
}
