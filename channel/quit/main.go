// https://github.com/kevinniechen/go-concurrency-patterns/blob/master/2-1-quit-select.go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	ch := generator("Hi!", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- "Bye!"
	fmt.Printf("Generator says %s", <-quit)
}
func generator(msg string, quit chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
			case <-quit:
				quit <- "See ya!"
				return
			}
		}
	}()
	return ch
}
