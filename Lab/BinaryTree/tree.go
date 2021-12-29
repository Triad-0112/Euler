package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	recursiveWalk(t, ch)
	close(ch)
}
func recursiveWalk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	recursiveWalk(t.Left, ch)
	ch <- t.Value
	recursiveWalk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		noMoreValuesInEitherTree := !(ok1 && ok2)
		if noMoreValuesInEitherTree {
			return true
		}

		isMoreNodesInOneTree := !(ok1 || ok2)
		if isMoreNodesInOneTree {
			return false
		}

		if v1 != v2 {
			return false
		}

		// v1 and v2 exist and are equal to each other, so continue
	}
}

func main() {
	fmt.Println("Test 1 (expected true):", Same(tree.New(1), tree.New(1)))
	fmt.Println("Test 2 (expected false):", Same(tree.New(1), tree.New(2)))
}
