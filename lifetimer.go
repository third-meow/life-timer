package main

import (
	"fmt"
	"time"
)

func timekeeper(name string, mins int, status, done chan string) {
	for i := 0; i < mins; i++ {
		minTimer := time.NewTimer(time.Second)
		<-minTimer.C
	}
	done <- name
}

func main() {
	blank := make(chan string)
	done_timers := make(chan string)
	go timekeeper("jeff", 25, blank, done_timers)
	dt := <-done_timers
	fmt.Println(dt)

}
