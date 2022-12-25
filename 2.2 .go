package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go Work("goroutine1", ch)
	ch <- true
	go Work("goroutine2", ch)
	ch <- true
	go Work("goroutine3", ch)
	ch <- true
	time.Sleep(time.Second)
	fmt.Println("successful")
}
func Work(workName string, ch chan bool) {

	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	<-ch
}
