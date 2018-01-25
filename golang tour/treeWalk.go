package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){

     if t.Left != nil {
	   Walk( t.Left, ch)
	 }
	 ch <- t.Value
	 //fmt.Println(t.Value)

	 if t.Right != nil {
	    Walk(t.Right, ch)
	 }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
    ch2 := make(chan int, 10)
	go Walk(t1, ch2)

	ch3 := make(chan int, 10)
	go Walk(t2, ch3)
	for i := 0; i < 10 ; i++{
	   if <-ch2 != <-ch3 {
	      return false
	   }
	}
	return true
}
func main() {
    ch1 := make(chan int)
	go Walk(tree.New(1), ch1)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch1)
	}


    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
