package main 

import (
	"time"
	"fmt"
)

func generateEvenNumbers(n int) chan int {
	ch := make(chan int)
	go func() {
		for i:= 1; i <= n; i++ {
			ch <- i * 2
			time.Sleep(time.Second)
		}
	}()
	return ch 
}

func main() {
	n := 10
	ch := generateEvenNumbers(n)
	for i := 0; i < n; i++ {
		num := <- ch 
		fmt.Println("Number is", num)
	} 
}