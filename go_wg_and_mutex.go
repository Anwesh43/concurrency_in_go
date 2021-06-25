package main 

import (
	"time"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup 
	var mutex sync.Mutex
	money := 100 
	wg.Add(2)
	go func() {
		time.Sleep(time.Second)
		mutex.Lock()
		money += 30
		mutex.Unlock()
		fmt.Println("after depositing", money)
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		mutex.Lock()
		money -= 10 
		mutex.Unlock()
		fmt.Println("after withdrawing", money)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("money currently", money)
}