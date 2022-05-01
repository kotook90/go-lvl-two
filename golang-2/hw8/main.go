//Написать программу, которая отменяет действие функции по кнтексту
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Microsecond)
	defer cancel()
	chanResult := make(chan bool)

	result(ctx, chanResult)
}

func result(ctx context.Context, chanResult chan bool) {

	go func(ctx context.Context, chanResult chan bool) {
		for i := 0; i < 10000000000000001; i++ {
			fmt.Println(i)
			if i == 1000000000000000 {
				fmt.Println("Горутина успела все сделать")
				chanResult <- true
				return
			}
		}

	}(ctx, chanResult)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-chanResult:
			fmt.Println("Функция выполнила свою работу")
			return
		default:
			continue
		}
	}

}
