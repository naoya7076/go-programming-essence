package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")
	r := rand.New(rand.NewSource(time.Now().Unix()))
	time.Sleep(time.Duration(r.Intn(3000)) * time.Millisecond)

	num := r.Intn(10)
	c <- num
}

func main() {
	fmt.Println("what is today's lucky number?")

	c := make(chan int)
	go getLuckyNum(c)
	num := <-c
	fmt.Printf("Today's your lucky number is %d!\n", num)

	close(c)
}
