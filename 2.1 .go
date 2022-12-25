package main

import (
	"fmt"
	"time"
)

func main() {
	var channel = make(chan int, 1)
	go func() {
		channel <- 1
		fmt.Println("出现")
	}()
	time.Sleep(time.Second)
}
