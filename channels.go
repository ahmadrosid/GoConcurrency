package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go sender(c)
	go ponger(c)
	go receiver(c)
	in := 0
	fmt.Scanln(&in)
}

func sender(c chan <- string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}

func ponger(c chan <- string) {
	c <- "Pong"
	c <- "Pong"
}

func receiver(c <- chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
