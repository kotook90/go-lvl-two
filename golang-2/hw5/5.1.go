//Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех
//Реализуйте функцию для разблокировки мьютекса с помощью defer

package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1000)
	ch := make(chan int, 1001)

	for i := 1; i < 1001; i++ {
		ch <- i

		go func(i int) {
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			counter += 1

		}(i)
		if i == 1000 {
			close(ch)
		}

	}

	wg.Wait()
	sum(ch)
	fmt.Println("Число выполненных потоков равно: ", counter)
}

func sum(ch chan int) {

	for val := range ch {
		fmt.Printf("Сигналов, поданых в канал %d\n", val)
	}

}
