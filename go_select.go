package main 

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		for i:=0; i < 5; i++ {
			ch1 <- "slept for 1 second"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for i:=0; i < 5; i++ {
			ch2 <- "slept for 3 seconds"
			time.Sleep(time.Second * 3)
		}
	}()
	
	for i := 0; i < 10; i++ {
		select {
			case message1 := <- ch1:
				fmt.Println(message1)
			case message2 := <- ch2:
				fmt.Println(message2)
		}
	}
}