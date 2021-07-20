package main 

import (
	"time"
	"fmt"
)

func consume(input chan int) chan int {
	ch := make(chan int)
	go func() {
		for data := range(input) {
			ch <- data 
		}
		close(ch)
	}()
	return ch 
}



func getFirst10Numbers() chan int {
	ch := make(chan int)
	go func() {
		for i:= 0; i < 10; i++ {
			ch <- i 
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	return ch 
}



func main() {
	ch := getFirst10Numbers()
	o1 := consume(ch)
	o2 := consume(ch)
	for i := 0; i < 10; i++ {
		select {
			case data1 := <- o1:
				fmt.Println(data1)
			case data2 := <- o2:
				fmt.Println(data2)
				
		}	
	}	
}