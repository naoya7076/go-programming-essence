package main

import (
	"fmt"
	"slices"
)

func main() {
	a := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	n := 50
	fmt.Printf("%v", slices.Delete(a, n, n+1))
}
