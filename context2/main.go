package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done() //go func() で起動したゴルーチンの処理が終了するのを待つためにここで実行

	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case out <- num:
			}
		}
		close(out)
		fmt.Println("generator close")
	}()
	return out
}

func main() {
	done := make(chan struct{})
	gen := generator(done, 1)
	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	close(done)
	wg.Wait()
}
