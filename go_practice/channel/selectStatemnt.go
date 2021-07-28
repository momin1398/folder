package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano() / 1000000
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "hello"
	}()
	go func() {
		time.Sleep(time.Second)
		c2 <- "salute"
	}()
	for i := 0; i < 2; i++ { //there are 2 channel thats y 0,1 in for loop
		select {
		case msg1 := <-c1:
			fmt.Printf("The msg is %v\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("The msg is %v\n", msg2)

		}

	}
	end := time.Now().UnixNano() / 1000000
	fmt.Println("The time taken", end-start)
}
