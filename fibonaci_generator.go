package main 

import (
	"time"
	"fmt"
)

func generateFibonaci(n int) chan int {
	ch := make(chan int)
	
	go func() {
		a := 0 
		b := 1
		c :=  a + b 
		for i := 0; i < n; i++ {
			time.Sleep(time.Second)
			if i == 0 {
				ch <- a 
			} else if i == 1 {
				ch <- b 
			} else {
				ch <- c 
				a = b 
				b = c  
				c = a + b 			
			}
		}
		close(ch)
	}()
	return ch 
}

func main() {
	for i := range generateFibonaci(10) {
		fmt.Println("number", i)
	}
}