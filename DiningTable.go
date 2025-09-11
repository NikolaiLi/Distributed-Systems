package main

func main() {
	forkChannels := make([]chan Request, 5)
	for i := 0; i < 5; i++ {
		forkChannels[i] = make(chan Request)
		go fork(i, forkChannels[i])
	}

	for i := 0; i < 5; i++ {
		// Instantiér filosofferne
		// Tildel filosofferne deres gafler
	}

}

func philosopher(id int, right chan Request, left chan Request) {
	// Sende requests til højre og venstre gaffel
	// Vente på svar
	// print eating
	// Release og print thinking
}

func fork(id int, requests chan Request) {
	// Venter på requests
	// Hvis en filosof beder om det, så giv gaflen
	// Filosof releaser gaflen igen
}

type Request struct {
	philosopherID int
	action        string
	reply         chan bool
}
