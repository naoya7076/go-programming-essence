package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}
	c := make(chan int)
	for _, s := range src {
		go func(s int) {
			result := s * 2
			c <- result
		}(s)
	}
	for range src {
		num := <-c
		dst = append(dst, num)
	}

	// var mu sync.Mutex
	// for _, s := range src {
	// 	go func(s int) {
	// 		result := s * 2
	// 		mu.Lock()
	// 		dst = append(dst, result)
	// 		mu.Unlock()
	// 	}(s)
	// }
	// time.Sleep(time.Second)
	fmt.Println(dst)
}
