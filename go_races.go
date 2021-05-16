package main 

import "fmt"

func addToNumber(number *int, a int) {
	*number += a 
}

func subToNumber(number *int, b int) {
	*number -= b
}

func main() {
	number := 0
	
	go addToNumber(&number, 50)
	subToNumber(&number, 10)
	fmt.Println(number)
}