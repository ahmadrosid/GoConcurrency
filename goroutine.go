package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go order(i)
	}
	in := 0
	fmt.Scanln(&in)
}

func order(number int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Order", number, "process", i)
		randTime := time.Duration(rand.Intn(240))
		time.Sleep(time.Millisecond * randTime)
	}
}
