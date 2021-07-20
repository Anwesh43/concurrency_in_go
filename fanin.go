package main 

import (
	"time"
	"fmt"
)

func getFirstFiveNNumbers(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- (i * n)
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	return ch 
}

func fanIn() chan int {
	ch := make(chan int)
	ch1 := getFirstFiveNNumbers(2)
	ch2 := getFirstFiveNNumbers(3)
	go func() {
		for i := 0; i < 5; i++ {
			 data :=  <-ch1
			ch <- data  
		}
	}()
	go func() {
		for i := 0;  i < 5; i++ {
			data :=  <-ch2
			ch <- data 
		}
	}()
	return ch 
}

func main() {
	ch := fanIn()
	for i := 0; i < 10; i++ {
		fmt.Println("Data", <-ch)
	}
}