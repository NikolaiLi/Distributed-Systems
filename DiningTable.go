package main

func main() {
	for i := 1; i <= 5; i++ {
		philosopher(i, false, false)
		fork(i, 0)
	}
}

func philosopher(id int, right bool, left bool) {
	if right && left == false {
		message := make(chan string)
		go func() { message <- "request" }()

		// Take both forks
		right = true
		left = true
	}
}

func fork(id int, requests int) {
	if requests == 1 {
		// Do a goroutine
	}
}
