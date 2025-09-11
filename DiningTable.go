package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		philosopher(i, false, false)
		fork(i, 0)
	}
}

func philosopher(id int, right bool, left bool) {
	if right && left == false {
		fmt.Println("Taking Fork")
		right = false
		left = false
	}
}

func fork(id int, requests int) {
	if requests == 1 {
		// Do a goroutine
	}
}
