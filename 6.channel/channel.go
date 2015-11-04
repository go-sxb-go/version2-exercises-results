package channel

import (
	"fmt"
	"time"
)

// Create a channel to let Ping and Pong communicate and start both routines
func PingPong() {
	c := make(chan bool)
	go Ping(c)
	go Pong(c)
}

// Method which sends 'true' every second
func Ping(c chan<- bool) {
	for {
		c <- true
		time.Sleep(time.Second)
	}
}

// Method which reads from channel and display "PONG"
func Pong(c <-chan bool) {
	for range c {
		fmt.Println("PONG")
	}
}
