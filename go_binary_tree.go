package main 

import "fmt"

type TreeNode struct {
	data int 
	left *TreeNode
	right *TreeNode 
}

func  create(i int, ch chan interface{}, n int) {
	if i >= n {
		ch <- nil  
		return 
	}
	t := TreeNode{data:i,}
	//fmt.Println(t.data)
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	go create(2 * i + 1, ch1, n)
	go create(2 * i + 2, ch2, n)
	left := <-ch1 
	right := <-ch2
	if left != nil {
		tLeft := left.(TreeNode)
		t.left = &tLeft 
	}
	if right != nil {
		tRight := right.(TreeNode)
		t.right = &tRight
	}
	fmt.Println("create ", i)
	ch <- t 
}

func inorder(t *TreeNode) {
	if t == nil {
		return 
	}
	left := t.left 
	right := t.right 
	inorder(left)
	fmt.Println(t.data)
	inorder(right)
}

func main() {
	ch := make(chan interface{})
	create(0, ch, 15)
	t := <- ch
	tree := t.(TreeNode)
	inorder(&tree)
}