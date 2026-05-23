//load1 - имитация загрузки файла + возвращение результата через канал
//write1(2) - запись в канал + fanIn сбор данных в результирующий канал
package main

import (
	"fmt"
	"time"
)

func main() {

	res := load1()

	x := <-res
	fmt.Println(x)

	myCh1 := make(chan int)
	myCh2:= make(chan int)

	go write1(myCh1)
	go write2(myCh2)

	result := fanIn(myCh1, myCh2)

	for {
		x := <-result
		fmt.Println(x)
	}

}
func load1() chan string {
	res := make(chan string)
	go func() {
		fmt.Println("download the file, please wait")
		time.Sleep(5 * time.Second)
	 	res <- "file is downloaded"
	}()
	return res
}

func write1(ch chan int) {
	for {
		time.Sleep(500*time.Millisecond)
		ch <- 1
	}
}

func write2(ch chan int) {
	for {
		time.Sleep(500*time.Millisecond)
		ch <- 2
	}
}

func fanIn(ch1, ch2 chan int) chan int {
	fanout := make(chan int, 2)

	go func() {
		for {
			select {
			case x:= <-ch1:
				fanout <- x
			case x:= <- ch2:
				fanout <- x
			}
		}
	}()

	return fanout
}
