package main 

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		ch1 <- 10
		fmt.Println(<-ch2)
		ch3 <- 1
	}()

	go func() {
		fmt.Println(<-ch1)
		ch2 <- 20
	}()
	<- ch3
}