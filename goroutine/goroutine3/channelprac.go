package main

import (
	"fmt"
)

func main() {

	one := make(chan string)
	two := make(chan string)
	quit := make(chan string)
	go func() {

		//time.Sleep(2 * time.Second)
		one <- "hey"
	}()

	go func() {
		//time.Sleep(1 * time.Second)
		two <- "hello"
	}()
	for {
		select {
		case rec1 := <-one:
			fmt.Println("I recieved from Channel one", rec1)
		case rec2 := <-two:
			fmt.Println("I recieved from channel two!", rec2)
		case <-quit:
			return
		}
	}
}
