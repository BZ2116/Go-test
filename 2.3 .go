package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go number(1, ch)
	ch <- true
	go number(2, ch)
	ch <- true
	go number(3, ch)
	ch <- true
	go number(4, ch)
	ch <- true
	go number(5, ch)
	ch <- true
	go number(6, ch)
	ch <- true
}
func number(n int, ch chan bool) {
	time.Sleep(time.Second)
	sum := 0
	t := n
	for i := shuwei(t); i < shuwei(t+1); i++ {
		sum = 0
		n = t
		for j := 1; j <= t; j++ {
			sum += multipe(t, (i/shuwei(n))%10)
			n--
		}
		if sum == i {
			fmt.Printf(" %d ", i)
		}
	}
	<-ch
}

func shuwei(n int) (r int) {
	r = 1
	for i := 1; i < n; i++ {
		r *= 10
	}
	return r
}
func multipe(n int, m int) (r int) {
	r = 1
	for i := 1; i <= n; i++ {
		r *= m
	}
	return r
}
