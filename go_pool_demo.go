package main 

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var pool sync.Pool 
	var wg sync.WaitGroup 

	wg.Add(2)
	go func() {
		time.Sleep(3 * time.Second)
		data1 := pool.Get()
		fmt.Println(data1)
		data2 := pool.Get()
		fmt.Println(data2)
		time.Sleep(4 * time.Second)
		data3 := pool.Get()
		fmt.Println(data3)
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second)
		pool.Put(1)
		time.Sleep(3 * time.Second)
		pool.Put(2)
		wg.Done()
	}()
	wg.Wait()
}