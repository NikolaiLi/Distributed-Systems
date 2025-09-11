package main

func main() {
	forkChannels := make([]chan Request, 5)
	for i := 0; i < 5; i++ {
		forkChannels[i] = make(chan Request)
		go fork(i, forkChannels[i])
	}

	for i := 0; i < 5; i++ {
		left := i
		right := (i + 1) % 5
		go philosopher(i, forkChannels[left], forkChannels[right])
	}

}

func philosopher(id int, left chan Request, right chan Request) {
	response := "need fork"
	replyChan := make(chan string)
	joker := Request{philosopherID: id, action: " need fork", reply: replyChan }

	// Sende requests til højre og venstre gaffel
	// Vente på svar
	// print eating
	// Release og print thinking
}

func fork(id int, requests chan Request) {
	taken := false
	for { req := <-requests } {
		if !taken && req == "need fork" {
			taken = true
			// send reply
		}
	}

	// Venter på requests
	// Hvis en filosof beder om det, så giv gaflen
	// Filosof releaser gaflen igen
}

type Request struct {
	philosopherID int
	action        string
	reply         chan bool
}
