package main 

import (
	"time"
	"fmt"
	"io/ioutil"
	"net/http"
)
func main() {
	ch := make(chan string)
	go func() {
		for i := 0; i <= 10; i++ {
			resp, err := http.Get("http://localhost:8000/data")
			if err == nil {
				data, er := ioutil.ReadAll(resp.Body)
				if er == nil {
					sb := string(data)
					ch <- sb 
					time.Sleep(2 * time.Second)
				} 
			}
			
		}
		
	}()

	for i := 0; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}