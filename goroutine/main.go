package main

import (
	"context"
	"fmt"
	"sync"
	"time"
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

// 拘束パターン
func restFunc() <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)

		for i := 0; i < 5; i++ {
			result <- 1
		}
	}()
	return result
}

// select文
func selectStatement() {
	gen1, gen2 := make(chan int), make(chan int)
	select {
	case num := <-gen1:
		fmt.Println(num)
	case num := <-gen2:
		fmt.Println(num)
	default:
		fmt.Println("neither chan cannot use")
	}
}

func fanIn2(ctx context.Context, cs ...<-chan int) <-chan int {
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))
	for i, c := range cs {
		// FanInの対象になるチャネルごとに個別にゴールーチンを立てる
		go func(c <-chan int, i int) {
			defer wg.Done()
			for num := range c {
				select {
				case <-ctx.Done():
					fmt.Println("wg.Done", i)
					return
				case result <- num:
					fmt.Println("send", i)
				}
			}
		}(c, i)
	}
	go func() {
		wg.Wait()
		fmt.Println("closing fanIn")
		close(result)
	}()
	return result
}

func ticker() {
	t := time.NewTicker(time.Millisecond * 100)
	defer t.Stop()
	for i := 0; i < 5; i++ {
		<-t.C
		fmt.Println("tick")
	}
}

func Query(conns []Conn, query string) Result {
	ch := make(chan Result, len(conns))

	for _, conn := range conns {
		go func(c Conn) {
			ch <- c.DoQuery(query)
		}(conn)
	}
	return <-ch
}

// func main() {
// 	result := Query(conns, query)
// 	fmt.Println(result)
// }
