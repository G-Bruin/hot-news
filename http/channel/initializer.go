package channel

import (
	"fmt"
	"time"
)

func Init() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	for {
		ch <- "Washington"
		ch <- "Tripoli"
		ch <- "London"
		ch <- "Beijing"
		ch <- "Tokyo"
	}

}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
