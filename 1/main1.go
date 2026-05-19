package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	myCh := make(chan int, 10)

	fmt.Println("start load")
	for range 10 {
		go load1(myCh)
	}
	fmt.Println("finish load")

	fmt.Println("start print")
	for range 10 {
		go print(myCh)
		time.Sleep(500*time.Millisecond)
	}
	fmt.Println("finish print")
	
}

func load1(ch chan int) {
	val := rand.Int()
	ch <- val
}

func print(ch chan int) {
	fmt.Println(<-ch)
}