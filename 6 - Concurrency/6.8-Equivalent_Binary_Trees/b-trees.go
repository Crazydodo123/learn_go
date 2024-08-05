package main

// import (
// 	"fmt"
// 	"golang.org/x/tour/tree"
// )

// // Walk walks the tree t sending all values
// // from the tree to the channel ch
// func Walk(t *tree.Tree, ch chan int) {
// 	if t.Left != nil {
// 		Walk(t.Left, ch)
// 	}
// 	ch <- t.Value
// 	if t.Right != nil {
// 		Walk(t.Right, ch)
// 	}
// }

// // Same determines whether the trees
// // ti and t2 contain the same values
// func Same(t1, t2 *tree.Tree) bool {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	Walk(t1, ch1)
// 	Walk(t2, ch2)

// 	for range 10 {
// 		if <-ch1 != <-ch2 {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	// ch := make(chan int)
	// go Walk(tree.New(1), ch)
	// for range 10 {
	// 	fmt.Println(<-ch)
	// }
}
