package main 

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup 
	var o sync.Once
	wg.Add(3) 
	go func() {
		time.Sleep(time.Second)
		o.Do(func() {
			fmt.Println("executed once")
		})
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		o.Do(func() {
			fmt.Println("execute once")
		})
		fmt.Println("wasn't executed after 2 seconds")
		wg.Done()
	}()
	go func() {
		time.Sleep(3 * time.Second)
		o.Do(func() {
			fmt.Println("execute once")
		})
		fmt.Println("wasn't executed after 3 seconds")
		wg.Done()
	}()
	wg.Wait()
}