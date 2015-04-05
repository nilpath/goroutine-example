package main

import (
	"fmt"
	"time"
)

func printAfterDelay(delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("Delayed Message!")
}

func add(a int, b int, out chan int) {
	c := a + b
	out <- c
}

func printer(in chan int) {
	for {
		result := <-in
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(result)
	}
}

func example1() {
	fmt.Println("example 1: start..")
	go printAfterDelay(2000 * time.Millisecond)

	fmt.Println("continue..")

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("example 1: end..")
}

func example2() {
	fmt.Println("example 2: start..")

	// setup chanel
	c := make(chan int)

	//init printer
	go printer(c)

	//add numbers
	go add(5, 4, c)

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("example 2: end..")
}

func main() {
	// First example
	example1()

	// Second example
	//example2()
}
