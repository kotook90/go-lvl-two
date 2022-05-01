//Написать программу, которая при получении в канал сигнала SIGTERM останавливается не
//позднее, чем за одну секунду (установить таймаут)

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	infinity(signalChanel)

}

func infinity(signalChanel chan os.Signal) {
	defer close(signalChanel)
	for i := 1; ; i++ {
		fmt.Println(i)
		time.Sleep(3 * time.Second)

		select {
		case v := <-signalChanel:
			time.After(50 * time.Microsecond)
			fmt.Printf("Получен сигнал ОС %s", v)
			return
		default:
			continue
		}
	}

}
