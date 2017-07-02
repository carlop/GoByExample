// We can use channels to synchronize execution across goroutines. Here's an
// example of using a blocking receive to wait for a goroutine to finish.

package main

import (
	"fmt"
	"time"
)

// This is the function we'll run in a goroutine. The done channel will be used
// to notify another goroutine thath this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	// Start a worker goroutine, givint it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done

}
