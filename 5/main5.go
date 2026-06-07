// context practice (context.WithTimeout/context.WithCancel)
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	// res := load2()
	// select {
	// case x := <-res:
	// 	fmt.Println(x)
	// case <-ctx.Done():
	// 	fmt.Println("Cancel operation:", ctx.Err())
	// }
	ctxWithCancel, cancelFunction := context.WithCancel(context.Background())
	defer cancelFunction()
	ch := make(chan int)

	go load3(ch)
	go func() {
		for range 20 {
			select {
			case num := <-ch:
				fmt.Println(num)
				if num == 13 {
					fmt.Println("Bad number")
					cancelFunction()
				}
			case <-ctxWithCancel.Done():
				fmt.Println("Cancel operation:", ctxWithCancel.Err())
				return
			}
		}
	}()
	time.Sleep(3 * time.Second)
}

func load3(ch chan int) {
	for i := range 20 {
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
}

// func load2() chan string {
// 	res := make(chan string)
// 	go func() {
// 		fmt.Println("download the file, please wait")
// 		time.Sleep(5 * time.Second)
// 		res <- "file is downloaded"
// 	}()
// 	return res
// }
