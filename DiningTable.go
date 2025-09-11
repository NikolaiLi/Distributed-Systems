package main

func main() {
	for i := 1; i <= 5; i++ {
		go philosopher(i, false, false)
		go fork(i, 0)
		message := make(chan string)
	}
}

func philosopher(id int, right bool, left bool) {
	// Sende requests til højre og venstre gaffel
	// Vente på svar
	// print eating
	// Release og print thinking
}

func fork(id int, requests int) {
	// Venter på requests
	// Hvis en filosof beder om det, så giv gaflen
	// Filosof releaser gaflen igen
}

type Request struct {
	philosopherID int
	action        string
	reply         chan bool
}
