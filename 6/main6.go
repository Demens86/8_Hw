// worker pool
// при worker = tasks результаты выведутся сразу все результаты, излишняя нагрузка?
// ----------
// при worker > tasks результаты выведутся сразу все результаты, но часть worker`ов
// будет простаивать без задачь = нет смысла
// ----------
// при worker < tasks результаты выводятся по мере обработки worker`ами, оптимальный вариант, т.к все worker`ы
// задействованы при решении задачь
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const workers = 4
	const tasks = 12
	jobs := make(chan int, tasks)
	results := make(chan string, tasks)

	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 1; w <= workers; w++ {
		go worker(w, jobs, results, &wg)
	}

	for v := range tasks {
		jobs <- v
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
		//time.Sleep(100 * time.Millisecond)
	}
}

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(100 * time.Millisecond)
		results <- fmt.Sprintf("worker:%d, %d*2 = %d", id, j, j*2)
	}
}
