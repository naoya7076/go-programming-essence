package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx0 := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx0)
	go func(ctx1 context.Context) {
		<-ctx1.Done()
		fmt.Println("G1-1 canceled")
	}(ctx1)

	go func(ctx1 context.Context) {
		<-ctx1.Done()
		fmt.Println("G1-2 canceled")
	}(ctx1)

	cancel1()
	time.Sleep(time.Second)
}
