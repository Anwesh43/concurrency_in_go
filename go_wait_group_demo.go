package main 
import (
	"fmt"
	"time"
	"sync"
) 

func main() {
	var wg sync.WaitGroup 
	wg.Add(3)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("after 1 second")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("after 2 seconds")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("after 3 seconds")
		wg.Done()
	}()
	fmt.Println("let's wait for 3 seconds")
	wg.Wait()
	fmt.Println("done executing the program")

}