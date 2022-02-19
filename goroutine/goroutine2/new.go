package main

import "fmt"

func main() {
	n := 4
	out := make(chan int)
	go multiplyByTwo(n, out)
	fmt.Println(<-out)
}

func multiplyByTwo(num int, oust chan<- int) {
	result := num * 2

	oust <- result

}
