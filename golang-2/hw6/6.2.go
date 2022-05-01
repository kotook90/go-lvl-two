//Написать многопоточную программу, в которой будет использоваться явный вызов
//планировщика. Выполните трассировку программы
//GOMAXPROCS=1 go run 6.2.go 2>trace.out
// go tool trace trace.out

package main

import (
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg := sync.WaitGroup{}
	for i := 0; i < 1<<4; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1e8; i += 1 {
			}
		}()
	}
	wg.Wait()
}
