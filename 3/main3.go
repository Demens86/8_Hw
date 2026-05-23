//Имитация долгой операции с таймаутом
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	timer := time.NewTimer(2*time.Second)
	go func() {
		time.Sleep(5*time.Second)
		ch <- true
	}()

	select {
	case <- ch:
		fmt.Println("operation done")
	//case <-time.After(2*time.Second):
	//	fmt.Println("timeout!")
	case <- timer.C:
		fmt.Println("time is gone!")
	}
}

