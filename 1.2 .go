package main

import "fmt"

//defer第一个func 跳过，执行if语句，输出2，return到第一个func，输出1，程序结束，不会再到第二个func

func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()
	if a {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
