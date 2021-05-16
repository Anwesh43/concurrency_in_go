package main 

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		fmt.Println(<-ch2)
		ch1 <- 10
	}()

	go func() {
		fmt.Println(<-ch1)
		ch2 <- 20
		ch3 <- 1
	}()
	<- ch3
}