//С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из
//которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться, что при
//каждом запуске программы итоговое число равно 1000.
package main

import (
	""
	"sync"
	"time"
)

func main() {
	chanResult := make(chan struct{}, 100)
	var (
		counter int
		mutex   = sync.Mutex{}
	)

	counter = 0

	for i := 0; i < 1001; i++ {
		chanResult <- struct{}{}

		go func(i int) {
			mutex.Lock()
			counter++
			mutex.Unlock()
			defer func() {
				<-chanResult
			}()
			time.Sleep(30 * time.Millisecond)

		}(i)

	}

	.Println("Число выполненных горутин = ", counter)

}
