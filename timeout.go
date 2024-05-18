package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		if err := ctx.Err(); errors.Is(err, context.Canceled) {
			// キャンセルされていた場合
			fmt.Println("canceled")
		} else if errors.Is(err, context.DeadlineExceeded) {
			// タイムアウトだった場合
			fmt.Println("deadline")
		}

		close(out)
		userID, authToken, traceID := ctx.Value("userID").(int), ctx.Value("autoToken").(string), ctx.Value("traceID").(int)
		fmt.Println("log: ", userID, authToken, traceID)
		fmt.Println("generator closed")
	}()
	return out
}

func main() {
	type contextKey string
	const (
		userIDKey    contextKey = "userID"
		authTokenKey contextKey = "authToken"
		traceIDKey   contextKey = "traceID"
	)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	ctx = context.WithValue(ctx, userIDKey, 2)
	ctx = context.WithValue(ctx, authTokenKey, "xxxxx")
	ctx = context.WithValue(ctx, traceIDKey, 3)
	gen := generator(ctx, 1)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		result, ok := <-gen
		if ok {
			fmt.Println(result)
		} else {
			fmt.Println("timeout")
			break LOOP
		}
	}
	cancel()

	wg.Wait()
}
