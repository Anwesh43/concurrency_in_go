package main 

import "fmt"

func mergeSort(arr []int) []int {
	l := len(arr)
	if (l == 1) {
		return arr 
	}
	ch1 := make(chan []int)
	go func() {
		ch1 <- mergeSort(arr[0 : l / 2])
	}()
	arr1 := <- ch1 
	arr2 :=mergeSort(arr[l / 2 : l])
	arr3 := make([]int, l)

	m := len(arr1) 
	n := len(arr2)
	k := 0
	i := 0
	p := 0
	for j := 0; ; j++ {
		if arr1[p] < arr2[i] {
			arr3[k] = arr1[p]
			k++
			p++
			if p == m {
				break 
			}
		} else {
			arr3[k] = arr2[i]
			k++ 
			i++
			if  i == n {
				break 
			}
		}
	}
	arrk := arr2
	jk := i 
	mk := n
	if p < m {
		arrk = arr1 
		jk = p
		mk = m
	}
	for e:=jk; e < mk; e++ {
		arr3[k] = arrk[e]
		k++  
	}
	return arr3
}

func main() {
	arr := []int{4, 3, 1, 2, 6, 5, 10, 8, 9, 7, 20, 11, 15, 19, 17, 18, 14, 12, 13, 16}
	sortedArr := mergeSort(arr)
	for _,k := range sortedArr {
		fmt.Println(k)
	}
}