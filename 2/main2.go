//Передача случайных чисел в канал и подсчет суммы за определенное время
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chInt1 := make(chan int)
	chInt2 := make(chan int)

	go load1(chInt1)
	go load2(chInt2)

	timer := time.After(1 * time.Second)
	sum := 0
	isActive := true
	for isActive {
		select {
		case msg1 := <-chInt1:
			sum += msg1
		case msg2 := <-chInt2:
			sum += msg2
		case <-timer:
			fmt.Println("timeout!")
			isActive = false
		}
	}
	fmt.Println(sum)
}

func load1(ch chan int) {
	for {
		val := rand.Intn(100)
		ch <- val
		time.Sleep(10 * time.Millisecond)
	}
}

func load2(ch chan int) {
	for {
		val := rand.Intn(100)
		ch <- val
		time.Sleep(10 * time.Millisecond)
	}
}
