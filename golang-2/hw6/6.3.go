// Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”
// go run -race 6.3.go

package main

import (
	"fmt"
	"sync"
)

const count = 10000

func main() {
	var (
		counter int
		wg      sync.WaitGroup
	)
	ch := make(chan int, 10000)

	wg.Add(count)
	for i := 0; i < count; i += 1 {
		ch <- i
		go func() {
			defer wg.Done()
			counter += 1

		}()
		if i == 9999 {
			close(ch)
		}
	}
	wg.Wait()

	for val := range ch {
		fmt.Printf("Количество записей в канал составило %d, что говорит нам о том, что на самом деле потоков было выполнено соответственно записям\n", val)

	}
	fmt.Printf("По данным waitGroup мы выполнили %d потоков\n", counter)
}
