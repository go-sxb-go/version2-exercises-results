package channelbis

import (
	"fmt"
	"time"
)

// Create a channel to let Ping and Pong communicate and start both routines
// The function should return a channel which will stop the ping routine when closed
func PingPong() chan bool {
	stop := make(chan bool)
	ping := make(chan bool)
	go Ping(ping, stop)
	go Pong(ping)
	return stop
}

// Method which sends 'true' every second until the channel stop is closed
func Ping(c chan<- bool, stop chan bool) {
	defer close(c)
	for {
		select {
		case <-stop:
			return
		default:
			c <- true
			time.Sleep(time.Second)
		}
	}
}

// Method which reads from channel and display "PONG"
// When channel is stopped, should display "END"
func Pong(c <-chan bool) {
	for range c {
		fmt.Println("PONG")
	}
	fmt.Println("END")
}
